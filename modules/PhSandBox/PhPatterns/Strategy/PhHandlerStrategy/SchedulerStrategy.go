package PhHandlerStrategy

import (
	"PhSandBox/PhRecord/PhEventMsg"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/log"
)

type SchedulerStrategy struct {}

func (ues *SchedulerStrategy) DoExec(msg PhEventMsg.EventMsg) (interface{}, error) {
	content := map[string]interface{}{}
	log.NewLogicLoggerBuilder().Build().Debug("进入 Scheduler")
	err := json.Unmarshal([]byte(msg.Data), &content)
	if err != nil {
		return nil, err
	}

	param, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}
	log.NewLogicLoggerBuilder().Build().Info(param)

	//jobId, err := uuid.GenerateUUID()
	//if err != nil {
	//	return "no", err
	//}

	//ues.pushBlood(jobId, content["status"].(string))

	return "ok", nil
}
