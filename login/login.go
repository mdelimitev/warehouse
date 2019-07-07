package login

import (
	"database/sql"
	"fmt"
	"net/http"
	"webstore/config"

	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

func Login(w http.ResponseWriter, req *http.Request) {
	if AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		p := req.FormValue("password")
		// is there a username?
		var eu user
		checkuser := `
		SELECT  id, password FROM users 
		WHERE username = lower($1)`
		err := config.DB.QueryRow(checkuser, un).Scan(&eu.id, &eu.Password)
		fmt.Println(eu.Password)
		if err == sql.ErrNoRows {
			http.Error(w, "Username and/or password do not match ", http.StatusForbidden)
			return
		}

		// does the entered password match the stored password?
		err = bcrypt.CompareHashAndPassword(eu.Password, []byte(p))
		if err != nil {
			http.Error(w, "Username and/or password do not match ", http.StatusForbidden)
			return
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
		VALUES ($1,$2) RETURNING id`
		_, err = config.DB.Exec(insertsession, sID.String(), eu.id)
		if err != nil {
			fmt.Println("error loggin session", err)
		}
		// redirect
		http.Redirect(w, req, "/home", http.StatusSeeOther)
		return
	}
	config.TPL.ExecuteTemplate(w, "login.gohtml", nil)
	fmt.Println("Executing login.gohtml")

}

func Logout(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		fmt.Println("err", err)
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	sdel := `DELETE FROM sessions where sid=$1`
	_, err = config.DB.Query(sdel, c.Value)
	if err != nil {
		fmt.Println("err", err)
		fmt.Println("Session not deleted", c.Value)
	}
	c = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, c)
	fmt.Println("Logged out")
	http.Redirect(w, req, "/login", http.StatusSeeOther)
}
