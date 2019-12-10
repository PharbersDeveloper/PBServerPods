package http

import (
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func Post(url string, header map[string]string, body io.Reader) ([]byte, error) {
	// 转发
	client := &http.Client{
		Timeout: 1 * time.Minute,
	}
	request, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		request.Header.Add(k, v)
	}
	response, err:= client.Do(request)
	if err != nil {
		return nil, err
	}
	responseBody, err:= ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return responseBody, nil
}
