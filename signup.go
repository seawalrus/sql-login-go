package main

import (
	"database/sql"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	db, err := sql.Open("mysql", "root:PASSWORD@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	var user string

	err = db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&user)
	log.Println(err)
	switch {
	case err == sql.ErrNoRows:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)

			return
		}

		_, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			log.Fatal(err)
			return
		}

		log.Println("user created")
		return
	case err != nil:
		log.Fatal(err)

		return
	default:
		log.Println("awesome sauce")
	}
}

func signupSubmit(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	log.Println(username, password)
	switch {
	case username == "" || password == "":
		log.Println("username or password field has no value")
	default:
		//upload shite here
		Upload(username, password)
		//upload(username, password)
		http.Redirect(w, r, "success", http.StatusSeeOther)
	}
}
