package main

import (
	"flag"

	"github.com/lesomnus/proto-orm/cmd/protoc-gen-orm-service/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	p := plugin.NewPlugin()

	var flags flag.FlagSet
	flags.StringVar(&p.Naming, "naming", "{{ .Name }}_svc.proto", "golang text template for output filename")

	opt := protogen.Options{ParamFunc: flags.Set}
	opt.Run(p.Run)
}
