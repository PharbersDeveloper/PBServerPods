package Handler

import (
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"io/ioutil"
	"net/http"
)

func SendEmail(w http.ResponseWriter, r *http.Request) {
	var result = ""
	params := map[string]string{}
	res, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(res, &params)
	result = "ok"
	log.NewLogicLoggerBuilder().Build().Info(result)
}
