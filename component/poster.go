package component

import (
	"encoding/json"
	"ergate/model"
	"github.com/kanyuanzhi/tialloy/tinet"
	"github.com/kanyuanzhi/tialloy/utils"
	"log"
	"net"
	"strings"
	"time"
)

type Poster struct {
	Collector *Collector
	Conn      net.Conn
	IP        string
	IsTouched bool
}

func NewPoster() *Poster {
	return &Poster{}
}

func (p *Poster) LoadCollector(collector *Collector) {
	p.Collector = collector
}

func (p *Poster) DialServer() error {
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		utils.GlobalLog.Error(err)
		p.IsTouched = false
		return err
	}
	p.Conn = conn
	p.IP = strings.Split(conn.LocalAddr().String(), ":")[0]
	log.Println(p.IP)
	p.IsTouched = true
	return nil
}

func (p *Poster) Send(data []byte) {
	p.Conn.Write(data)
}

func (p *Poster) Start() {
	if p.IsTouched {
		info := &model.Info{}
		info.BasicInfo = p.Collector.GetBasicInfo()
		dp := tinet.NewDataPack()
		for {
			info.RunningInfo = p.Collector.GetRunningInfo()
			data, err := json.Marshal(info)
			if err != nil {
				utils.GlobalLog.Error(err)
				return
			}
			utils.GlobalLog.Println(len(data))
			msg := tinet.NewMessage(1, data)
			binaryMsg, _ := dp.Pack(msg)
			p.Send(binaryMsg)
			time.Sleep(time.Second)
		}
	} else {
		utils.GlobalLog.Warn("ergate does not touch the traffic hub")
		return
	}
}
