package face

type IPoster interface {
	LoadCollector(collector ICollector)
	DialServer() error
	Send(data []byte)
	Start()
}
