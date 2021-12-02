package component

import (
	"ergate/face"
	"errors"
	"fmt"
	"github.com/kanyuanzhi/tialloy/tinet"
	"github.com/kanyuanzhi/tialloy/utils"
	"net"
)

type Connection struct {
	Client       face.IClient
	Conn         net.Conn
	MsgChan      chan []byte
	ExitBuffChan chan bool
}

func NewConnection(client face.IClient, conn net.Conn) face.IConnection {
	return &Connection{
		Client: client,
		Conn:   conn,
	}
}

func (c *Connection) StartReader() {
	// TODO unfinished
	utils.GlobalLog.Info("tcp reader goroutine is running")
	defer utils.GlobalLog.Warn("tcp reader goroutine exited")
	defer c.Stop()
	for {
		select {
		case msg := <-c.MsgChan:
			if _, err := c.Conn.Write(msg); err != nil {
				utils.GlobalLog.Error(err)
				break
			}
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) StartWriter() {
	utils.GlobalLog.Info("tcp writer goroutine is running")
	for {
		select {
		case msg := <-c.MsgChan:
			if _, err := c.Conn.Write(msg); err != nil {
				utils.GlobalLog.Error(err)
				break
			}
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Start() {
	go c.StartWriter()
	//go c.StartReader()

	c.Client.CallOnConnStart(c)

	for {
		select {
		case <-c.ExitBuffChan:
			return
		}
	}
}

func (c *Connection) Stop() {

}

func (c *Connection) SendMsg(msgID uint32, data []byte) error {
	dp := tinet.NewDataPack()
	binaryMessage, err := dp.Pack(tinet.NewMessage(msgID, data))
	if err != nil {
		return errors.New(fmt.Sprintf("pack tcp msgID=%d err", msgID))
	}
	c.MsgChan <- binaryMessage
	return nil
}

func (c *Connection) GetConn() net.Conn {
	return c.Conn
}
