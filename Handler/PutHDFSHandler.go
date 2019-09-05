// Handler
package Handler

import (
	"encoding/json"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/elodina/go-avro"
	kafkaAvro "github.com/elodina/go-kafka-avro"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type PutHDFSHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
}

func (h PutHDFSHandler) NewPutHDFSHandler(args ...interface{}) PutHDFSHandler {
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

	return PutHDFSHandler{ Method: md, HttpMethod: hm, Args: ag, db: m, rd: r }
}

func (h PutHDFSHandler) PutHDFS(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {

	params := map[string]string{}
	data, _ := ioutil.ReadFile(os.Getenv("HDFSAVROCONF"))
	rawMetricsSchema := strings.ReplaceAll(strings.ReplaceAll(string(data), "\n", ""), " ", "")
	bkc, _ := bmkafka.GetConfigInstance()
	res, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(res, &params)

	schema, _ := avro.ParseSchema(rawMetricsSchema)
	record := avro.NewGenericRecord(schema)
	record.Set("Path", params["Path"])

	encoder := kafkaAvro.NewKafkaAvroEncoder(bkc.SchemaRepositoryUrl)
	recordByteArr, _ := encoder.Encode(record)

	topic := "ListeningSandBoxOss"
	bkc.Produce(&topic, recordByteArr)

	return 0
}

func (h PutHDFSHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h PutHDFSHandler) GetHandlerMethod() string {
	return h.Method
}
