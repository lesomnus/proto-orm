package main

import (
	"flag"

	"github.com/lesomnus/proto-orm/cmd/protoc-gen-orm-ent-grpc/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	p := plugin.NewPlugin()

	var flags flag.FlagSet
	flags.StringVar((*string)(&p.EntPkg), "ent", "", "full package name of generate code by Ent")
	flags.StringVar(&p.Naming, "naming", "{{ kebab .Entity }}.g.go", "golang text template for output filename")

	opt := protogen.Options{ParamFunc: flags.Set}
	opt.Run(p.Run)
}
