package Handler

import (
	"PhSandBox/PhPatterns/Strategy/PhHandlerStrategy"
	"PhSandBox/PhRecord/PhEventMsg"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"github.com/PharbersDeveloper/bp-go-lib/log"
)

func EventMsgConsumerHandler() {
	log.NewLogicLoggerBuilder().Build().Info("EventMsg Open")
	c, err := kafka.NewKafkaBuilder().SetGroupId("oss_msg_golang").BuildConsumer()
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
	err = c.Consume("oss_msg_test", eventMsgFunc)
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
}

func eventMsgFunc(key interface{}, value interface{}) {
	var msgValue PhEventMsg.EventMsg
	err := kafka.DecodeAvroRecord(value.([]byte), &msgValue)
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}

	context := PhHandlerStrategy.HandlerContext{EventMsg: msgValue}
	res, e := context.DoExec()
	if e != nil {
		log.NewLogicLoggerBuilder().Build().Error(e.Error())
	}

	log.NewLogicLoggerBuilder().Build().Info(res.(string))
}