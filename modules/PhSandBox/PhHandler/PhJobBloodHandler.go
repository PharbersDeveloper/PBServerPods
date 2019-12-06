package Handler

import (
	"PhSandBox/PhRecord/PhDataSet"
	"PhSandBox/Uitl/http"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"strings"
)

func DataSetConsumerHandler() {
	c, err := kafka.NewKafkaBuilder().SetGroupId("data_set_job").BuildConsumer()
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
	err = c.Consume("data_set_job", dataSetFunc)
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
}

func dataSetFunc(key interface{}, value interface{}) {
	log.NewLogicLoggerBuilder().Build().Debug("进入 DataSet Kafka")
	var msgValue PhDataSet.DataSet
	err := kafka.DecodeAvroRecord(value.([]byte), &msgValue)

	if err != nil {
		return
	}

	param, err := json.Marshal(map[string]interface{}{
		"jobContainerId": msgValue.JobContainerId,
		"mongoId": msgValue.MongoId,
		"parent": msgValue.ParentIds,
		"colNames": msgValue.ColName,
		"length": msgValue.Length,
		"tabName": msgValue.TabName,
		"url": msgValue.Url,
		"description": msgValue.Description,
	})
	if err != nil {
		return
	}

	go func() {
		http.Post("http://localhost:8080/createDataSetsAndJob",
			map[string]string{"Content-Type": "application/json"},
			strings.NewReader(string(param)))
	}()
}
