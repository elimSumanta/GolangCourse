package database

import (
	mod "github.com/elim/GoCourses/Model"
	"log"
)

func (DBRes *DBResource) SelectCarByIDGerage(idGerage string) []mod.GerageStatus {
	rows, err := DBRes.stmtGetCarByIDGerage.Query(idGerage)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	responseSlice := make([]mod.GerageStatus, 0)
	for rows.Next() {
		var idCar string
		var idOwner string
		var carName string
		if err := rows.Scan(&idCar, &idOwner, &carName); err != nil {
			log.Fatal(err)
		}
		responseSlice = append(responseSlice, mod.GerageStatus{idCar, idOwner, carName})
	}

	return responseSlice
}

func (DBRes *DBResource) SelectCarByIDCar(idCar string) []mod.GerageStatus {
	rows, err := DBRes.stmtGetCarByIDCar.Query(idCar)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	responseSlice := make([]mod.GerageStatus, 0)
	for rows.Next() {
		var idCar string
		var idOwner string
		var carName string
		if err := rows.Scan(&idCar, &idOwner, &carName); err != nil {
			log.Fatal(err)
		}
		responseSlice = append(responseSlice, mod.GerageStatus{idCar, idOwner, carName})
	}

	return responseSlice
}
