package service

import (
	mod "github.com/elim/GoCourses/Model"
	model "github.com/elim/GoCourses/Model"
	api "github.com/elim/GoCourses/Repository/APIService"
	repo "github.com/elim/GoCourses/Repository/Database"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/tokopedia/officialstore-home/pkg/log"
)

func init() {
	log.InitMock()
}

func TestService_GetCarByIDCar(t *testing.T) {
	type fields struct {
		api  api.Repo
		repo repo.Repository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   mod.DetailCarStatus
	}{
		{
			name: "case 1",
			fields: fields{
				repo: &repo.RepositoryMock{
					SelectCarByIDCarFunc:    func(id string) (a []model.GerageStatus) { return },
					SelectCarByIDGerageFunc: func(id string) (a []model.GerageStatus) { return },
				},
				api: &api.RepoMock{
					GetDetailDataFunc: func(id string) (a []model.DetailPosition) { return },
				},
			},
			args: args{
				id: "",
			},
			want: mod.DetailCarStatus{
				OwnerName:    "J",
				CarName:      "",
				IDCar:        "",
				GarageName:   "",
				Longtitude:   "",
				Latitude:     "",
				PositionName: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				api:  tt.fields.api,
				repo: tt.fields.repo,
			}
			if got := s.GetCarByIDCar(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCarByIDCar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_GetDepedancyData(t *testing.T) {
	type fields struct {
		api  api.Repo
		repo repo.Repository
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []mod.DetailPosition
	}{
		{
			name: "case 1",
			fields: fields{
				api: &api.RepoMock{
					GetDetailDataFunc: func(id string) (a []model.DetailPosition) {
						handler := func(w http.ResponseWriter, r *http.Request) {
							arr := []mod.DetailPosition{
								{
									GarageName:   "cek",
									Latitude:     "1000",
									Longtitude:   "20000",
									PositionName: "cek",
								},
							}
							b, err := json.Marshal(arr)
							if err != nil {
								panic(err)
							}

							w.Header().Set("Content-Type", "application/json")
							fmt.Fprintf(w, string(b))
						}
						serv := httptest.NewServer(http.HandlerFunc(handler))

						res, err := http.Get(serv.URL)
						if err != nil {
							log.Error(err, "HTTP Error")
						}

						body, err := ioutil.ReadAll(res.Body)
						if err != nil {
							panic(err.Error())
						}
						var cek []mod.DetailPosition
						err = json.Unmarshal(body, &cek)
						return cek
					},
				},
				repo: &repo.RepositoryMock{
					SelectCarByIDCarFunc:    func(id string) (a []model.GerageStatus) { return },
					SelectCarByIDGerageFunc: func(id string) (a []model.GerageStatus) { return },
				},
			},
			args: args{
				id: "1",
			},
			want: []mod.DetailPosition{
				{
					GarageName:   "cek",
					Latitude:     "1000",
					Longtitude:   "20000",
					PositionName: "cek",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				api:  tt.fields.api,
				repo: tt.fields.repo,
			}

			if got := s.GetDepedancyData(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetDepedancyData() = %v, want %v", got, tt.want)
			}
		})
	}
}
