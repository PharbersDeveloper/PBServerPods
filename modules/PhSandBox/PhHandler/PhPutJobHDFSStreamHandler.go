package Handler

import (
	"PhSandBox/PhRecord/PhOssTask"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"io/ioutil"
	"net/http"
)

func PutJobHDFS2Stream(w http.ResponseWriter, r *http.Request) {
	var result = ""
	params := map[string]interface{}{}
	res, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(res, &params)

	//data, _ := ioutil.ReadFile("avsc/OssTask.avsc")
	//rawMetricsSchema := strings.ReplaceAll(strings.ReplaceAll(string(data), "\n", ""), " ", "")
	//bkc, _ := bmkafka.GetConfigInstance()
	//schema, _ := avro.ParseSchema(rawMetricsSchema)
	//record := avro.NewGenericRecord(schema)
	//
	//record.Set("jobId", "")
	//record.Set("traceId", params["traceId"])
	//record.Set("ossKey", params["ossKey"])
	//record.Set("fileType", params["fileType"])
	//record.Set("fileName", params["fileName"])
	//record.Set("sheetName","")
	//record.Set("labels", params["labels"])
	//record.Set("dataCover", params["dataCover"])
	//record.Set("geoCover", params["geoCover"])
	//record.Set("markets", params["markets"])
	//record.Set("molecules", params["molecules"])
	//record.Set("providers", params["providers"])
	//
	//encoder := kafkaAvro.NewKafkaAvroEncoder(bkc.SchemaRegistryUrl)
	//recordByteArr, _ := encoder.Encode(record)
	//
	//topic := "oss_task_submit"
	//bkc.Produce(&topic, recordByteArr)

	p, err := kafka.NewKafkaBuilder().BuildProducer()
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		result = err.Error()
		return
	}
	record := PhOssTask.OssTask {
		TitleIndex: nil,
 		JobId: "",
		TraceId: params["traceId"].(string),
		OssKey: params["ossKey"].(string),
		FileType: params["fileType"].(string),
		FileName: params["fileName"].(string),
		SheetName: "",
		Labels: interface2ArrString(params["labels"]),
		DataCover: interface2ArrString(params["dataCover"]),
		GeoCover: interface2ArrString(params["geoCover"]),
		Markets: interface2ArrString(params["markets"]),
		Molecules: interface2ArrString(params["molecules"]),
		Providers: interface2ArrString(params["providers"]),
	}

	specificRecordByteArr, err := kafka.EncodeAvroRecord(&record)
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		result = err.Error()
		return
	}
	//log.NewLogicLoggerBuilder().Build().Info(specificRecordByteArr)
	//log.NewLogicLoggerBuilder().Build().Info(p)
	err = p.Produce("oss_task_submit", []byte("value"), specificRecordByteArr)
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		result = err.Error()
		return
	}

	result = "ok"
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(result))
}

func interface2ArrString(in interface{}) []string {
	var tmp []string
	for _, elem := range in.([]interface{}) {
		tmp = append(tmp, elem.(string))
	}
	return tmp
}
