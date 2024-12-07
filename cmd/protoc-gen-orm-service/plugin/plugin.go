package plugin

import (
	"fmt"

	"github.com/lesomnus/proto-orm/graph"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Plugin struct {
}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Run(plugin *protogen.Plugin) error {
	plugin.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_PROTO3
	plugin.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_PROTO3
	plugin.SupportedFeatures = uint64(0 |
		pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	g, err := graph.NewGraph(plugin.Files)
	if err != nil {
		return fmt.Errorf("graph: %w", err)
	}

	fs := map[string]*File{}
	for _, e := range g {
		if !e.File.Generate {
			continue
		}

		d := e.File.GeneratedFilenamePrefix + ".svc.proto"
		if f, ok := fs[d]; ok {
			f.Entities = append(f.Entities, e)
			continue
		}

		fs[d] = &File{
			GeneratedFile: plugin.NewGeneratedFile(d, e.File.GoImportPath),
			Entities:      []*graph.Entity{e},
		}
	}
	for _, f := range fs {
		NewPrinter().Print(f)
	}

	return nil
}
