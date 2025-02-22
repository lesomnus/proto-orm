package plugin

import (
	"errors"
	"fmt"
	"maps"
	"path/filepath"
	"slices"
	"strings"

	"github.com/lesomnus/proto-orm/graph"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type Plugin struct {
	Graph  graph.Graph
	EntPkg protogen.GoImportPath

	EntConvertFile *protogen.GeneratedFile
}

func (p *Plugin) Run(plugin *protogen.Plugin) error {
	plugin.SupportedEditionsMinimum = descriptorpb.Edition_EDITION_PROTO3
	plugin.SupportedEditionsMaximum = descriptorpb.Edition_EDITION_PROTO3
	plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)

	if p.EntPkg == "" {
		return errors.New(`"ent" option must be given`)
	}

	{
		i := string(p.EntPkg)
		p_ := filepath.Join(i, "convert.pb.go")
		f := plugin.NewGeneratedFile(p_, p.EntPkg)
		p.EntConvertFile = f

		p.printPrelude(f)
		f.P("package ", filepath.Base(i))
		f.P("")
	}

	g, err := graph.NewGraph(plugin.Files)
	if err != nil {
		return fmt.Errorf("graph: %w", err)
	}

	es := slices.Collect(maps.Values(g))
	slices.SortFunc(es, func(a *graph.Entity, b *graph.Entity) int {
		return strings.Compare(a.Name(), b.Name())
	})

	errs := []error{}
	for _, e := range es {
		if !e.File.Generate {
			continue
		}

		p_ := filepath.Join(filepath.Dir(e.File.GeneratedFilenamePrefix), "schema", e.Filename())
		f := plugin.NewGeneratedFile(p_, e.File.GoImportPath+"/schema")

		p.printPrelude(f)
		f.P("package schema")
		f.P("")

		if err := p.Print(e, f); err != nil {
			err = fmt.Errorf("%s: %w", e.FullName(), err)
			errs = append(errs, err)
			continue
		}
	}
	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (p *Plugin) printPrelude(f *protogen.GeneratedFile) {
	f.P(`// Code generated by "protoc-gen-orm-ent", DO NOT EDIT.`)
	f.P("")
}
