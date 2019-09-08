package Middleware

import (
	http2 "SandBox/Util/http"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"github.com/manyminds/api2go"
)

var CheckToken CheckTokenMiddleware

type CheckTokenMiddleware struct {
	Args []string
	db   *BmMongodb.BmMongodb
	rd   *BmRedis.BmRedis
}

type result struct {
	AllScope         string  `json:"all_scope"`
	AuthScope        string  `json:"auth_scope"`
	UserID           string  `json:"user_id"`
	ClientID         string  `json:"client_id"`
	Expires          float64 `json:"expires_in"`
	RefreshExpires   float64 `json:"refresh_expires_in"`
	Error            string  `json:"error"`
	ErrorDescription string  `json:"error_description"`
}

func (ctm CheckTokenMiddleware) NewCheckTokenMiddleware(args ...interface{}) CheckTokenMiddleware {
	var r *BmRedis.BmRedis
	var m *BmMongodb.BmMongodb
	var ag []string
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmRedis" {
					r = dm.(*BmRedis.BmRedis)
				}
				if tm.Name() == "BmMongodb" {
					m = dm.(*BmMongodb.BmMongodb)
				}
			}
		} else if i == 1 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {
		}
	}

	CheckToken = CheckTokenMiddleware{Args: ag, rd: r, db: m}
	return CheckToken
}

func (ctm CheckTokenMiddleware) DoMiddleware(c api2go.APIContexter, w http.ResponseWriter, r *http.Request) {
	bmlog.StandardLogger().Info("Token Middleware")
	if _, err := ctm.CheckTokenFormFunction(w, r); err != nil {
		panic(err.Error())
	}
}

func (ctm CheckTokenMiddleware) CheckTokenFormFunction(w http.ResponseWriter, r *http.Request) (rst *result, err error) {
	w.Header().Add("Content-Type", "application/json")

	// 拼接转发的URL
	scheme := "http://"
	if r.TLS != nil {
		scheme = "https://"
	}
	version := strings.Split(r.URL.Path, "/")[1]
	resource := fmt.Sprint(ctm.Args[0], "/"+version+"/", "TokenValidation")
	mergeURL := strings.Join([]string{scheme, resource}, "")

	response :=http2.Post(mergeURL,r.Header, nil)
	// 转发
	//client := &http.Client{}
	//req, _ := http.NewRequest("POST", mergeURL, nil)
	//for k, v := range r.Header {
	//	req.Header.Add(k, v[0])
	//}
	//response, err := client.Do(req)
	//if err != nil {
	//	return
	//}

	//body, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	return
	//}

	temp := result{}
	err = json.Unmarshal(response, &temp)
	if err != nil {
		return
	}

	if temp.Error != "" {
		err = errors.New(temp.ErrorDescription)
		return
	}

	return &temp, err
}
