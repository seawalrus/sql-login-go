package main

import (
	_ "github.com/go-sql-driver/mysql"
)

/*
func Match(username string, password string) bool {
	var databaseUsername string
	var databasePassword string
	db, err := sql.Open("mysql", "root:PASSWORD@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		log.Fatal("passwords do not match")

	}
	log.Println("you should be logged in!")
	return true
}
*/
