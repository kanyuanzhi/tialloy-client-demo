package main

import (
	"encoding/json"
	"github.com/kanyuanzhi/tialloy-client/ticface"
	"github.com/kanyuanzhi/tialloy-client/ticnet"
	"github.com/kanyuanzhi/tialloy-client/utils"
	"strings"
	"tialloy-client-demo/component"
	utils2 "tialloy-client-demo/utils"
	"time"
)

func OnConnStartHook(connection ticface.IConnection) {
	utils.GlobalLog.Infoln("connection is established, ready to send device info")
	ip := strings.Split(connection.GetConn().LocalAddr().String(), ":")[0]

	collector := component.NewCollector()
	collector.NetCollector.SetIP(ip)

	basicInfo := collector.GetBasicInfo()
	mac := basicInfo.Data.NetBasicInfo.Mac
	basicInfo.Key = mac
	utils.GlobalLog.Infoln(basicInfo.Data.NetBasicInfo)

	basicInfoBytes, _ := json.Marshal(basicInfo)
	connection.SendMsg(101, basicInfoBytes)

	ticker := time.NewTicker(utils2.CustomGlobalObject.CollectInterval * time.Second)
	for {
		select {
		case <-connection.Context().Done():
			utils.GlobalLog.Warnln("lose connection")
			return
		case <-ticker.C:
			runningInfo := collector.GetRunningInfo()
			runningInfo.Key = mac

			runningInfoBytes, _ := json.Marshal(runningInfo)
			connection.SendMsg(102, runningInfoBytes)
		}
	}
}

func main() {
	client := ticnet.NewClient()
	client.SetOnConnStart(OnConnStartHook)
	client.Serve()
}
