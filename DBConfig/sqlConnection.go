package dbconfig

import (
	"database/sql"
	"fmt"

	//DB Connection
	_ "github.com/lib/pq"
)

const (
	host     = "0.0.0.0"
	port     = 5432
	user     = "postgres"
	password = "changeme"
	dbname   = "DB_Mobilku"
)

//ExeConnection Generate Connection
func ExeConnection() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db, nil
}
