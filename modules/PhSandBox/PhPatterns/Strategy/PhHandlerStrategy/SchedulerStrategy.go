package PhHandlerStrategy

import (
	"PhSandBox/PhRecord/PhEventMsg"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/log"
)

type SchedulerStrategy struct {}

func (ues *SchedulerStrategy) DoExec(msg PhEventMsg.EventMsg) (interface{}, error) {
	context := map[string]interface{}{}
	log.NewLogicLoggerBuilder().Build().Debug("进入 Scheduler")
	err := json.Unmarshal([]byte(msg.Data), &context)
	if err != nil {
		return nil, err
	}

	param, err := json.Marshal(context)
	if err != nil {
		return nil, err
	}
	log.NewLogicLoggerBuilder().Build().Info(param)


	return "ok", nil
}
