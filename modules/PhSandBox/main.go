package main

import (
	Handler "PhSandBox/PhHandler"
	"PhSandBox/env"
	"fmt"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"net/http"
)

func main() {

	// 本地调试打开
	env.SetLocalEnv()

	mux := http.NewServeMux()

	mux.HandleFunc("/identify", Handler.IdentifyHandler)
	mux.HandleFunc("/putJob2Stream", Handler.PutJobHDFS2Stream)
	mux.HandleFunc("/sendEmail", Handler.SendEmail)


	go func() {
		Handler.EventMsgConsumerHandler()
	}()

	port := "30001"

	log.NewLogicLoggerBuilder().Build().Debug("Listening...", port)
	http.ListenAndServe(fmt.Sprint(":", port), mux)
}
