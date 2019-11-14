package env

import (
	"flag"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"os"
)

const (
	//Project env key
	ProjectName = "PROJECT_NAME"

	//Log env key
	LogTimeFormat = "BP_LOG_TIME_FORMAT"
	LogOutput     = "BP_LOG_OUTPUT"
	LogLevel      = "BP_LOG_LEVEL"

	//kafka env key
	KafkaConfigEnable = "BP_KAFKA_CONFIG_ENABLE"
	KafkaConfigPath = "BP_KAFKA_CONFIG_PATH"

	KafkaBroker = "BM_KAFKA_BROKER"
	SchemaRegistryUrl = "BM_KAFKA_SCHEMA_REGISTRY_URL"
	KafkaGroup = "BM_KAFKA_CONSUMER_GROUP"
	CaLocation = "BM_KAFKA_CA_LOCATION"
	CaSignedLocation = "BM_KAFKA_CA_SIGNED_LOCATION"
	SslKeyLocation = "BM_KAFKA_SSL_KEY_LOCATION"
	SslPass = "BM_KAFKA_SSL_PASS"
)

func SetLocalEnv() {
	//项目范围内的环境变量
	_ = os.Setenv(ProjectName, "SandBox")

	//log
	_ = os.Setenv(LogTimeFormat, "2006-01-02 15:04:05")
	_ = os.Setenv(LogOutput, "console")
	//_ = os.Setenv(env.LogOutput, "/Users/qianpeng/bplogs/mqtt-message-storage.log")
	_ = os.Setenv(LogLevel, "info")
	//_ = os.Setenv("EMAIL_TEMPLATE", fmt.Sprint(os.Getenv("SANDBOX_HOME"), "/resource/email-template.txt"))
	//_ = os.Setenv("EMAILADDRESS", fmt.Sprint(os.Getenv("SANDBOX_HOME"), "/resource/emails.json"))

	//kafka
	_ = os.Setenv(KafkaConfigPath, "../../conf/kafka_config.json")

	_ = os.Setenv(KafkaBroker, "pharbers.com:9092")
	_ = os.Setenv(SchemaRegistryUrl, "http://pharbers.com:8081")
	_ = os.Setenv(KafkaGroup, "")
	_ = os.Setenv(CaLocation, "/Users/qianpeng/kafka/secrets/snakeoil-ca-1.crt")
	_ = os.Setenv(CaSignedLocation, "/Users/qianpeng/kafka/secrets/kafkacat-ca1-signed.pem")
	_ = os.Setenv(SslKeyLocation, "/Users/qianpeng/kafka/secrets/kafkacat.client.key")
	_ = os.Setenv(SslPass, "pharbers")
}

func SetStartingParameter() {
	mode := flag.String("mode", "prod", "Development Mode")
	flag.Parse()
	_ = os.Setenv("mode", *mode)
	log.NewLogicLoggerBuilder().Build().Info("Starting Parameter，Development Mode ", *mode)
}
