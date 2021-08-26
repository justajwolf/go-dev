package client

import "github.com/changbaihe/go-manual/go-basic/practices/chatserver/protocol"

type Client interface {
	Dial(addrss string) error
	Start()
	Close()
	Send(cmd interface{}) error
	SetName(name string) error
	SendMess(msg string) error
	InComing() chan protocol.CmdMess
}
