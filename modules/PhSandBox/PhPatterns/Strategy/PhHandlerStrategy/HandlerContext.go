package PhHandlerStrategy

import (
	"PhSandBox/PhRecord/PhEventMsg"
	"errors"
	"github.com/PharbersDeveloper/bp-go-lib/log"
)

type HandlerContext struct {
	strategy HandlerInterface
	EventMsg PhEventMsg.EventMsg
}

func (hc *HandlerContext) mapping() error {
	var err error
	switch hc.EventMsg.Type {
	case "UploadEndPoint":
		hc.strategy = &UploadEndStrategy{}
	case "AssetDataMart":
		hc.strategy = &DataMartStrategy{}
	case "ComplementAsset":
		hc.strategy = &ComplementAssetStrategy{}
	case "SetMartTags":
		hc.strategy = &SetMartTagsStrategy{}

	case "SandBoxDataSet":
		hc.strategy = &JobBloodStrategy{}
	case "PushJob":
		hc.strategy = &PushJobStrategy{}
	case "PushDs":
		hc.strategy = &PushDsStrategy{}
	case "Scheduler": // 暂时没用到 PushDs全都做了，对接完删除更改处理方式
		hc.strategy = &SchedulerStrategy{}
		
	//case "Python-FileMetaData-Test": // Test
	//	hc.strategy = &TestStrategy{}
	default:
		log.NewLogicLoggerBuilder().Build().Warn("is not implementation")
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
