package Handler

import (
	"PhSandBox/PhModel"
	"PhSandBox/PhPatterns/Strategy/PhMailStrategy"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"io/ioutil"
	"net/http"
)

// TODO: 重新设计发送不同形式的Email strategy
func SendEmail(w http.ResponseWriter, r *http.Request) {
	var result = ""
	mail := PhModel.Mail{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &mail)

	context := PhMailStrategy.MailContext{MailModel: mail}
	res, e := context.DoExec()
	if e != nil {
		log.NewLogicLoggerBuilder().Build().Error(e.Error())
	}
	result = `{"status"": "邮件已发送"}`
	log.NewLogicLoggerBuilder().Build().Info(res)
	log.NewLogicLoggerBuilder().Build().Info(result)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write([]byte(result))
}
