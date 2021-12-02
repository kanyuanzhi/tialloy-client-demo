package face

type IClient interface {
	Start()
	Stop()
	Serve()

	SetOnConnStart(func(connection IConnection))
	CallOnConnStart(connection IConnection)
}
