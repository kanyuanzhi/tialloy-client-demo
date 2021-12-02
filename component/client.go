package component

import (
	"ergate/face"
	"fmt"
	"github.com/kanyuanzhi/tialloy/utils"
	"net"
	"time"
)

type Client struct {
	ServerHost string
	ServerPort int

	OnConnStart func(connection face.IConnection)
}

func NewClient() face.IClient {
	return &Client{
		ServerHost: utils.GlobalObject.Host,
		ServerPort: utils.GlobalObject.TcpPort,
	}
}

func (c *Client) Start() {
	utils.GlobalLog.Info("client starts")
	go func() {
		var conn net.Conn
		var err error
		for {
			conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", c.ServerHost, c.ServerPort))
			if err != nil {
				utils.GlobalLog.Errorf("touch server failed, retouch now... err: %s", err)
				time.Sleep(time.Second)
				continue
			}
			break
		}
		utils.GlobalLog.Info("touch server successfully")
		dealConn := NewConnection(c, conn)
		go dealConn.Start()
	}()
}

func (c *Client) Stop() {
	panic("implement me")
}

func (c *Client) Serve() {
	c.Start()
	for {
		time.Sleep(time.Minute)
	}
}

func (c *Client) SetOnConnStart(hookFunc func(connection face.IConnection)) {
	c.OnConnStart = hookFunc
}

func (c *Client) CallOnConnStart(connection face.IConnection) {
	if c.OnConnStart != nil {
		utils.GlobalLog.Tracef("call DoConnStartHook")
		c.OnConnStart(connection)
	}
}
