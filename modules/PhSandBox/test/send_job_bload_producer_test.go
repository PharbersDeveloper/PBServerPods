package test

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestSendJobBloodProducer(t *testing.T) {
	//env.SetLocalEnv()
	t.Parallel()

	//Convey("SendDataSet Producer Test", t, func() {
	//	p, err := kafka.NewKafkaBuilder().BuildProducer()
	//	if err != nil {
	//		log.NewLogicLoggerBuilder().Build().Error(err.Error())
	//		return
	//	}
	//
	//	record, err := kafka.EncodeAvroRecord(&PhDataSet.DataSet{
	//		ParentIds: []string{"001"}, //JobID
	//		JobId: "002",
	//		ColName: []string{"A", "B", "C"},
	//		TabName: "TabName",
	//		Length: 10,
	//		Url: "/aa/bb/cc/001",
	//		Description: "",
	//	})
	//	log.NewLogicLoggerBuilder().Build().Infof("Serializing struct: %#v\n", record)
	//	if err != nil {
	//		log.NewLogicLoggerBuilder().Build().Error(err.Error())
	//		return
	//	}
	//	err = p.Produce("data_set_job", []byte("value"), record)
	//	if err != nil {
	//		log.NewLogicLoggerBuilder().Build().Error(err.Error())
	//		return
	//	}
	//})

	//Convey("SendJob Producer Test", t, func() {
	//	p, err := kafka.NewKafkaBuilder().BuildProducer()
	//	if err != nil {
	//		log.NewLogicLoggerBuilder().Build().Error(err.Error())
	//		return
	//	}
	//
	//	record, err := kafka.EncodeAvroRecord(&PhJob.Job{
	//		JobId: "001",
	//		Status: "pending",
	//		Error: "",
	//		Description: "",
	//	})
	//	log.NewLogicLoggerBuilder().Build().Infof("Serializing struct: %#v\n", record)
	//	if err != nil {
	//		log.NewLogicLoggerBuilder().Build().Error(err.Error())
	//		return
	//	}
	//	err = p.Produce("job_status", []byte("value"), record)
	//	if err != nil {
	//		log.NewLogicLoggerBuilder().Build().Error(err.Error())
	//		return
	//	}
	//})

	//Convey("SendJob Producer Test", t, func() {
	//	p, err := kafka.NewKafkaBuilder().BuildProducer()
	//	if err != nil {
	//		log.NewLogicLoggerBuilder().Build().Error(err.Error())
	//		return
	//	}
	//
	//	record, err := kafka.EncodeAvroRecord(&PhUploadEnd.UploadEnd{TraceId: "000", DataSetId: "000"})
	//	log.NewLogicLoggerBuilder().Build().Infof("Serializing struct: %#v\n", record)
	//	if err != nil {
	//		log.NewLogicLoggerBuilder().Build().Error(err.Error())
	//		return
	//	}
	//	err = p.Produce("upload_end_job", []byte("value"), record)
	//	if err != nil {
	//		log.NewLogicLoggerBuilder().Build().Error(err.Error())
	//		return
	//	}
	//})

	Convey("push job", t, func() {

		jobs := map[string]interface{}{
			"schemaJob": map[string]string{"jobId": bson.NewObjectId().Hex()},
			"pyJob": map[string]string{"jobId": bson.NewObjectId().Hex()},
			"dataMartJob": map[string]string{"jobId": bson.NewObjectId().Hex()},
		}


		for k, v := range jobs {
			m := v.(map[string]string)
			fmt.Println(m["jobId"])
			fmt.Println(k)
		}

		//param, _ := json.Marshal(map[string]string{
		//
		//})
		//go func() {
		//	_, _ = http.Post("http://localhost:8080/putJob2Stream",
		//		map[string]string{"Content-Type": "application/json"},
		//		strings.NewReader(string(param)))
		//}()
	})
}

