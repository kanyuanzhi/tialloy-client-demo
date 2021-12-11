package component

import (
	"fmt"
	"github.com/shirou/gopsutil/net"
	"math/rand"
	"strings"
	"tialloy-client-demo/global"
	"tialloy-client-demo/model"
)

type NetCollector struct {
	Basic   *model.TerminalNetBasic
	Running *model.TerminalNetRunning
}

func NewNetCollector() *NetCollector {
	return &NetCollector{
		Basic:   model.NewTerminalNetBasic(),
		Running: model.NewTerminalNetRunning(),
	}
}

func (nc *NetCollector) GetBasic() *model.TerminalNetBasic {
	nc.Basic.IP = global.IP
	interfaces, _ := net.Interfaces()
	for _, inter := range interfaces {
		if global.IP == "127.0.0.1" {
			if inter.Name == global.Object.NetName {
				nc.Basic.Name = inter.Name
				nc.Basic.Mac = inter.HardwareAddr
				return nc.Basic
			} else {
				for _, addr := range inter.Addrs {
					ip_ := strings.Split(addr.Addr, "/")[0]
					if ip_ == global.IP {
						nc.Basic.Name = inter.Name
						if inter.HardwareAddr == "" {
							nc.Basic.Mac = fmt.Sprintf("uncertain-%d", rand.Int31n(1000))
						} else {
							nc.Basic.Mac = inter.HardwareAddr
						}
						return nc.Basic
					}
				}
			}
		} else {
			for _, addr := range inter.Addrs {
				ip_ := strings.Split(addr.Addr, "/")[0]
				if ip_ == global.IP {
					nc.Basic.Name = inter.Name
					nc.Basic.Mac = inter.HardwareAddr
					return nc.Basic
				}
			}
			nc.Basic.Name = "uncertain"
			nc.Basic.Mac = fmt.Sprintf("uncertain-%d", rand.Int31n(1000))
			return nc.Basic
		}
	}
	return nc.Basic
}

func (nc *NetCollector) GetRunning() *model.TerminalNetRunning {
	IOCounters, _ := net.IOCounters(true) // true统计每一个网卡的信息，false统计总信息
	for _, IOCounter := range IOCounters {
		if IOCounter.Name == nc.Basic.Name {
			nc.Running.BytesSent = IOCounter.BytesSent
			nc.Running.BytesRecv = IOCounter.BytesRecv
			nc.Running.PacketsSent = IOCounter.PacketsSent
			nc.Running.PacketsRecv = IOCounter.PacketsRecv
			nc.Running.Errin = IOCounter.Errin
			nc.Running.Errout = IOCounter.Errout
			nc.Running.Dropin = IOCounter.Dropin
			nc.Running.Dropout = IOCounter.Dropout
		}
	}
	return nc.Running
}
