package main

import (
	"encoding/json"
	"github.com/kanyuanzhi/tialloy-client/ticface"
	"github.com/kanyuanzhi/tialloy-client/ticnet"
	"github.com/kanyuanzhi/tialloy-client/utils"
	"strings"
	"tialloy-client-demo/component"
	"tialloy-client-demo/model"
	utils2 "tialloy-client-demo/utils"
	"time"
)

func OnConnStartHook(connection ticface.IConnection) {
	utils.GlobalLog.Infoln("connection is established, ready to send device info")
	ip := strings.Split(connection.GetConn().LocalAddr().String(), ":")[0]

	info := model.NewInfo()
	var infoData []byte

	collector := component.NewCollector()
	collector.NetCollector.SetIP(ip)

	basicInfo := collector.GetBasicInfo()
	utils.GlobalLog.Infoln(basicInfo.NetBasicInfo)
	info.BasicInfo = basicInfo

	infoData, _ = json.Marshal(info)
	connection.SendMsg(1, infoData)

	ticker := time.NewTicker(utils2.CustomGlobalObject.CollectInterval * time.Second)
	for {
		select {
		case <-connection.Context().Done():
			utils.GlobalLog.Warnln("lose connection")
			return
		case <-ticker.C:
			runningInfo := collector.GetRunningInfo()
			info.RunningInfo = runningInfo
			utils.GlobalLog.Traceln(info.BasicInfo.NetBasicInfo.Mac)
			infoData, _ = json.Marshal(info)
			connection.SendMsg(1, infoData)
		}
	}
}

func main() {
	client := ticnet.NewClient()
	client.SetOnConnStart(OnConnStartHook)
	client.Serve()
}
