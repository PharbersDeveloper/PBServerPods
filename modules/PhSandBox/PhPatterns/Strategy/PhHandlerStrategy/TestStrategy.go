package PhHandlerStrategy

import (
	"PhSandBox/PhRecord/PhEventMsg"
	"encoding/json"
)

type TestStrategy struct {}

func (ues *TestStrategy) DoExec(msg PhEventMsg.EventMsg) (interface{}, error) {
	context := map[string]interface{}{}

	err := json.Unmarshal([]byte(msg.Data), &context)
	if err != nil {
		return nil, err
	}

	return "nice", nil
}

