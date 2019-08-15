package restapi

import (
	mod "github.com/elim/GoCourses/Model"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestArifsServer_GetDetailData(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(testArifsServer))
	type fields struct {
		urlHost string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantCode int
		wantData []mod.DetailPosition
	}{
		{
			name: "case 1 -> checking data",
			fields: fields{
				urlHost: testServer.URL,
			},
			args: args{
				id: "1",
			},
			wantCode: http.StatusOK,
			wantData: []mod.DetailPosition{
				{
					GarageName:   "cek",
					Longtitude:   "cek",
					Latitude:     "cek",
					PositionName: "cek",
				},
			},
		},
		{
			name: "case 2 -> Another Server not found",
			fields: fields{
				urlHost: "",
			},
			args: args{
				id: "1",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			serv := Init(tt.fields.urlHost)

			if got := serv.GetDetailData(tt.args.id); !reflect.DeepEqual(got, tt.wantData) {
				t.Errorf("ArifsServer.GetDetailData() = %v, want %v", got, tt.wantData)
			}
		})
	}
}

func testArifsServer(w http.ResponseWriter, r *http.Request) {
	resp := []mod.DetailPosition{
		{
			GarageName:   "cek",
			Longtitude:   "cek",
			Latitude:     "cek",
			PositionName: "cek",
		},
	}

	respjson, err := json.Marshal(resp)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, string(respjson))
}
