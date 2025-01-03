package plugin

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/lesomnus/proto-orm/graph"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Plugin struct {
	Naming string
}

func NewPlugin() *Plugin {
	return &Plugin{}
}

func (p *Plugin) Run(plugin *protogen.Plugin) error {
	plugin.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_PROTO3
	plugin.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_PROTO3
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	namer, err := template.New("naming").Parse(p.Naming)
	if err != nil {
		return fmt.Errorf(`invalid option value "naming=%s": %w`, p.Naming, err)
	}

	g, err := graph.NewGraph(plugin.Files)
	if err != nil {
		return fmt.Errorf("graph: %w", err)
	}

	printer := &Printer{
		namer: func(p string) string {
			d, n := filepath.Split(p)
			ext := filepath.Ext(n)
			n, _ = strings.CutSuffix(n, ext)

			b := &bytes.Buffer{}
			err := namer.Execute(b, struct{ Name string }{Name: n})
			if err != nil {
				panic("namer")
			}

			return filepath.Join(d, b.String())
		},

		messages: map[protoreflect.FullName]*generatedMessage{},
	}

	fs := map[string]*File{}
	for _, e := range g {
		if !e.File.Generate {
			continue
		}

		n := printer.namer(e.File.GeneratedFilenamePrefix)
		if f, ok := fs[n]; ok {
			f.Entities = append(f.Entities, e)
			continue
		}

		fs[n] = &File{
			Path:          printer.namer(e.File.Desc.Path()),
			GeneratedFile: plugin.NewGeneratedFile(n, e.File.GoImportPath),
			Entities:      []*graph.Entity{e},
		}
	}
	for _, f := range fs {
		printer.Print(f)
	}

	return nil
}
