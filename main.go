package main

import (
	"encoding/json"
	"github.com/kanyuanzhi/tialloy-client/ticface"
	"github.com/kanyuanzhi/tialloy-client/ticlog"
	"github.com/kanyuanzhi/tialloy-client/ticnet"
	"strings"
	"tialloy-client-demo/component"
	"tialloy-client-demo/global"
	"tialloy-client-demo/model"
	"time"
)

func GetBasic(collector *component.Collector) []byte {
	terminalBasic := collector.GetBasic()

	terminalBasicPack := model.TerminalBasicPack{
		Key:  global.Mac,
		Data: terminalBasic,
	}
	data, _ := json.Marshal(terminalBasicPack)
	return data
}

func GetRunning(collector *component.Collector) []byte {
	terminalRunning := collector.GetRunning()
	terminalRunningPack := model.TerminalRunningPack{
		Key:  global.Mac,
		Data: terminalRunning,
	}
	data, _ := json.Marshal(terminalRunningPack)
	return data
}

func OnConnStartHook(connection ticface.IConnection) {
	ticlog.Log.Infoln("connection is established, ready to send device info")
	global.IP = strings.Split(connection.GetConn().LocalAddr().String(), ":")[0]

	collector := component.NewCollector()
	global.OS = collector.HostCollector.GetBasic().OS
	global.Mac = collector.NetCollector.GetBasic().Mac
	ticlog.Log.Infof("terminal IP=%s, Mac=%s, OS=%s", global.IP, global.Mac, global.OS)

	basicData := GetBasic(collector)
	err := connection.SendMsg(101, basicData)
	if err != nil {
		ticlog.Log.Error(err)
		return
	}

	ticker := time.NewTicker(global.Object.CollectInterval * time.Second)
	for {
		select {
		case <-connection.Context().Done():
			ticlog.Log.Warnln("lose connection")
			return
		case <-ticker.C:
			runningData := GetRunning(collector)
			err = connection.SendMsg(102, runningData)
			if err != nil {
				ticlog.Log.Error(err)
				return
			}
		}
	}
}

func main() {
	client := ticnet.NewClient()
	client.SetOnConnStart(OnConnStartHook)
	client.Serve()
	//collector := component.NewCollector()
	//result := collector.HostCollector.GetBasic()
	//fmt.Printf("%s\n", result)
	//hostInfo, _ := host.Info()
	//fmt.Printf("%s", hostInfo)
}
