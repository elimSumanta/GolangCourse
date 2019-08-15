package apiservice

import (
	mod "github.com/elim/GoCourses/Model"
)

//go:generate moq -out init_moq.go . Repo

//Repo for detail data
type Repo interface {
	GetDetailData(id string) []mod.DetailPosition
}
