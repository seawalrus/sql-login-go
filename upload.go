package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func Upload(username string, password string) {
	db, err := sql.Open("mysql", "root:PASSWORD@tcp(127.0.0.1:3306)/testing")
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
		}
		_, err = db.Exec("INSERT INTO users(username, password) VALUES(?, ?)", username, hashedPassword)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("user created")
	case err != nil:
		log.Fatal(err)
	default:
		log.Println("hash matches success")
	}

}
