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

	p, err := kafka.NewKafkaBuilder().BuildProducer()

	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		result = err.Error()
		return
	}

	record := PhOssTask.OssTask {
		AssetId: params["assetId"].(string),
		Owner: params["owner"].(string),
		CreateTime: int64(params["createTime"].(float64)),
 		JobId: params["jobId"].(string),
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
	err = p.Produce("oss_task_submit", []byte(params["jobId"].(string)), specificRecordByteArr)
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
