package restapi

import (
	mod "github.com/elim/GoCourses/Model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ArifsServer struct {
	urlHost string
}

func Init(url string) *ArifsServer {
	return &ArifsServer{
		urlHost: url,
	}
}

func (serv *ArifsServer) GetDetailData(id string) []mod.DetailPosition {
	uri := fmt.Sprintf("%s/getcarposition/%s", serv.urlHost, id)
	resp, err := http.Get(uri)

	if err != nil {
		//		log.Fatal(err)
		return nil
	}

	defer resp.Body.Close()

	if resp.Status == "200 OK" {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err.Error())
		}
		var cek []mod.DetailPosition
		err = json.Unmarshal(body, &cek)
		return cek
	}
	return nil
}
