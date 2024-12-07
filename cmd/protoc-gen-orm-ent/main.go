package main

import (
	"flag"

	"github.com/lesomnus/proto-orm/cmd/protoc-gen-orm-ent/plugin"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	p := plugin.Plugin{}

	var flags flag.FlagSet
	flags.StringVar((*string)(&p.EntPkg), "ent", "", "full package name of generate code by Ent")

	opt := protogen.Options{ParamFunc: flags.Set}
	opt.Run(p.Run)
}
