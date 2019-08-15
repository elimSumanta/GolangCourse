package main

import (
	"fmt"

	con "github.com/elim/GoCourses/DBConfig"
	mod "github.com/elim/GoCourses/Model"

	"github.com/julienschmidt/httprouter"
)

type details interface {
	getDetails() []mod.DetailPosition
}

/*func loadData() {

	for {
		var detail []mod.DetailPosition
		detail, stat := del.GetDetailData("http://172.31.4.89:3000/api/getcarposition/3")
		if stat.StatusCode != 200 {
			fmt.Println(stat.Status)
		}
		fmt.Println(time.Now(), detail)
		time.Sleep(5 * time.Second)
	}
}*/

func init() {
	fmt.Println("Initialize Server")
	db, error := con.ExeConnection()
	//defer db.Close()
	if error != nil {
		panic(error)
	}

	engine := Init(db)
	router := httprouter.New()
	engine.StartEndPoint(router)
}

func main() {
	/*router := httprouter.New()
	router.GET("/getgeragestatus", del.GetGerageStatus)
	router.GET("/getCarStatus", del.GetCarStatus)

	log.Fatal(http.ListenAndServe(":8978", router))*/

}
