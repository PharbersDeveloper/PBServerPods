package test

import (
	"SandBox/Handler"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alfredyang1986/blackmirror/bmalioss"
	. "github.com/smartystreets/goconvey/convey"
	"io/ioutil"
	"net/http/httptest"
	"os"
	"testing"
)

func TestUploadOSS(t *testing.T) {
	t.Parallel()
	Convey("Generate Link to OSS", t, func() {
		data := map[string]interface{}{
			"account-id": "5ce6d782aa60bdae2e8656e1",
		}
		gl := Handler.GenerateLinkHandler{}

		jsonData, _ := json.Marshal(data)

		req := httptest.NewRequest("POST", "/v0/GenerateLink", bytes.NewBuffer(jsonData))
		resp := httptest.NewRecorder()
		Convey("When the request is handled by the Router", func() {

			gl.GenerateLink(resp, req, nil)

			type response struct {
				Status string `json:"status"`
				Link   string `json:"link"`
			}
			body, _ := ioutil.ReadAll(resp.Body)

			rp := response{}
			err := json.Unmarshal(body, &rp)
			if err != nil {
				t.Error(err)
			}

			Convey("Response Status Is Ok", func() {
				So(rp.Status, ShouldEqual, "ok")
			})

			Convey("Upload OSS", func() {
				file, _ := os.Open("./面试问题.txt")
				objectValue, _ := ioutil.ReadAll(file)
				err := bmalioss.PutObject("pharbers-sandbox", fmt.Sprint(rp.Link, "/", "面试问题.txt"), objectValue)
				So(err, ShouldEqual, nil)
			})
		})
	})
}