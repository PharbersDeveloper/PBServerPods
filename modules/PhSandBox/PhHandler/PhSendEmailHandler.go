package Handler

import (
	http2 "PhSandBox/Uitl/http"
	"encoding/json"
	"github.com/PharbersDeveloper/bp-go-lib/log"
	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// TODO: 重新设计发送不同形式的Email strategy
func SendEmail(w http.ResponseWriter, r *http.Request) {
	var result = ""
	params := map[string]string{}
	res, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(res, &params)


	timeLayout := "2006-01-02 15:04:05"
	// TODO: 未曾设计参数
	dataTimeStr := time.Unix(1575882043871 / 1000, 0).Format(timeLayout)

	b, _ := ioutil.ReadFile(os.Getenv("EMAIL_TEMPLATE"))
	reg := regexp.MustCompile("\t|\r|\n")
	userName := strings.ReplaceAll(string(b), "**UserName**", "")
	fileName := strings.ReplaceAll(userName, "**FileName**", "")
	fileType := strings.ReplaceAll(fileName, "**FileType**", "")
	fileLength := strings.ReplaceAll(fileType, "**FileLength**", strconv.Itoa(0) + "行")
	html := strings.ReplaceAll(fileLength, "**UploadTime**", dataTimeStr)
	content := reg.ReplaceAllString(html, "")

	for _, e := range BmConfig.BmGetConfigMap(os.Getenv("EMAILADDRESS"))["address"].([]interface{}) {
		body := strings.NewReader(`{
			"email": "`+ e.(string) +`",
			"subject": "SandBox文件上传记录",
			"content": "`+ content +`",
			"content-type": "text/html; charset=UTF-8"
		}`)
		_, _ = http2.Post("http://www.pharbers.com:60106/v0/SendMail",
			map[string]string{"Content-Type": "application/json"},
			body)
	}

	result = "ok"
	log.NewLogicLoggerBuilder().Build().Info(result)
}
