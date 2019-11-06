package Handler

import (
	"PhSandBox/PhPatterns/State"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"net/http"
)

func IdentifyHandler(w http.ResponseWriter, r *http.Request) {
	if len(r.Header.Get("Target")) > 0 {
		who := map[string]string {}
		context := State.AuthContext{}
		err :=	json.Unmarshal([]byte(r.Header.Get("Who")), &who)
		if err != nil { panic(err.Error()) }

		context.NewAuthContext(
				r.Header.Get("Target"),
				r.Header.Get("Action"),
				who,
				r.Header.Get("RealMod"),
			)

		_, err = context.DoExecute()
		var result = ""
		if err != nil {
			log.NewLogicLoggerBuilder().Build().Error(err)
			result = err.Error()
		} else {
			result = "ok"
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(result))

	}
}
