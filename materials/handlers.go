package materials

import (
	"fmt"
	"net/http"
	"strings"
	"webstore/config"
	"webstore/login"
)

var mts []Material

//view all materials
func VewAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		fmt.Println("Method not GET")
		return
	}
	mts, err := AllMaterials()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		fmt.Println("12", err)
		return
	}
	fmt.Println(mts)

	config.TPL.ExecuteTemplate(w, "viewall.gohtml", mts)
}

//ADd new material to Db
func AddMaterial(w http.ResponseWriter, req *http.Request) {
	if !login.AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		mname := req.FormValue("mname")
		mcode := req.FormValue("mcode")
		if strings.TrimSpace(mname) == "" || strings.TrimSpace(mcode) == "" {
			http.Error(w, "name or code are empty ", http.StatusForbidden)
			return
		}
		//add material in DB

		insertMaterial(config.DB, mname, mcode)
		http.Redirect(w, req, "/view", http.StatusSeeOther)

		/*addM := `
		INSERT INTO materials (name, code)
		VALUES ($1,$2) RETURNING id`
		result, err := config.DB.Exec(addM, mname, mcode)
		if err != nil {
			fmt.Println("error add new material", err)
		}
		fmt.Println(result.LastInsertId)
		// redirect
		return*/
	}
	config.TPL.ExecuteTemplate(w, "add.gohtml", nil)
	fmt.Println("Executing add.gohtml")

}

//Edit material in Db
func EditMaterial(w http.ResponseWriter, req *http.Request) {
	if !login.AlreadyLoggedIn(req) {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}
	// process form submission
	if req.Method == http.MethodPost {
		mname := req.FormValue("mname")
		mcode := req.FormValue("mcode")
		if strings.TrimSpace(mname) == "" || strings.TrimSpace(mcode) == "" {
			http.Error(w, "name or code are empty ", http.StatusForbidden)
			return
		}
		//add material in DB

		insertMaterial(config.DB, mname, mcode)
		http.Redirect(w, req, "/view", http.StatusSeeOther)

		/*addM := `
		INSERT INTO materials (name, code)
		VALUES ($1,$2) RETURNING id`
		result, err := config.DB.Exec(addM, mname, mcode)
		if err != nil {
			fmt.Println("error add new material", err)
		}
		fmt.Println(result.LastInsertId)
		// redirect
		return*/
	}
	config.TPL.ExecuteTemplate(w, "add.gohtml", nil)
	fmt.Println("Executing add.gohtml")

}
