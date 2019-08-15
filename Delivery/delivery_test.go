package delivery

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/elim/GoCourses/Model"
	apiResource "github.com/elim/GoCourses/Repository/APIService"
	dbResource "github.com/elim/GoCourses/Repository/Database"
	serv "github.com/elim/GoCourses/Service"

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
		id string
	}

	apiMock := &apiResource.RepoMock{
		GetDetailDataFunc: func(id string) (res []model.DetailPosition) {
			return []model.DetailPosition{
				{
					GarageName:   "Test1",
					Longtitude:   "1000",
					Latitude:     "1000",
					PositionName: "A",
				},
			}
		},
	}
	repoMock := &dbResource.RepositoryMock{
		SelectCarByIDCarFunc: func(id string) (res []model.GerageStatus) {
			return []model.GerageStatus{
				{
					OwnerName: "J",
					CarName:   "Test",
					IDCar:     "1",
				},
			}
		},
		SelectCarByIDGerageFunc: func(id string) (res []model.GerageStatus) { return nil },
	}

	servMock := serv.New(apiMock, repoMock)

	tests := []struct {
		name         string
		fields       fields
		args         args
		wantCode     int
		wantResponse model.DetailCarStatus
	}{
		{
			name: "Case 1",
			fields: fields{
				services: servMock,
			},
			args: args{
				id: "1",
			},
			wantCode: http.StatusOK,
			wantResponse: model.DetailCarStatus{
				OwnerName:    "J",
				CarName:      "Test",
				IDCar:        "1",
				GarageName:   "Test1",
				Longtitude:   "1000",
				Latitude:     "1000",
				PositionName: "A",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hanlder := &InitHTTPHandler{
				services: tt.fields.services,
			}

			router := httprouter.New()
			router.GET("/getCarStatus", hanlder.GetCarStatus)
			recorder := httptest.NewRecorder()
			request, _ := http.NewRequest("GET", fmt.Sprintf("/getCarStatus?id_car=%s", tt.args.id), nil)
			router.ServeHTTP(recorder, request)

			assert.Equal(t, recorder.Code, tt.wantCode, "error code")
			var response model.DetailCarStatus
			err := json.Unmarshal(recorder.Body.Bytes(), &response)
			assert.NoError(t, err, "")
			assert.Equal(t, tt.wantResponse, response, "error")
		})
	}
}
