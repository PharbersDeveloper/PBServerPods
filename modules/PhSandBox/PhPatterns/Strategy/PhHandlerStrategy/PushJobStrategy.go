package PhHandlerStrategy

import (
	"PhSandBox/PhRecord/PhEventMsg"
	"PhSandBox/PhRecord/PhOssTask"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"gopkg.in/mgo.v2/bson"
)

type PushJobStrategy struct {}

func (ues *PushJobStrategy) DoExec(msg PhEventMsg.EventMsg) (interface{}, error) {
	log.NewLogicLoggerBuilder().Build().Debug("进入 PushJob")
	p, err := kafka.NewKafkaBuilder().BuildProducer()

	var record PhOssTask.OssTask
	err = json.Unmarshal([]byte(msg.Data), &record)
	if err != nil {
		return "no", err
	}

	ues.pushBlood(record.AssetId, record.JobId)

	specificRecordByteArr, err := kafka.EncodeAvroRecord(&record)
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return "no", err
	}

	err = p.Produce("oss_task_submit", []byte(record.TraceId), specificRecordByteArr)
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return "no", err
	}
	return "ok", nil
}

func (ues *PushJobStrategy) pushBlood(assetId string, jobId string) {
	ds := map[string]interface{}{
		"mongoId": bson.NewObjectId().Hex(),
		"assetId": assetId,
		"parentIds": []string{},
		"jobId": jobId,
		"columnNames": []string{},
		"tabName": "",
		"length": 0,
		"url": "",
		"description": "PushJob",
		"status": "pending",

	}

	param, err := json.Marshal(ds)
	if err != nil {
		return
	}

	msgValue := PhEventMsg.EventMsg {
		JobId: bson.NewObjectId().Hex(),
		TraceId: bson.NewObjectId().Hex(),
		Type   : "SandBoxDataSet",
		Data   : string(param),
	}

	context := HandlerContext{EventMsg: msgValue}
	_, e := context.DoExec()
	if e != nil && e.Error() != "is not implementation" {
		log.NewLogicLoggerBuilder().Build().Error(e.Error())
		return

	}
}
