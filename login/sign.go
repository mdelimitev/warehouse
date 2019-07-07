package login

import (
	"database/sql"
	"fmt"
	"net/http"
	"webstore/config"

	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

type sessions struct {
	sID    int
	userId int
}

type user struct {
	id       int
	UserName string
	Password []byte
	First    string
	Last     string
}

func Signup(w http.ResponseWriter, req *http.Request) {
	if AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		// get form values
		id := 0
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		// username taken?
		var eu user
		checkuser := `
		SELECT (username, password) FROM users 
		WHERE username =  lower($1)`
		err := config.DB.QueryRow(checkuser, un).Scan(&eu.UserName, &eu.Password)
		if err != sql.ErrNoRows {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		//encrypt pass
		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// store user in dbUsers
		u := user{id, un, bs, f, l}
		fmt.Println(u.UserName, u.Password, u.First, u.Last)
		insertuser := `
		INSERT INTO users (username, password, firstname, lastname)
		VALUES (lower($1), $2, $3, $4) RETURNING id
		`
		_, err = config.DB.Exec(insertuser, u.UserName, u.Password, u.First, u.Last)
		if err != nil {
			fmt.Println("error", err)
			http.Redirect(w, req, "/authfail", 301)
		}

		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)

		//set session in DB
		insertsession := `
		INSERT INTO sessions (sID, user_id)
		VALUES ($1,(SELECT id from users WHERE username = lower($2)) RETURNING id`
		_, err = config.DB.Exec(insertsession, sID.String(), u.UserName)
		if err != nil {
			fmt.Println("error signup session", err)
			return
		}
		// redirect
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return

	}
	config.TPL.ExecuteTemplate(w, "signup.gohtml", nil)

}
