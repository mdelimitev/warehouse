package login

import (
	"fmt"
	"net/http"
	"webstore/config"
)

func AlreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	selectuser := `
		SELECT user_id from sessions WHERE sID = $1`
	_, err = config.DB.Query(selectuser, (c.Value))
	if err != nil {
		fmt.Println("error get user", err)
		return false
	}
	return true
}

func GetUserID(w http.ResponseWriter, req *http.Request) int {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
	}
	var userID int
	selectuser := `
		SELECT user_id from sessions WHERE sID = $1`
	err = config.DB.QueryRow(selectuser, (c.Value)).Scan(&userID)
	if err != nil {
		fmt.Println("error get user", err)
	}
	return userID
}
