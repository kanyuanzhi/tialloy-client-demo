package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/kanyuanzhi/tialloy-client/ticface"
	"github.com/kanyuanzhi/tialloy-client/ticlog"
	"github.com/kanyuanzhi/tialloy-client/ticnet"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os/exec"
	"strings"
	"tialloy-client-demo/component"
	"tialloy-client-demo/global"
	"tialloy-client-demo/model"
	"time"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
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

	hostCollector := component.NewHostCollector()
	global.OS = hostCollector.GetBasic().OS
	netCollector := component.NewNetCollector()
	global.Mac = netCollector.GetBasic().Mac
	ticlog.Log.Infof("terminal IP=%s, Mac=%s, OS=%s", global.IP, global.Mac, global.OS)

	collector := component.NewCollector()
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

type CommandExecRouter struct {
	*ticnet.BaseRouter
}

func NewCommandExecRouter() ticface.IRouter {
	return &CommandExecRouter{}
}

func (cer *CommandExecRouter) Handle(request ticface.IRequest) {
	var tcpCommandRequest = model.TcpCommandRequest{}
	if err := json.Unmarshal(request.GetData(), &tcpCommandRequest); err != nil {
		ticlog.Log.Error(err)
		return
	}
	command := tcpCommandRequest.Command
	commandList := strings.Split(command, " ")
	commandName := commandList[0]
	commandParams := commandList[1:]
	ticlog.Log.Infoln(strings.Join(commandList, " "))

	cmd := exec.Command(commandName, commandParams...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		ticlog.Log.Error(err)
		return
	}
	err = cmd.Start()
	if err != nil {
		ticlog.Log.Error(err)
		return
	}
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		cmdRe := ConvertByte2String(in.Bytes(), "GB18030")
		tcpCommandResponse := &model.TcpCommandResponse{
			Key:  global.Mac,
			Data: cmdRe,
		}
		cmdReBytes, _ := json.Marshal(tcpCommandResponse)
		err = request.GetConnection().SendMsg(110, cmdReBytes)
		if err != nil {
			ticlog.Log.Error(err)
			return
		}
		fmt.Println(cmdRe)
	}
	err = cmd.Wait()
	if err != nil {
		ticlog.Log.Error(err)
	}
	tcpCommandFinishedResponse := &model.TcpCommandResponse{
		Key:  global.Mac,
		Data: "FINISHED",
	}
	finishedBytes, _ := json.Marshal(tcpCommandFinishedResponse)
	_ = request.GetConnection().SendMsg(110, finishedBytes)
}

func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

func main() {
	client := ticnet.NewClient()
	client.SetOnConnStart(OnConnStartHook)
	router := NewCommandExecRouter()
	client.AddRouter(111, router)
	client.Serve()

	//collector := component.NewCollector()
	//result := collector.HostCollector.GetBasic()
	//fmt.Printf("%s\n", result)
	//hostInfo, _ := host.Info()
	//fmt.Printf("%s", hostInfo)
}
