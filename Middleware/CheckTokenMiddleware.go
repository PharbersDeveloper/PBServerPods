package Middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/manyminds/api2go"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

var CheckToken CheckTokenMiddleware

type CheckTokenMiddleware struct {
	Args []string
	db         *BmMongodb.BmMongodb
	rd   *BmRedis.BmRedis
}

type roleResult struct {
	Data map[string]interface{}	`json:"data"`
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
	// Group下的权限为Admin和Owner才能修改数据，普通用户只能上传文件
	//mdb := []BmDaemons.BmDaemon{ctm.db}
	//groupMetaDataStorage := DataStorage.GroupMetaDataStorage{}.NewGroupMetaDataStorage(mdb)
	//fileMetaDataStorage := DataStorage.FileMetaDataStorage{}.NewFileMetaDataStorage(mdb)
	//req := getApi2goRequest(r, w.Header())


	if _, err := ctm.CheckTokenFormFunction(w, r); err != nil {
		panic(err.Error())
	}

	//accountId, aok := r.URL.Query()["account-id"]
	//groupId, gok := r.URL.Query()["group-id"]
	//
	//if r.Method == "PATCH" && aok && gok {
	//	req.QueryParams["account-id"] = accountId
	//	req.QueryParams["group-id"] = groupId
	//	gmd := groupMetaDataStorage.GetAll(req, -1, -1)
	//
	//	// 拼接转发的URL
	//	scheme := "http://"
	//	if r.TLS != nil {
	//		scheme = "https://"
	//	}
	//
	//	var (
	//		roleId string
	//		attributes map[string]interface{}
	//	)
	//
	//	if len(gmd) > 0 {
	//		roleId = gmd[0].RoleID
	//		version := strings.Split(r.URL.Path, "/")[1]
	//		resource := fmt.Sprint(ctm.Args[0], "/", version, "/", "roles/", roleId)
	//		mergeURL := strings.Join([]string{scheme, resource}, "")
	//
	//		// 转发
	//		response := http2.Get(mergeURL, r.Header)
	//		result := roleResult{}
	//		json.Unmarshal(response, &result)
	//		attributes = result.Data["attributes"].(map[string]interface{})
	//	}
	//
	//	req.QueryParams = map[string][]string{}
	//
	//	req.QueryParams["owner-id"] = accountId
	//
	//	fmd := fileMetaDataStorage.GetAll(req, -1, -1)
	//
	//	level, lok := attributes["level"]
	//	if (lok && level.(float64) != 1) || (len(fmd) > 0 && fmd[0].OwnerID != accountId[0]){
	//		panic("权限不足")
	//	}
	//}
}

func getApi2goRequest(r *http.Request, header http.Header) api2go.Request{
	return api2go.Request{
		PlainRequest: r,
		Header: header,
		QueryParams: map[string][]string{},
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

	// 转发
	client := &http.Client{}
	req, _ := http.NewRequest("POST", mergeURL, nil)
	for k, v := range r.Header {
		req.Header.Add(k, v[0])
	}
	response, err := client.Do(req)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	temp := result{}
	err = json.Unmarshal(body, &temp)
	if err != nil {
		return
	}

	if temp.Error != "" {
		err = errors.New(temp.ErrorDescription)
		return
	}

	return &temp, err
}