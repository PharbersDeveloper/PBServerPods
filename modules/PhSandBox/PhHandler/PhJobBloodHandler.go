package Handler

import (
	"PhSandBox/PhRecord/PhDataSet"
	"PhSandBox/PhRecord/PhJob"
	"PhSandBox/Uitl/http"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"strings"
	"time"
)
// TODO：没用TS写kafka，没写好，巨丑无比，很畸形，写的真累，没设计好
type findResult struct {
	Data []map[string]interface{} `json:"data"`
}

type insertResult struct {
	Data map[string]interface{} `json:"data"`
}

type job struct {
	JobId string `json:"jobId"`
	Status string `json:"status"`
	Create int64 `json:"create"`
	Update int64 `json:"update"`
}

type dataSet struct {
	ColNames []string `json:"colNames"`
	TabName string `json:"tabName"`
	Length int32 `json:"length"`
	Url string `json:"url"`
}

func DataSetConsumerHandler() {
	c, err := kafka.NewKafkaBuilder().BuildConsumer()
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

func JobConsumerHandler() {
	c, err := kafka.NewKafkaBuilder().BuildConsumer()
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
	err = c.Consume("job_status", jobFunc)
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
}

func dataSetFunc(key interface{}, value interface{}) {
	var msgValue PhDataSet.DataSet
	err := kafka.DecodeAvroRecord(value.([]byte), &msgValue)
	if err != nil { log.NewLogicLoggerBuilder().Build().Error(err.Error());return }
	// 查询是否有Job
	jobId := msgValue.JobId
	jobResult := http.Get("http://localhost:8080/jobs?filter=(:and,(jobId,:eq,`"+ jobId +"`))")
	isFindResult := findResult{}
	err = json.Unmarshal(jobResult, &isFindResult)
	if err != nil { log.NewLogicLoggerBuilder().Build().Error(err.Error());return }

	var jobObjId = ""

	if len(isFindResult.Data) > 0 {
		jobObjId = isFindResult.Data[0]["id"].(string)
	} else {
		// 创建Job记录
		timeN := time.Now().UnixNano() / 1e6
		jobParam, err := json.Marshal(map[string]interface{}{
			"data": map[string]interface{}{
				"attributes": job{
					JobId: jobId,
					Status: "commit",
					Create: timeN,
					Update: timeN,
				},
				"type": "jobs",
			},
		})
		jobInsResult := http.Post("http://localhost:8080/jobs",
			map[string]string{"Content-Type": "application/vnd.api+json"},
			strings.NewReader(string(jobParam)))
		jobInsertResult := insertResult{}
		err = json.Unmarshal(jobInsResult, &jobInsertResult)
		if err != nil {log.NewLogicLoggerBuilder().Build().Error(err.Error());return}
		jobObjId = jobInsertResult.Data["id"].(string)
	}

	// 创建DataSets记录并与Job记录关联
	var parentNode []interface{} = []interface{}{}
	var jobObjectIds []string
	jobIds, _ := json.Marshal(msgValue.ParentIds)
	// 查出jobs表中是否包含这些id的，包含的取出id，再到dataSets表中查询
	isContainsJobResult := http.Get("http://localhost:8080/jobs?filter=(jobId,:in,"+strings.ReplaceAll(string(jobIds), "\"", "`")+")")
	jobData := findResult{}
	_ = json.Unmarshal(isContainsJobResult, &jobData)

	for _, v := range jobData.Data {
		jobObjectIds = append(jobObjectIds, v["id"].(string))
	}
	jobObjectIdsJson,  _ := json.Marshal(jobObjectIds)

	isContainsDataSetResult := http.Get("http://localhost:8080/data-sets?filter=(job,:in,"+strings.ReplaceAll(string(jobObjectIdsJson), "\"", "`")+")")
	dataSetData := findResult{}
	err = json.Unmarshal(isContainsDataSetResult, &dataSetData)
	for _, v := range dataSetData.Data {
		parentNode = append(parentNode, map[string]interface{}{
			"id": v["id"],
			"type": "data-sets",
		})
	}

	dataSetParam, err := json.Marshal(map[string]interface{}{
		"data": map[string]interface{}{
			"attributes": dataSet{
				ColNames: msgValue.ColName,
				TabName: msgValue.TabName,
				Length: msgValue.Length,
				Url: msgValue.Url,
			},
			"relationships": map[string]interface{}{
				"job": map[string]interface{}{
					"data": map[string]interface{}{
						"type": "jobs",
						"id": jobObjId,
					},
				},
				"parent" : map[string]interface{}{
					"data": parentNode,
				},
			},
			"type": "data-sets",
		},
	})

	dataSetInsResult := http.Post("http://localhost:8080/data-sets",
		map[string]string{"Content-Type": "application/vnd.api+json"},
		strings.NewReader(string(dataSetParam)))

	dataSetInsertResult := insertResult{}
	err = json.Unmarshal(dataSetInsResult, &dataSetInsertResult)
	if err != nil {log.NewLogicLoggerBuilder().Build().Error(err.Error());return}

}

func jobFunc(key interface{}, value interface{}) {
	var msgValue PhJob.Job
	err := kafka.DecodeAvroRecord(value.([]byte), &msgValue)
	if err != nil { log.NewLogicLoggerBuilder().Build().Error(err.Error()) }

	jobId := msgValue.JobId
	jobResult := http.Get("http://localhost:8080/jobs?filter=(jobId,:eq,`"+ jobId +"`)")
	findResult := findResult{}
	err = json.Unmarshal(jobResult, &findResult)
	if err != nil { log.NewLogicLoggerBuilder().Build().Error(err.Error());return }

	if len(findResult.Data) > 0 {
		jobObjId := findResult.Data[0]["id"].(string)
		timeN := time.Now().UnixNano() / 1e6
		jobParam, err := json.Marshal(map[string]interface{}{
			"data": map[string]interface{}{
				"id": jobObjId,
				"attributes": map[string]interface{}{
					"status": msgValue.Status,
					"update": timeN,
				},
				"type": "jobs",
			},
		})
		if err != nil { log.NewLogicLoggerBuilder().Build().Error(err.Error());return }
		jobResult := http.Patch("http://localhost:8080/jobs/" + jobObjId,
			map[string]string{"Content-Type": "application/vnd.api+json"},
			strings.NewReader(string(jobParam)))

		jobInsertResult := insertResult{}
		err = json.Unmarshal(jobResult, &jobInsertResult)
		if err != nil {log.NewLogicLoggerBuilder().Build().Error(err.Error());return}
	}
}