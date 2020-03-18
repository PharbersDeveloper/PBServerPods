package Handler

import (
	"PhSandBox/PhRecord/PhAssetDataMart"
	"PhSandBox/Uitl/http"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/kafka"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"strings"
)

func DataMartConsumerHandler() {
	c, err := kafka.NewKafkaBuilder().SetGroupId("AssetDataMart").BuildConsumer()
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
	err = c.Consume("AssetDataMart", dataMartFunc)
	if err != nil {
		log.NewLogicLoggerBuilder().Build().Error(err.Error())
		return
	}
}

func dataMartFunc(key interface{}, value interface{}) {
	log.NewLogicLoggerBuilder().Build().Debug("进入 AssetDataMart Kafka")
	var msgValue PhAssetDataMart.AssetDataMart
	err := kafka.DecodeAvroRecord(value.([]byte), &msgValue)

	if err != nil {
		return
	}

	param, err := json.Marshal(map[string]interface{}{
		"assetName": msgValue.AssetName,
		"assetDescription": msgValue.AssetDescription,
		"assetVersion": msgValue.AssetVersion,
		"assetDataType": msgValue.AssetDataType,
		"providers": msgValue.Providers,
		"markets": msgValue.Markets,
		"molecules": msgValue.Molecules,
		"dataCover": msgValue.DataCover,
		"geoCover": msgValue.GeoCover,
		"labels": msgValue.Labels,
		"dfs": msgValue.Dfs,
		"martName": msgValue.MartName,
		"martUrl": msgValue.MartUrl,
		"martDataType": msgValue.MartDataType,
		"saveMode": msgValue.SaveMode,
	})

	if err != nil {
		return
	}

	go func() {
		_, _ = http.Post("http://localhost:8080/assetDataMart",
			map[string]string{"Content-Type": "application/json"},
			strings.NewReader(string(param)))
	}()
}
