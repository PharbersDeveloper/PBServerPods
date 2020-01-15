package PhModel

type Mail struct {
	Operation 				string	`json:"operation"`
	TraceId					string 	`json:"traceId"`
	JobId					string	`json:"jobId"`
	FileName				string	`json:"fileName"`
	FileType				string	`json:"fileType"`
	CreateTime				int64	`json:"createTime"`
	Status					string	`json:"status"`
	Type					int		`json:"type"`
	Address					[]string `json:"address"`// 最终将统一发送者
	ProcessEmailAddress		[]string `json:"processEmailAddress"`
	DeliveryEmailAddress	[]string `json:"deliveryEmailAddress"`
}