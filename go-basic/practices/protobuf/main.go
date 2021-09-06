package main

import (
	"log"

	"github.com/changbaihe/go-manual/go-basic/practices/protobuf/pbs"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := &pbs.UserInfo{
		Message: *proto.String("changbaihenx"),
		Length:  *proto.Int32(12),
	}
	data, err := proto.Marshal(u)
	if err != nil {
		log.Fatalf("marshling error: %v\n", err)
	}

	newInfo := &pbs.UserInfo{}
	err = proto.Unmarshal(data, newInfo)
	if err != nil {
		log.Fatalf("unmarshling error: %v\n", err)
	}
	log.Println(newInfo.GetMessage())
}
