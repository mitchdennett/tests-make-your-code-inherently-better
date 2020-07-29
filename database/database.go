package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	better "github.com/mitchdennett/tests-make-your-code-inherently-better"
)

//SetupDatabase is used to setup the database on startup
func SetupDatabase() *sql.DB {
	var err error
	host := "host"
	port := 5432
	user := "postgres"
	password := "password"
	dbname := "dbname"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	conn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	return conn
}

type DB struct {
	DB *sql.DB
}

//ListRecipes will list all the recipes for a given page
func (d *DB) ListRecipes(page int) ([]*better.Recipe, error) {
	queryDocStmt := `SELECT recipe_id, title from recipe limit 50 offset $1`

	var offset int
	if page-1 < 0 {
		return nil, errors.New("Bad Request")
	}

	offset = (page - 1) * 50

	rows, err := d.DB.Query(queryDocStmt, offset)
	itemsList := make([]*better.Recipe, 0)

	if err != nil || rows == nil {
		log.Println(err)
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var item better.Recipe

		if err := rows.Scan(&item.ID, &item.Title); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Println(err)
			continue
		}

		itemsList = append(itemsList, &item)

	}

	return itemsList, nil

}
