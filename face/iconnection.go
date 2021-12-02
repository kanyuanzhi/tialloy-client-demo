package face

import "net"

type IConnection interface {
	Start()
	Stop()
	SendMsg(msgID uint32, data []byte) error
	GetConn() net.Conn
}
