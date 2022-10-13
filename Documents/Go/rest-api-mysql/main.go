package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/getUsers", returnAllUsers).Methods("GET")
	router.HandleFunc("/saveUser", insertUser).Methods("POST")
	router.HandleFunc("/updateUser", updateUser).Methods("PUT")
	router.HandleFunc("/deleteUser", deleteUser).Methods("DELETE")
	http.Handle("/", router)
	fmt.Printf("Connected to Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func returnAllUsers(w http.ResponseWriter, r *http.Request) {
	var user User
	var arr_user []User
	var response Response

	db := connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, first_name, last_name FROM person")

	if err != nil {
		log.Print(err)
	}

	// defined Struct by field
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName); err != nil {
			log.Fatal(err.Error())
		} else {
			arr_user = append(arr_user, user)
		}
	}

	// dynamic defined for Struct
	// for rows.Next() {
	// 	//user := Users{}

	// 	s := reflect.ValueOf(&users).Elem()
	// 	numCols := s.NumField()
	// 	columns := make([]interface{}, numCols)
	// 	for i := 0; i < numCols; i++ {
	// 		field := s.Field(i)
	// 		columns[i] = field.Addr().Interface()
	// 	}

	// 	err := rows.Scan(columns...)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	} else {
	// 		arr_users = append(arr_users, users)
	// 	}
	// 	log.Println(users)
	// }

	response.Status = http.StatusOK
	response.Message = "Success"
	response.Data = arr_user

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func insertUser(w http.ResponseWriter, r *http.Request) {
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")

	_, err = db.Exec("INSERT INTO person (first_name, last_name) VALUES(?,?)",
		first_name, last_name)

	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusCreated
	response.Message = "Sucess Created"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("user_id")
	first_name := r.FormValue("first_name")
	last_name := r.FormValue("last_name")

	_, err = db.Exec("UPDATE person set first_name = ?, last_name = ? WHERE id = ?",
		first_name, last_name, id)

	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Message = "Success Update Data"
	log.Print("Updated data to Database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	var response Response

	db := connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("user_id")

	_, err = db.Exec("DELETE person WHERE id = ?", id)

	if err != nil {
		log.Print(err)
	}

	response.Status = http.StatusOK
	response.Message = "User Success Deleted"
	log.Print("User deleted")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
