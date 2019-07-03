// Package Main Program Entrance
package main

import (
	"SandBox/Factory"
	"fmt"
	"github.com/alfredyang1986/BmServiceDef/BmApiResolver"
	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"github.com/alfredyang1986/BmServiceDef/BmPodsDefine"
	"github.com/alfredyang1986/blackmirror/bmlog"
	"github.com/julienschmidt/httprouter"
	"github.com/manyminds/api2go"
	"net/http"
	"os"
)

func main() {
	logEnv := "LOG_PATH"
	logPath := os.Getenv(logEnv)
	_ = os.Setenv(logEnv, fmt.Sprint(logPath, "/SandBox/logs/Log.log"))

	version := "v0"
	prodEnv := "SANDBOX_HOME"
	bmlog.StandardLogger().Info("SandBoxPods begins, version = ", version)

	fac := Factory.Table{}
	pod := BmPodsDefine.Pod{Name: "new SandBox", Factory: fac}
	home := os.Getenv(prodEnv)
	pod.RegisterSerFromYAML(home + "/resource/service-def.yaml")

	var bmRouter BmConfig.BmRouterConfig
	bmRouter.GenerateConfig(prodEnv)

	addr := fmt.Sprint(bmRouter.Host, ":", bmRouter.Port)
	bmlog.StandardLogger().Info("Listening on", addr)

	api := api2go.NewAPIWithResolver(version, &BmApiResolver.RequestURL{Addr: addr})
	pod.RegisterAllResource(api)
	pod.RegisterAllFunctions(version, api)
	pod.RegisterAllMiddleware(api)

	handler := api.Handler().(*httprouter.Router)
	pod.RegisterPanicHandler(handler)
	err := http.ListenAndServe(":" + bmRouter.Port, handler)
	bmlog.StandardLogger().Error(err)
}
