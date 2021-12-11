package face

type ICollector interface {
	GetBasic() interface{}
	GetRunning() interface{}
}
