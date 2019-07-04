// Handler
package Handler

import (
	"SandBox/Util/uuid"
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"reflect"
)

type GenerateLinkHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
}

func (h GenerateLinkHandler) NewGenerateLinkHandler(args ...interface{}) GenerateLinkHandler {
	var m *BmMongodb.BmMongodb
	var r *BmRedis.BmRedis
	var hm string
	var md string
	var ag []string
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmMongodb" {
					m = dm.(*BmMongodb.BmMongodb)
				}
				if tm.Name() == "BmRedis" {
					r = dm.(*BmRedis.BmRedis)
				}
			}
		} else if i == 1 {
			md = arg.(string)
		} else if i == 2 {
			hm = arg.(string)
		} else if i == 3 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		}
	}

	return GenerateLinkHandler{ Method: md, HttpMethod: hm, Args: ag, db: m, rd: r }
}

// GenerateLink oss or hdfs link return json
func (h GenerateLinkHandler) GenerateLink(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	// @Alex 差个OAuth的Token验证
	w.Header().Add("Content-Type", "application/json")
	params := map[string]string{}
	res, _ := ioutil.ReadAll(r.Body)
	result := map[string]interface{}{}
	enc := json.NewEncoder(w)
	err := json.Unmarshal(res, &params)
	if err != nil {
		bmlog.StandardLogger().Error(err)
		result["status"] = "error"
		result["msg"] = "解析失败"
		err = enc.Encode(result)
		bmlog.StandardLogger().Error(err)
		return 1
	}
	accountId, aok := params["account-id"]

	if !aok {
		bmlog.StandardLogger().Warning("Account 参数缺失")
		result["status"] = "error"
		result["msg"] = "Account 参数缺失"
		err = enc.Encode(result)
		bmlog.StandardLogger().Error(err)
		return 1
	}
	fmt.Println(accountId)

	uid, _ := uuid.NewRandom()

	link := fmt.Sprint("/", uid.String())
	bmlog.StandardLogger().Info(link)
	result["status"] = "ok"
	result["link"] = link
	enc.Encode(result)
	return 0
}

func (h GenerateLinkHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h GenerateLinkHandler) GetHandlerMethod() string {
	return h.Method
}
