package plugin

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"github.com/go-openapi/inflect"
	orm "github.com/lesomnus/proto-orm"
	"github.com/lesomnus/proto-orm/graph"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	//go:embed *.go.tpl
	template_files embed.FS
)

func (p *Plugin) Print(e *graph.Entity, f *protogen.GeneratedFile) error {
	t := p.NewTemplate(e, f)
	if err := t.ExecuteTemplate(f, "pick.go.tpl", e); err != nil {
		return err
	}
	if err := t.ExecuteTemplate(f, "select.go.tpl", e); err != nil {
		return err
	}
	if err := t.ExecuteTemplate(f, "getter.go.tpl", e); err != nil {
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

		"def_attr_eq": func(attr *graph.Attr, a, b string) string {
			t := attr.Type()
			switch t {
			case orm.Type_TYPE_BOOL,
				orm.Type_TYPE_ENUM,
				orm.Type_TYPE_INT32,
				orm.Type_TYPE_SINT32,
				orm.Type_TYPE_SFIXED32,
				orm.Type_TYPE_UINT32,
				orm.Type_TYPE_FIXED32,
				orm.Type_TYPE_INT64,
				orm.Type_TYPE_SINT64,
				orm.Type_TYPE_SFIXED64,
				orm.Type_TYPE_UINT64,
				orm.Type_TYPE_FIXED64,
				orm.Type_TYPE_FLOAT,
				orm.Type_TYPE_DOUBLE,
				orm.Type_TYPE_STRING:
				return fmt.Sprintf("%s == %s", a, b)

			case orm.Type_TYPE_BYTES,
				orm.Type_TYPE_UUID:
				ident := f.QualifiedGoIdent(protogen.GoImportPath("bytes").Ident("Equal"))
				return fmt.Sprintf("%s(%s, %s)", ident, a, b)

			case orm.Type_TYPE_JSON:
				panic("JSON equal not implement")

			case orm.Type_TYPE_TIME:
				return fmt.Sprintf("%s.Equal(%s)", a, b)

			default:
				panic(fmt.Sprintf("type not supported for default value: %v", t))
			}
		},
	})

	t, err := t.ParseFS(template_files, "*")
	if err != nil {
		panic(err)
	}

	return t
}
