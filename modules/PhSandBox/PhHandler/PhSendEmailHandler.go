package Handler

import (
	"PhSandBox/PhModel"
	"PhSandBox/PhPatterns/Strategy/PhMailStrategy"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"io/ioutil"
	"net/http"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var result = ""
	mail := PhModel.Mail{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &mail)

	context := PhMailStrategy.MailContext{MailModel: mail}
	res, e := context.DoExec()
	if e != nil {
		log.NewLogicLoggerBuilder().Build().Error(e.Error())
		result = `{"status": "ERROR"}`
		_, _ = w.Write([]byte(result))
		return
	}
	result = `{"status": "SUCCESS"}`
	log.NewLogicLoggerBuilder().Build().Info(res)
	log.NewLogicLoggerBuilder().Build().Info(result)
	_, _ = w.Write([]byte(result))
}
