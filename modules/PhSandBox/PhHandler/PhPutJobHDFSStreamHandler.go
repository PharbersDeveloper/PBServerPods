package Handler

import (
	Record "PhSandBox/PhRecord"
	"bytes"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"io/ioutil"
	"net/http"
)

func PutJobHDFS2Stream(w http.ResponseWriter, r *http.Request) {
	params := map[string]string{}
	res, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(res, &params)

	p, err := kafka.NewKafkaBuilder().BuildProducer()
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
	record := Record.OssTask {
 		JobId: "",
		TraceId: params["traceId"],
		OssKey: params["ossKey"],
		FileType: params["fileType"],
	}

	var buf bytes.Buffer
	log.NewLogicLoggerBuilder().Build().Infof("Serializing struct: %#v\n", record)
	err = record.Serialize(&buf)
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}

	err = p.Produce("oss_task_submit", []byte(""), buf.Bytes())
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
}
