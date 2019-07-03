package Factory

import (
	"SandBox/Handler"
	"SandBox/Middleware"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
)

type Table struct{}

var MODEL_FACTORY = map[string]interface{}{

}

var STORAGE_FACTORY = map[string]interface{}{

}

var RESOURCE_FACTORY = map[string]interface{}{

}

var FUNCTION_FACTORY = map[string]interface{}{
	"CommonPanicHandle": Handler.CommonPanicHandle{},
}
var MIDDLEWARE_FACTORY = map[string]interface{}{
	"CheckTokenMiddleware": Middleware.CheckTokenMiddleware{},
}

var DAEMON_FACTORY = map[string]interface{}{
	"BmMongodbDaemon": BmMongodb.BmMongodb{},
	"BmRedisDaemon":   BmRedis.BmRedis{},
}

func (t Table) GetModelByName(name string) interface{} {
	return MODEL_FACTORY[name]
}

func (t Table) GetResourceByName(name string) interface{} {
	return RESOURCE_FACTORY[name]
}

func (t Table) GetStorageByName(name string) interface{} {
	return STORAGE_FACTORY[name]
}

func (t Table) GetDaemonByName(name string) interface{} {
	return DAEMON_FACTORY[name]
}

func (t Table) GetFunctionByName(name string) interface{} {
	return FUNCTION_FACTORY[name]
}

func (t Table) GetMiddlewareByName(name string) interface{} {
	return MIDDLEWARE_FACTORY[name]
}
