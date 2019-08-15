//Gunanya untuk init si Service
//si Service nanti butuh repository, jadi init repository dulu

package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	httpApi "github.com/elim/GoCourses/Delivery"
	api "github.com/elim/GoCourses/Repository/APIService/RestAPI"
	repo "github.com/elim/GoCourses/Repository/Database/postgre"
	serv "github.com/elim/GoCourses/Service"
)

//Engine struct for Engine or Delivery area
type Engine struct {
	serviceName string
	services    *serv.Service
}

//Init setup for initialize Service and Resources
func Init(dbConn *sql.DB) *Engine {

	//init Repository Dulu, karena service butuh repository

	arifServer := api.Init("http://172.31.4.89:3000/api")
	repos := repo.NewConnection(dbConn)
	repos.InitQuery()

	e := &Engine{
		serviceName: "POSTGRE Services",
		services:    serv.New(arifServer, repos),
	}

	return e
}

//StartEndPoint start End Point or API
func (en *Engine) StartEndPoint(router *httprouter.Router) {
	handler, err := httpApi.New(en.services)
	if err != nil {
		panic(err)
	}

	router.GET("/getgeragestatus", handler.GetGerageStatus)
	router.GET("/getCarStatus", handler.GetCarStatus)
	log.Fatal(http.ListenAndServe(":8977", router))
}
