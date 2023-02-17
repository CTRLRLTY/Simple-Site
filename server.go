package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var db *sql.DB

func handleRegister(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		confirm := r.PostFormValue("password-confirm")

		if password != confirm {
			w.Header().Set("Content-Type", "application/json")
			data := make(map[string]string)
			data["reason"] = "Password don't match"
			json.NewEncoder(w).Encode(data)
		} else {
			insert_transact := fmt.Sprintf(`insert into "User" values (%s, %s)`, email, password)
			_, err := db.Exec(insert_transact)

			if err != nil {
				panic(err)
			}

			w.Header().Set("Goto", "/menu.html")
			w.WriteHeader(http.StatusSeeOther)
		}
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/create-account", handleRegister)

	const (
		host     = "localhost"
		port     = 5400
		user     = "postgres"
		password = "password"
		dbname   = "some_db"
	)

	connection_string := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connection_string)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	if err != nil {
		panic(err)
	}

	fmt.Println("Starting!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
