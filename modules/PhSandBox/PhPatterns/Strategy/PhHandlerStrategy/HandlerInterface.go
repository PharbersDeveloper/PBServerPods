package PhHandlerStrategy

import "PhSandBox/PhRecord/PhEventMsg"

type HandlerInterface interface {
	DoExec(eventMsg PhEventMsg.EventMsg) (interface{}, error)
}