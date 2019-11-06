package model

import "strings"

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

func (esr EsSQLResponse) FormatSource() [][]interface{} {

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

	return dealAggregations(esr.Aggregations.(map[string]interface{}))
	return nil
}

func dealAggregations(data map[string]interface{}) [][]interface{} {
	//TODO:
	resultMap := dealAboveBuckets(data)
	result := make([][]interface{},0)

	for k , v := range resultMap {
		tmp := make([]interface{}, 0 )
		key := strings.ReplaceAll(k, ".keyword", "")
		tmp = append(tmp, key)
		tmp = append(tmp, v[:]...)
		result = append(result, tmp)
	}

	return result
}

func dealAboveBuckets(data map[string]interface{}) map[string][]interface{} {

	result := make(map[string][]interface{})

	for bucketsKey, bucketsValue := range data {
		valueMap := bucketsValue.(map[string]interface{})
		if buckets, ok := valueMap["buckets"]; ok {
			for _, item := range buckets.([]interface{}) {
				bucket := item.(map[string]interface{})
				key := bucket["key"]
				delete(bucket, "key")
				delete(bucket, "doc_count")

				sub := dealAboveBuckets(bucket)
				//TODO: check one map
				if len(sub) > 1 {
					for k, _ := range bucket {
						delete(sub, k)
					}
				}

				for subK, subV := range sub {
					for range subV {
						result[bucketsKey] = append(result[bucketsKey], key)
					}
					result[subK] = append(result[subK], subV[:]...)
				}
			}
		} else {
			result[bucketsKey] = append(result[bucketsKey], valueMap["value"])
		}
	}

	return result
}



