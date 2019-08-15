package delivery

import (
	model "github.com/elim/GoCourses/Model"
	apiResource "github.com/elim/GoCourses/Repository/APIService"
	dbResource "github.com/elim/GoCourses/Repository/Database"
	serv "github.com/elim/GoCourses/Service"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestInitHTTPHandler_GetGerageStatus(t *testing.T) {
	type fields struct {
		services *serv.Service
	}
	type args struct {
		id string
	}

	apiRepo := &apiResource.RepoMock{
		GetDetailDataFunc: func(id string) (res []model.DetailPosition) { return },
	}
	dbRepo := &dbResource.RepositoryMock{
		SelectCarByIDCarFunc: func(id string) (res []model.GerageStatus) { return },
		SelectCarByIDGerageFunc: func(id string) (res []model.GerageStatus) {
			return []model.GerageStatus{
				{
					OwnerName: "Ceki",
					CarName:   "J",
					IDCar:     "1",
				},
			}
		},
	}
	servMock := serv.New(apiRepo, dbRepo)

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantCode     int
		wantResponse []model.GerageStatus
	}{
		{
			name: "Sample Test",
			fields: fields{
				services: servMock,
			},
			args: args{
				id: "1",
			},
			wantCode: http.StatusOK,
			wantResponse: []model.GerageStatus{
				{
					OwnerName: "Ceki",
					CarName:   "J",
					IDCar:     "1",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hanlder := &InitHTTPHandler{
				services: tt.fields.services,
			}
			router := httprouter.New()
			router.GET("/geragestatus", hanlder.GetGerageStatus)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", fmt.Sprintf("/geragestatus?id_Gerage=%s", tt.args.id), nil)
			router.ServeHTTP(recorder, request)

			assert.Equal(t, recorder.Code, tt.wantCode, "error code")

			var response []model.GerageStatus
			err := json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.NoError(t, err, "")
			assert.Equal(t, tt.wantResponse, response, "error")

		})
	}
}

func TestInitHTTPHandler_GetCarStatus(t *testing.T) {
	type fields struct {
		services *serv.Service
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ps httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hanlder := &InitHTTPHandler{
				services: tt.fields.services,
			}
			hanlder.GetCarStatus(tt.args.w, tt.args.r, tt.args.ps)
		})
	}
}
