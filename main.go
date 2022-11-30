package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var filename = "signup.html"
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = t.ExecuteTemplate(w, filename, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		signup(w, r)
	case "/signup-submit":
		signupSubmit(w, r)
	case "/login":
		login(w, r)
	case "/login-submit":
		Match(w, r)
	case "/signupfail":
		signupFailed(w, r)
	case "/loginfail":
		loginFailed(w, r)
	case "/success":
		success(w, r)
	default:
		log.Println("User has logged onto page!")
	}
}

func main() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
	log.Println("Started site on :8080")
	login(nil, nil)
}

func signupFailed(w http.ResponseWriter, r *http.Request) {
	var filename = "signupfail.html"
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = t.ExecuteTemplate(w, filename, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func loginFailed(w http.ResponseWriter, r *http.Request) {
	var filename = "loginfail.html"
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = t.ExecuteTemplate(w, filename, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	var filename = "login.html"
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = t.ExecuteTemplate(w, filename, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func success(w http.ResponseWriter, r *http.Request) {
	var filename = "success.html"
	t, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = t.ExecuteTemplate(w, filename, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func Match(w http.ResponseWriter, r *http.Request) {
	var databaseUsername string
	var databasePassword string
	username := r.FormValue("username")
	password := r.FormValue("password")
	db, err := sql.Open("mysql", "root:PASSWORD@tcp(127.0.0.1:3306)/testing")
	if err != nil {
		log.Fatal(err)
	}
	err = db.QueryRow("SELECT username, password FROM users WHERE username=?", username).Scan(&databaseUsername, &databasePassword)
	if err != nil {
		http.Redirect(w, r, "loginfail", http.StatusSeeOther)
	}

	err = bcrypt.CompareHashAndPassword([]byte(databasePassword), []byte(password))
	if err != nil {
		http.Redirect(w, r, "loginfail", http.StatusSeeOther)
	}
	log.Println("you should be logged in!")
}

/*
func Upload(w http.ResponseWriter, r *http.Request) {
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
	http.Redirect(w, r, "success.html", http.StatusSeeOther)

}
*/
