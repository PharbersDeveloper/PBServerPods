package PhHandlerStrategy

import (
	"PhSandBox/PhRecord/PhEventMsg"
	"errors"
)

type HandlerContext struct {
	strategy HandlerInterface
	EventMsg PhEventMsg.EventMsg
}

func (hc *HandlerContext) mapping() error {
	var err error
	switch hc.EventMsg.Type {
	case "SandBoxDataSet-Test":
		hc.strategy = &JobBloodStrategy{}
	case "UploadEndPoint-Test":
		hc.strategy = &UploadEndStrategy{}
	case "AssetDataMart-Test":
		hc.strategy = &DataMartStrategy{}
	case "Python-FileMetaData-Test":
		hc.strategy = &TestStrategy{}
	default:
		err = errors.New("is not implementation")
	}
	return err
}

func (hc *HandlerContext) DoExec() (interface{}, error)  {
	err := hc.mapping()
	if err != nil {
		return nil, err
	}
	return hc.strategy.DoExec(hc.EventMsg)
}
