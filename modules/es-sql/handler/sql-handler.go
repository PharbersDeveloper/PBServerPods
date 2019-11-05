package handler

import (
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/es-sql-pods/model"
	"io/ioutil"
	"net/http"
)

const EsServer string = "http://192.168.100.174:9200"

func SqlHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Post(fmt.Sprint(EsServer, "/_sql"),
		"application/json",
		r.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var model model.EsSQLResponse
	err = json.Unmarshal(body, &model)
	if err != nil {
		fmt.Println(err)
		return
	}

	source := model.FormatSource()

	result, err := json.Marshal(source)
	if err != nil {
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}
