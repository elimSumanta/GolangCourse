package database

import (
	"database/sql"

	//DB Connection
	_ "github.com/lib/pq"
)

//Clean Architecture Implementation

type DBResource struct {
	conn                 *sql.DB
	stmtGetCarByIDGerage *sql.Stmt
	stmtGetCarByIDCar    *sql.Stmt
}

func NewConnection(db *sql.DB) *DBResource {
	return &DBResource{
		conn: db,
	}
}

func (dbRes *DBResource) InitQuery() (err error) {
	dbRes.stmtGetCarByIDGerage, err = dbRes.conn.Prepare(SelectCarByIDGerage)
	if err != nil {
		return err
	}

	dbRes.stmtGetCarByIDCar, err = dbRes.conn.Prepare(SelectCarByIDCar)
	if err != nil {
		return err
	}
	return nil
}
