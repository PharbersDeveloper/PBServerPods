package model

import (
	"errors"
)

type EsSQLResponse struct {
	Took 			int64		`json:"took"`
	TimeOut 		bool		`json:"time_out"`
	Hits			HitsDetail	`json:"hits"`
	Aggregations	interface{}	`json:"aggregations"`
}

type HitsDetail struct {
	Total		interface{}	`json:"total"`
	MaxScore	float64		`json:"max_score"`
	Hits		[]Hit
}

type Hit struct {
	Index 	string		`json:"_index"`
	Type 	string		`json:"_type"`
	Id		string		`json:"_id"`
	Score	float64		`json:"_score"`
	Source 	interface{}	`json:"_source"`
}

func (esr EsSQLResponse) FormatSource() interface{} {

	if esr.Aggregations == nil {
		if len(esr.Hits.Hits) > 0 {

			firstSource := esr.Hits.Hits[0].Source.(map[string]interface{})
			l1 := len(firstSource)
			result := make([][]interface{}, l1)
			count := 0
			for k, v := range firstSource {
				result[count] = append(result[count], k, v)
				count++
			}

			for _, hit := range esr.Hits.Hits[1:] {
				s := hit.Source.(map[string]interface{})
				for i := 0; i < l1; i++ {
					result[i] = append(result[i], s[result[i][0].(string)])
				}
			}

			return result
		}

		return nil
	}

	//TODO；重构：多维数据的简单聚合查询
	agg := getAggRec(esr.Aggregations.(map[string]interface{}))
	// 临时给杨总看下效果
	// 请使用 { "sql": "select 月份, 药品名, count(药品名.keyword) as 数量  from hx2 group by 月份.keyword, 药品名.keyword order by 月份.keyword" }
	// 测试
	tempPivot := map[string]interface{}{
		"yAxis": "药品名.keyword",
		"xAxis": "月份.keyword",
		"value": "数量",
	}
	result, err := dealAggRes(agg, tempPivot)
	if err != nil {
		panic(err.Error())
	}
	return result

}

func getAggRec(data map[string]interface{}) (result []map[string]interface{}) {
	lastMap := make(map[string]interface{})

	for aggKey, aggValue := range data {
		valueMap := aggValue.(map[string]interface{})
		if buckets, ok := valueMap["buckets"]; ok {
			for _, item := range buckets.([]interface{}) {
				bucket := item.(map[string]interface{})
				key := bucket["key"]
				delete(bucket, "key")
				delete(bucket, "doc_count")
				for _, sub := range getAggRec(bucket) {
					sub[aggKey] = key
					result = append(result, sub)
				}
			}
		} else {
			lastMap[aggKey] = valueMap["value"]
		}
	}
	if len(lastMap) != 0 {
		result = append(result, lastMap)
	}
	return result
}

func dealAggRes(dataMap []map[string]interface{}, args map[string]interface{}) (result interface{}, err error) {
	yAxis := args["yAxis"].(string)
	xAxis := args["xAxis"].(string)
	value := args["value"].(string)
	head := value
	if val, ok := args["head"]; ok {
		head = val.(string)
	}

	ySlice, yAxis := extractAxis(dataMap, yAxis)
	xSlice, xAxis := extractAxis(dataMap, xAxis)
	if len(xSlice) == 0 {
		err = errors.New("X 轴错误，key不存在")
		return
	}
	if len(ySlice) == 0 {
		err = errors.New("Y 轴错误，key不存在")
		return
	}

	// 初始化二维数组
	tmpResult := make([]interface{}, len(ySlice))
	for i := 0; i < len(ySlice); i++ {
		tmpX := make([]interface{}, len(xSlice)+1)
		tmpX[0] = ySlice[i]
		tmpResult[i] = tmpX
	}

	// 写入内容
	for _, item := range dataMap {
		y := sliceIndex(ySlice, item[yAxis])
		x := sliceIndex(xSlice, item[xAxis])
		if y > -1 && x > -1 {
			tmpY := tmpResult[y].([]interface{})
			tmpY[x+1] = item[value]
		}
	}

	// 写入表头
	firstRow := append([]interface{}{head}, xSlice...)
	tmpResult = append([]interface{}{firstRow}, tmpResult...)

	result = tmpResult
	return
}

func extractAxis(data []map[string]interface{}, key string) ([]interface{}, string) {

	arr := make([]interface{}, 0)
	// 提取
	for _, item := range data {
		if sliceIndex(arr, item[key]) != -1 {
			continue
		}
		if v, ok := item[key]; ok {
			arr = append(arr, v)
		}
	}

	return arr, key
}

func sliceIndex(slice []interface{}, item interface{}) int {
	for i, val := range slice {
		if val == item {
			return i
		}
	}
	return -1
}
