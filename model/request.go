package model

type TcpCommandRequest struct {
	Command string `json:"command"`
}

type TcpCommandResponse struct {
	Key  string `json:"key"`
	Data string `json:"data"`
}
