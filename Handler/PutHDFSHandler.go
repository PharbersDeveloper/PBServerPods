// Handler
package Handler

import (
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/alfredyang1986/blackmirror/bmkafka"
	"github.com/elodina/go-avro"
	kafkaAvro "github.com/elodina/go-kafka-avro"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"reflect"
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
	rawMetricsSchema := `{"type":"record","name":"ListeningSandBoxOss","namespace":"com.pharbers.kafka.schema","fields":[{"name":"Path","type":"string"}]}`
	schemaRepositoryUrl := "http://pharbers.com:8081"

	schema, _ := avro.ParseSchema(rawMetricsSchema)
	record := avro.NewGenericRecord(schema)
	record.Set("Path", "f172b80b-9d13-45b1-b1e0-050ef6e0eeb7/1567682024667")

	encoder := kafkaAvro.NewKafkaAvroEncoder(schemaRepositoryUrl)
	recordByteArr, _ := encoder.Encode(record)

	bkc, _ := bmkafka.GetConfigInstance()

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
