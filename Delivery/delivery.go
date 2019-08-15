package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//GetGerageStatus to get gerage status to support arif
func (hanlder *InitHTTPHandler) GetGerageStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idGerage := r.FormValue("id_Gerage")
	//service := services.New(DBresource)

	data := hanlder.services.GetCarByIDGerage(idGerage)
	b, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

//GetGerageStatus to get gerage status to support arif
func (hanlder *InitHTTPHandler) GetCarStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	idCar := r.FormValue("id_car")
	statusByCarID := hanlder.services.GetCarByIDCar(idCar)

	respjson, err := json.Marshal(statusByCarID)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(respjson))
}
