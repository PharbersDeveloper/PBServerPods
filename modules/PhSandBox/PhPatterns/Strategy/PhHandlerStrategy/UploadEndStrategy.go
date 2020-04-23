package PhHandlerStrategy

import (
	"PhSandBox/PhRecord/PhEventMsg"
	"PhSandBox/Uitl/http"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"strings"
)

type UploadEndStrategy struct {}

func (ues *UploadEndStrategy) DoExec(msg PhEventMsg.EventMsg) (interface{}, error) {
	context := map[string]interface{}{}
	log.NewLogicLoggerBuilder().Build().Debug("进入 Upload End ")
	err := json.Unmarshal([]byte(msg.Data), &context)
	if err != nil {
		return nil, err
	}

	param, err := json.Marshal(context)
	if err != nil {
		return nil, err
	}

	go func() {
		_, _ = http.Post("http://localhost:8080/uploadFileEnd",
			map[string]string{"Content-Type": "application/json"},
			strings.NewReader(string(param)))
	}()

	return "ok", nil
}