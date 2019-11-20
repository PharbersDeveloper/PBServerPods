package main

import (
	Handler "PhSandBox/PhHandler"
	"PhSandBox/env"
	"fmt"
	"log"
	"net/http"
)


func main() {

	// 本地调试打开
	env.SetLocalEnv()

	mux := http.NewServeMux()

	mux.HandleFunc("/identify", Handler.IdentifyHandler)
	mux.HandleFunc("/putJob2Stream", Handler.PutJobHDFS2Stream)
	//mux.HandleFunc("/", nil)

	// TODO: 简略粗暴启动Consumer
	go func(){
		Handler.DataSetConsumerHandler()
	}()
	go func() {
		Handler.JobConsumerHandler()
	}()

	port := "30001"


	log.Println("Listening...", port)
	http.ListenAndServe(fmt.Sprint(":", port), mux)
}