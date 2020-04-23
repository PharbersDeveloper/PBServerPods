package PhHandlerStrategy

import (
	"PhSandBox/PhRecord/PhEventMsg"
	"PhSandBox/Uitl/http"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"strings"
)

type JobBloodStrategy struct {}

func (ues *JobBloodStrategy) DoExec(msg PhEventMsg.EventMsg) (interface{}, error) {
	context := map[string]interface{}{}
	log.NewLogicLoggerBuilder().Build().Debug("进入 Blood DataSet")
	err := json.Unmarshal([]byte(msg.Data), &context)
	if err != nil {
		return nil, err
	}

	param, err := json.Marshal(context)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = http.Post("http://localhost:8080/createDataSetsAndJob",
			map[string]string{"Content-Type": "application/json"},
			strings.NewReader(string(param)))
	}()

	return "ok", nil
}
