package service

import "ergate/model"

type Collector struct {
	BasicInfo   *model.BasicInfo   `json:"basic_info"`
	RunningInfo *model.RunningInfo `json:"running_info"`
}

func NewCollector() *Collector {
	return &Collector{
		BasicInfo:   model.NewBasicInfo(),
		RunningInfo: model.NewRunningInfo(),
	}
}

func (c Collector) GetBasicInfo() interface{} {
	panic("implement me")
}

func (c Collector) GetRunningInfo() interface{} {
	panic("implement me")
}
