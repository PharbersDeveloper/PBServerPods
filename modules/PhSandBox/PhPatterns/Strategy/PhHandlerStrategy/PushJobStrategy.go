package PhHandlerStrategy

import (
	"PhSandBox/PhRecord/PhEventMsg"
	"PhSandBox/PhRecord/PhOssTask"
	"PhSandBox/Uitl/http"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"gopkg.in/mgo.v2/bson"
	"strings"
)

type PushJobStrategy struct {}

type Job struct {
	assetId, jobId, description, mongoId string
	parentIds []string
}

//TODO：暂时写死，还没想到用什么做动态DAG
func (j *Job) getJobDefine() map[string]interface{} {
	jobs := map[string]interface{}{
		"schemaJob": map[string]string{
			"next": "pyJob",
		},
		"pyJob": map[string]string{
			"next": "null",
		},
	}
	return jobs
}

func (ues *PushJobStrategy) DoExec(msg PhEventMsg.EventMsg) (interface{}, error) {
	log.NewLogicLoggerBuilder().Build().Debug("进入 PushJob")
	p, err := kafka.NewKafkaBuilder().BuildProducer()

	var record PhOssTask.OssTask
	err = json.Unmarshal([]byte(msg.Data), &record)
	if err != nil {
		return "no", err
	}

	jobId := "schema_job_" + bson.NewObjectId().Hex()
	for i := 0; i <= 4; i++ { // 写死创建5个
		parentIds := []string{}
		schemaMongoId := bson.NewObjectId().Hex()
		schemaJob := Job{
			mongoId: schemaMongoId,
			assetId: record.AssetId,
			jobId: jobId,
			description: "schemaJob",
			parentIds: parentIds,
		}
		ues.pushBlood(schemaJob)
		parentIds = []string{schemaMongoId}

		pyJob := Job{
			mongoId: bson.NewObjectId().Hex(),
			assetId: record.AssetId,
			jobId: "clean_job_" + bson.NewObjectId().Hex(),
			description: "pyJob",
			parentIds: parentIds,
		}
		ues.pushBlood(pyJob)
	}

	record.JobId = jobId
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

func (ues *PushJobStrategy) pushBlood(job Job) {
	status := "pending"
	ds := map[string]interface{}{
		"mongoId": job.mongoId,
		"assetId": job.assetId,
		"parentIds": job.parentIds,
		"jobId": job.jobId,
		"columnNames": []string{},
		"tabName": "",
		"length": 0,
		"url": "",
		"description": job.description,
		"status": status,
	}

	param, err := json.Marshal(ds)
	if err != nil {
		return
	}

	_, _ = http.Post("http://localhost:8080/initJobs",
		map[string]string{"Content-Type": "application/json"},
		strings.NewReader(string(param)))
}
