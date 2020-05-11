package PhMailStrategy

import (
	"PhSandBox/PhModel"
	"PhSandBox/Uitl/http"
	"encoding/json"
	"github.com/alfredyang1986/BmServiceDef/BmConfig"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

type WebUploadEndEmailStrategy struct {}

func (w * WebUploadEndEmailStrategy) DoExec(mail PhModel.Mail) (interface{}, error) {
	dataTimeStr := time.Unix( mail.CreateTime / 1000, 0).Format("2006-01-02 15:04:05")
	b, _ := ioutil.ReadFile(os.Getenv("EMAIL_TEMPLATE"))
	reg := regexp.MustCompile("\t|\r|\n")

	param, _ := json.Marshal(map[string]string{
		"account": mail.Operation,
	})
	result, err := http.Post("http://oauth.pharbers.com/v0/GetAccountNameById",
		map[string]string{"Content-Type": "application/json"},
		strings.NewReader(string(param)))

	if err != nil {
		return "no", err
	}

	context := map[string]interface{}{}

	err = json.Unmarshal(result, &context)
	if err != nil {
		return nil, err
	}

	userName := strings.ReplaceAll(string(b), "**UserName**", context["accountName"].(string))
	fileName := strings.ReplaceAll(userName, "**FileName**", mail.FileName)
	fileType := strings.ReplaceAll(fileName, "**FileType**", mail.FileType)
	uploadTime := strings.ReplaceAll(fileType, "**UploadTime**", dataTimeStr)
	html := strings.ReplaceAll(uploadTime, "**Status**", mail.Status)
	content := reg.ReplaceAllString(html, "")

	for _, e := range BmConfig.BmGetConfigMap(os.Getenv("EMAILADDRESS"))["address"].([]interface{}) {
		body := strings.NewReader(`{
			"email": "`+ e.(string) +`",
			"subject": "SandBox文件上传记录",
			"content": "`+ content +`",
			"content-type": "text/html; charset=UTF-8"
		}`)
		_, _ = http.Post("http://www.pharbers.com:60106/v0/SendMail",
			map[string]string{"Content-Type": "application/json"},
			body)
	}

	return "ok", nil
}
