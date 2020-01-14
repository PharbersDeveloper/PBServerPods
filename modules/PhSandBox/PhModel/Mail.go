package PhModel

type Mail struct {
	Operation 				string
	TraceId					string
	JobId					string
	FileName				string
	Status					string
	Type					int
	Address					[]string
	ProcessEmailAddress		[]string
	DeliveryEmailAddress	[]string
}