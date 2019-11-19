package http

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) []byte {
	// 转发
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	response, _:= client.Do(request)

	responseBody, _:= ioutil.ReadAll(response.Body)
	return responseBody
}
