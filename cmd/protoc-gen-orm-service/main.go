package main

import (
	"github.com/lesomnus/proto-orm/cmd/protoc-gen-orm-service/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	p := plugin.NewPlugin()

	opt := protogen.Options{}
	opt.Run(p.Run)
}
