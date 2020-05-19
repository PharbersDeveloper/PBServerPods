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
	case "SandBoxDataSet":
		hc.strategy = &JobBloodStrategy{}
	case "UploadEndPoint":
		hc.strategy = &UploadEndStrategy{}
	case "AssetDataMart":
		hc.strategy = &DataMartStrategy{}
	case "ComplementAsset":
		hc.strategy = &ComplementAssetStrategy{}
	case "SetMartTags":
		hc.strategy = &SetMartTagsStrategy{}
	case "PushJob":
		hc.strategy = &PushJobStrategy{}
	case "Scheduler":
		hc.strategy = &SchedulerStrategy{}
		
	//case "Python-FileMetaData-Test": // Test
	//	hc.strategy = &TestStrategy{}
	default:
		//log.NewLogicLoggerBuilder().Build().Warn("is not implementation")
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
