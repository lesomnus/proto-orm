package plugin

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/go-openapi/inflect"
	"github.com/lesomnus/proto-orm/graph"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	//go:embed *.go.tpl
	template_files embed.FS
)

func (p *Plugin) Print(e *graph.Entity, f *protogen.GeneratedFile) error {
	if err := p.NewTemplate(e, f).ExecuteTemplate(f, "getter.go.tpl", e); err != nil {
		return err
	}

	return nil
}

func (p *Plugin) NewTemplate(e *graph.Entity, f *protogen.GeneratedFile) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{
		"debug": func(v any) string {
			fmt.Fprintln(os.Stderr, v)
			return ""
		},
		"pascal": func(v string) string {
			return inflect.Camelize(v)
		},

		"pb": func(name string) string {
			return f.QualifiedGoIdent(e.File.GoImportPath.Ident(name))
		},
		"go_type": func(f_ graph.Field) string {
			switch v := f_.(type) {
			case *graph.Attr:
				return v.GoType(f.QualifiedGoIdent)
			default:
				panic("invalid field")
			}
		},
		"proto_go_type": func(f_ graph.Field) string {
			switch v := f_.(type) {
			case *graph.Attr:
				return v.ProtoGoType(f.QualifiedGoIdent)
			default:
				panic("invalid field")
			}
		},

		"is_attr": func(f graph.Field) bool {
			_, ok := f.(*graph.Attr)
			return ok
		},
		"as_attr": func(f graph.Field) *graph.Attr {
			v, ok := f.(*graph.Attr)
			if !ok {
				return nil
			}
			return v
		},
		"is_edge": func(f graph.Field) bool {
			_, ok := f.(*graph.Edge)
			return ok
		},
		"as_edge": func(f graph.Field) *graph.Edge {
			v, ok := f.(*graph.Edge)
			if !ok {
				return nil
			}
			return v
		},
		"key_is_field": func(k graph.Key) bool {
			_, ok := k.(graph.Field)
			return ok
		},
		"key_as_field": func(k graph.Key) graph.Field {
			v, ok := k.(graph.Field)
			if !ok {
				return nil
			}
			return v
		},
		"key_is_index": func(k graph.Key) bool {
			_, ok := k.(*graph.Index)
			return ok
		},
		"key_as_index": func(k graph.Key) *graph.Index {
			v, ok := k.(*graph.Index)
			if !ok {
				return nil
			}
			return v
		},
	})

	t, err := t.ParseFS(template_files, "*")
	if err != nil {
		panic(err)
	}

	return t
}
