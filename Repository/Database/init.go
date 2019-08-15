package database

import (
	mod "github.com/elim/GoCourses/Model"
)

//go:generate moq -out init_moq.go . Repository

//Repository to handle request
type Repository interface {
	SelectCarByIDGerage(id string) []mod.GerageStatus
	SelectCarByIDCar(id string) []mod.GerageStatus
}
