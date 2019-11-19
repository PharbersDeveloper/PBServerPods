package http

import (
	"io"
	"io/ioutil"
	"net/http"
)

func Patch(url string, header map[string]string, body io.Reader) []byte {
	// 转发
	client := &http.Client{}
	request, _ := http.NewRequest("PATCH", url, body)
	for k, v := range header {
		request.Header.Add(k, v)
	}
	response, _:= client.Do(request)

	responseBody, _:= ioutil.ReadAll(response.Body)
	return responseBody
}
