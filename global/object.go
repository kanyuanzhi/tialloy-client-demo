package global

import (
	"encoding/json"
	"github.com/kanyuanzhi/tialloy-client/global"
	"io/ioutil"
	"time"
)

var (
	IP     string
	Mac    string
	OS     string
	Object *Obj
)

type Obj struct {
	*global.Obj
	CollectInterval time.Duration `json:"collect_interval,omitempty"`
	NetName         string        `json:"net_name"`
	WindowsPath     string        `json:"windows_path"`
}

func (c *Obj) Reload() {
	data, err := ioutil.ReadFile("conf/tialloy_client.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &Object)
	if err != nil {
		panic(err)
	}
}

func init() {
	Object = &Obj{
		Obj:             global.Object,
		CollectInterval: 1,
		NetName:         "以太网",
		WindowsPath:     "c:",
	}
	Object.Reload()
}
