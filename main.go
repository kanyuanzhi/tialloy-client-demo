package main

import (
	"ergate/component"
	"ergate/face"
)

func OnConnStartHook(connection face.IConnection) {

}

func main() {
	client := component.NewClient()
	client.SetOnConnStart(OnConnStartHook)
	client.Serve()
}
