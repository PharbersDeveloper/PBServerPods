package PhHandlerStrategy

import (
	"PhSandBox/PhRecord/PhEventMsg"
	"PhSandBox/Uitl/http"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"strings"
)

type SetMartTagsStrategy struct {}

func (ues *SetMartTagsStrategy) DoExec(msg PhEventMsg.EventMsg) (interface{}, error) {
	context := map[string]interface{}{}
	log.NewLogicLoggerBuilder().Build().Debug("进入 SetMartTags")
	err := json.Unmarshal([]byte(msg.Data), &context)
	if err != nil {
		return nil, err
	}

	param, err := json.Marshal(context)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = http.Post("http://localhost:8080/setMartTags2Asset",
			map[string]string{"Content-Type": "application/json"},
			strings.NewReader(string(param)))
	}()

	return "ok", nil
}
