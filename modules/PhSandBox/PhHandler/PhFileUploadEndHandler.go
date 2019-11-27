package Handler

import (
	"PhSandBox/PhRecord/PhUploadEnd"
	"PhSandBox/Uitl/http"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"strings"
)

func FileUploadEndHandler() {
	c, err := kafka.NewKafkaBuilder().BuildConsumer()
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
	err = c.Consume("upload_end_job", uploadEndFunc)
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
}

func uploadEndFunc(key interface{}, value interface{}) {
	var msgValue PhUploadEnd.UploadEnd
	err := kafka.DecodeAvroRecord(value.([]byte), &msgValue)
	if err != nil { log.NewLogicLoggerBuilder().Build().Error(err.Error());return }

	param, err := json.Marshal(map[string]string{
		"jobId": msgValue.JobId,
		"traceId": msgValue.TraceId,
	})
	http.Post("http://localhost:8080/updateJobId2DataSets",
		map[string]string{"Content-Type": "application/json"},
		strings.NewReader(string(param)))

}
