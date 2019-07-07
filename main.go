package main

import (
	"net/http"
	"webstore/config"
	"webstore/login"
	"webstore/materials"
)

func main() {
	http.HandleFunc("/signup", login.Signup)
	http.HandleFunc("/", login.Login)
	http.HandleFunc("/login", login.Login)
	http.HandleFunc("/logout", login.Logout)
	http.HandleFunc("/home", home)
	http.HandleFunc("/add", materials.AddMaterial)
	http.HandleFunc("/view", materials.VewAll)
	http.HandleFunc("/edit", materials.EditMaterial)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets"))))

	http.ListenAndServe(":8080", nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	if !login.AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	config.TPL.ExecuteTemplate(w, "home.gohtml", nil)
}
