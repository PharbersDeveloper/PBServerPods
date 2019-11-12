package Handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func UpdateMongoWithJobId(w http.ResponseWriter, r *http.Request) {
	params := map[string]interface{}{}
	res, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(res, &params)


	//in := Model.FileMetaDatum{}
	//out := Model.FileMetaDatum{}
	//cond := bson.M{ "trace-id": traceId }
	//err := h.db.FindOneByCondition(&in, &out, cond)
	//if err != nil {
	//	log.NewLogicLoggerBuilder().Build().Error(err.Error())
	//	return
	//} else {
	//	out.JobID = append(out.JobID, params["jobId"].(string))
	//	err = h.db.Update(&out)
	//	if err != nil {
	//		return
	//	}
	//}
}
