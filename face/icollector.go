package face

type ICollector interface {
	GetBasicInfo() interface{}
	GetRunningInfo() interface{}
}
