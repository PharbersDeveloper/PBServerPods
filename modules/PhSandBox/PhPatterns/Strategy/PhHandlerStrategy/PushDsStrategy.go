package PhHandlerStrategy

import (
	"PhSandBox/PhRecord/PhEventMsg"
	"PhSandBox/Uitl/http"
	"encoding/json"
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"strings"
)

type PushDsStrategy struct {}

type PyJob struct {
	JobId string `json:"jobId"`
	NoticeTopic string `json:"noticeTopic"`
	MetadataPath string `json:"metadataPath"`
	FilesPath string `json:"filesPath"`
}

func (ues *PushDsStrategy) DoExec(msg PhEventMsg.EventMsg) (interface{}, error) {
	context := map[string]interface{}{}
	log.NewLogicLoggerBuilder().Build().Debug("进入 PushDs")
	err := json.Unmarshal([]byte(msg.Data), &context)
	if err != nil {
		return nil, err
	}

	param, err := json.Marshal(context)
	if err != nil {
		return nil, err
	}

	resByte, _ := http.Post("http://localhost:8080/pushDs",
		map[string]string{"Content-Type": "application/json"},
		strings.NewReader(string(param)))

	fmt.Println(string(resByte))

	var res map[string]string
	_ = json.Unmarshal(resByte, &res)

	jobs := Job{}
	nextJob := jobs.getJobDefine()[res["description"]].(map[string]string)["next"]

	if res["dsId"] != "-1" && res["status"] == "end" && nextJob != "null" {
		// TODO 执行下一个job也就是PyJob，根据空闲的py ds 查询关联的jobId ✅
		// TODO 其实JobId和RunnerId都应该变成一个微服务 ❎
		// TODO 现在还差一个initJobs ✅
		// TODO 与老邓再次对接Connector的发送消息 ❎
		// TODO 数据存储路径  res["path"]  什么Job运行结束 res["description"] ✅

		param, _ = json.Marshal(map[string]string{
			"jobName": nextJob,
			"dsId": res["dsId"],
		})
		resByte, _ = http.Post("http://localhost:8080/findNextJob",
			map[string]string{"Content-Type": "application/json"},
			strings.NewReader(string(param)))
		_ = json.Unmarshal(resByte, &res)
		jobId := res["jobId"]

		// TODO：差截断路径加拼接
		metaDataPath := res["path"][:strings.LastIndex(res["path"], "/")] + "/metadata"
		filesPath := res["path"]

		pyJob := PyJob{
			JobId: jobId,
			NoticeTopic: "HiveTaskNone",
			MetadataPath: metaDataPath,
			FilesPath: filesPath,
		}

		json, _ := json.Marshal(pyJob)

		eventMsg := PhEventMsg.EventMsg{
			TraceId: "",
			JobId: jobId,
			Type: "Python-FileMetaData",
			Data: string(json),
		}
		fmt.Println(eventMsg)
		p, _ := kafka.NewKafkaBuilder().BuildProducer()

		specificRecordByteArr, _ := kafka.EncodeAvroRecord(&eventMsg)

		err = p.Produce("oss_msg", []byte(eventMsg.TraceId), specificRecordByteArr)
		if err != nil {
			return "no", err
		}
		return "ok", nil
	}

	return "ok", nil
}
