package utils

import (
	"encoding/json"
	"github.com/kanyuanzhi/tialloy-client/utils"
	"io/ioutil"
	"time"
)

var CustomGlobalObject *CustomGlobalObj

type CustomGlobalObj struct {
	*utils.GlobalObj
	CollectInterval time.Duration `json:"collect_interval,omitempty"`
}

func (c *CustomGlobalObj) Reload() {
	data, err := ioutil.ReadFile("conf/tialloy_client.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &CustomGlobalObject)
	if err != nil {
		panic(err)
	}
}

func init() {
	CustomGlobalObject = &CustomGlobalObj{utils.GlobalObject, 1}
	CustomGlobalObject.Reload()
}
