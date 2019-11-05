package model

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

	return dealAggregations(esr.Aggregations)
}

func dealAggregations(agg interface{}) [][]interface{} {
	//TODO:
	return nil
}

