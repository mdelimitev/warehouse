package materials

import (
	"database/sql"
	"fmt"
	"log"
	"webstore/config"

	"github.com/lib/pq"
)

type Material struct {
	Id   int
	Name string
	Code string
}

func AllMaterials() ([]Material, error) {
	rows, err := config.DB.Query("select * FROM materials;")
	if err, ok := err.(*pq.Error); ok {
		// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
		fmt.Println("pq error:", err.Code.Name())
		return nil, err

	}

	defer rows.Close()

	mts := make([]Material, 0)
	for rows.Next() {
		mt := Material{}
		err := rows.Scan(&mt.Id, &mt.Name, &mt.Code) // order matters
		if err != nil {
			fmt.Println("2", err)
			return nil, err
		}

		mts = append(mts, mt)
	}
	if err = rows.Err(); err != nil {
		fmt.Println("3", err)
		return nil, err
	}
	return mts, nil

}

func insertMaterial(db *sql.DB, name string, code string) {

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare(`INSERT INTO materials (name, code)
                     VALUES($1,$2);`)
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(name, code)
	if err, ok := err.(*pq.Error); ok {
		// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
		fmt.Println("pq error:", err.Code.Name())
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

}

func editMaterial(db *sql.DB, name string, code string) {

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("UPDATE  materials SET name = 1$, code = 2$ WHERE id = $3;")
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

	defer stmt.Close()

	_, err = stmt.Exec(name, code)
	if err, ok := err.(*pq.Error); ok {
		// Here err is of type *pq.Error, you may inspect all its fields, e.g.:
		fmt.Println("pq error:", err.Code.Name())
		log.Fatal(err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
	}

}
