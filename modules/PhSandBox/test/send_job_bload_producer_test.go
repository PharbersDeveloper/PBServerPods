package test

import (
	"PhSandBox/PhRecord/PhDataSet"
	"PhSandBox/PhRecord/PhJob"
	"PhSandBox/env"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSendJobBloodProducer(t *testing.T) {
	env.SetLocalEnv()
	t.Parallel()

	Convey("SendDataSet Producer Test", t, func() {
		p, err := kafka.NewKafkaBuilder().BuildProducer()
		if err != nil {
			log.NewLogicLoggerBuilder().Build().Error(err.Error())
			return
		}

		record, err := kafka.EncodeAvroRecord(&PhDataSet.DataSet{
			ParentIds: []string{}, //JobID
			JobId: "001",
			ColName: []string{"A", "B", "C"},
			TabName: "TabName",
			Length: 10,
			Url: "/aa/bb/cc/001",
			Description: "",
		})
		log.NewLogicLoggerBuilder().Build().Infof("Serializing struct: %#v\n", record)
		if err != nil {
			log.NewLogicLoggerBuilder().Build().Error(err.Error())
			return
		}
		err = p.Produce("data_set_job", []byte("value"), record)
		if err != nil {
			log.NewLogicLoggerBuilder().Build().Error(err.Error())
			return
		}
	})

	Convey("SendJob Producer Test", t, func() {
		p, err := kafka.NewKafkaBuilder().BuildProducer()
		if err != nil {
			log.NewLogicLoggerBuilder().Build().Error(err.Error())
			return
		}

		record, err := kafka.EncodeAvroRecord(&PhJob.Job{
			JobId: "001",
			Status: "commit",
			Error: "",
			Description: "",
		})
		log.NewLogicLoggerBuilder().Build().Infof("Serializing struct: %#v\n", record)
		if err != nil {
			log.NewLogicLoggerBuilder().Build().Error(err.Error())
			return
		}
		err = p.Produce("job_status", []byte("value"), record)
		if err != nil {
			log.NewLogicLoggerBuilder().Build().Error(err.Error())
			return
		}
	})
}

