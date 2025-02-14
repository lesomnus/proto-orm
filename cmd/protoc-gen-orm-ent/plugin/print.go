package plugin

import (
	"embed"
	"fmt"
	"os"
	"strings"
	"text/template"

	orm "github.com/lesomnus/proto-orm"
	"github.com/lesomnus/proto-orm/graph"
	"github.com/lesomnus/proto-orm/internal/rules"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	import_ent       protogen.GoImportPath = "entgo.io/ent"
	import_ent_field protogen.GoImportPath = "entgo.io/ent/schema/field"
	import_ent_edge  protogen.GoImportPath = "entgo.io/ent/schema/edge"
	import_ent_index protogen.GoImportPath = "entgo.io/ent/schema/index"
	import_uuid      protogen.GoImportPath = "github.com/google/uuid"
	import_pb_ts     protogen.GoImportPath = "google.golang.org/protobuf/types/known/timestamppb"
)

var (
	//go:embed *.go.tpl
	template_files embed.FS
)

func (p *Plugin) Print(e *graph.Entity, f *protogen.GeneratedFile) error {
	if err := p.NewTemplate(e, f).ExecuteTemplate(f, "entity.go.tpl", e); err != nil {
		return err
	}
	if err := p.NewTemplate(e, p.EntConvertFile).ExecuteTemplate(p.EntConvertFile, "convert.go.tpl", e); err != nil {
		return err
	} else {
		p.EntConvertFile.P("")
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
		"ent_pascal": func(v string) string {
			return rules.EntPascal(v)
		},

		"ent": func(name string) string {
			return f.QualifiedGoIdent(import_ent.Ident(name))
		},
		"pb": func(name string) string {
			return f.QualifiedGoIdent(e.File.GoImportPath.Ident(name))
		},

		"is_attr": func(f graph.Field) bool {
			_, ok := f.(*graph.Attr)
			return ok
		},
		"as_attr": func(f graph.Field) *graph.Attr {
			v, ok := f.(*graph.Attr)
			if !ok {
				panic("field must be an attribute")
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
				panic("field must be an edge")
			}
			return v
		},

		"def_field": func(field *graph.Attr) string {
			t := field.Type()
			n := field.Name()
			v := ""

			i := toEntIdent(t)
			if field.IsMap() {
				v = field.GoType(f.QualifiedGoIdent) + "{}"
			} else if field.IsList() {
				i = "JSON"
				v = field.GoType(f.QualifiedGoIdent) + "{}"
			} else {
				switch t {
				case orm.Type_TYPE_ENUM:
					i = "Int"
				case orm.Type_TYPE_UUID:
					v = f.QualifiedGoIdent(import_uuid.Ident("UUID")) + "{}"
				case orm.Type_TYPE_JSON:
					v = "&" + field.GoType(f.QualifiedGoIdent) + "{}"
				}
			}

			ident := f.QualifiedGoIdent(import_ent_field.Ident(i))
			if len(v) > 0 {
				return fmt.Sprintf(`%s(%q, %s)`, ident, n, v)
			} else {
				return fmt.Sprintf(`%s(%q)`, ident, n)
			}
		},
		"def_field_default": func(field *graph.Attr) string {
			if !field.HasDefault() {
				panic("it must have a default")
			}

			t := field.Type()
			switch t {
			case orm.Type_TYPE_BOOL:
				return "false"
			case orm.Type_TYPE_ENUM:
				return "0"
			case orm.Type_TYPE_INT32,
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
				orm.Type_TYPE_DOUBLE:
				return "0"
			case orm.Type_TYPE_STRING:
				return `""`
			case orm.Type_TYPE_BYTES:
				return "[]byte{}"

			case orm.Type_TYPE_JSON:
				v := field.GoType(f.QualifiedGoIdent) + "{}"
				if field.Kind() == protoreflect.MessageKind && !field.IsMap() {
					v = "&" + v
				}
				return v
			case orm.Type_TYPE_UUID:
				return f.QualifiedGoIdent(import_uuid.Ident("New"))
			case orm.Type_TYPE_TIME:
				return f.QualifiedGoIdent(protogen.GoImportPath("time").Ident("Now"))
			default:
				panic(fmt.Sprintf("type not supported for default value: %v", t))
			}
		},
		"def_edge": func(edge *graph.Edge) string {
			name := edge.Name()
			if edge.Inverse != nil {
				return fmt.Sprintf(
					"%s(%q, %s.Type).Ref(%q)",
					f.QualifiedGoIdent(import_ent_edge.Ident("From")),
					name,
					edge.Target.Source.GoIdent.GoName,
					edge.Inverse.Name(),
				)
			}

			v := fmt.Sprintf(
				`%s(%q, %s.Type)`,
				f.QualifiedGoIdent(import_ent_edge.Ident("To")),
				name,
				edge.Target.Source.GoIdent.GoName,
			)
			if edge.IsSelfLoop() {
				v += fmt.Sprintf(".From(%q)", edge.Inverse.Name())
			}
			return v
		},
		"def_index": func(index *graph.Index) string {
			fs, es := index.Split()
			f_keys := []string{}
			for _, f := range fs {
				f_keys = append(f_keys, fmt.Sprintf(`%q`, string(f.Name())))
			}
			e_keys := []string{}
			for _, e := range es {
				e_keys = append(e_keys, fmt.Sprintf(`%q`, string(e.Name())))
			}

			ident := "Fields"
			has_fields := len(fs) > 0
			has_edges := len(es) > 0
			if !has_fields {
				ident = "Edges"
			}

			v := f.QualifiedGoIdent(import_ent_index.Ident(ident))
			if has_fields {
				v += fmt.Sprintf("(%s)", strings.Join(f_keys, ", "))
			}
			if has_fields && has_edges {
				v += ".Edges"
			}
			if has_edges {
				v += fmt.Sprintf("(%s)", strings.Join(e_keys, ", "))
			}

			return v
		},
		"ent_type": func(field *graph.Attr) string {
			return field.GoType(f.QualifiedGoIdent)
		},
		"ent_value_to_proto": func(ident string, field *graph.Attr) string {
			n := fmt.Sprintf("%s.%s", ident, rules.EntPascal(string(field.Name())))
			if field.IsList() {
				return n
			}

			t := field.Type()
			switch t {
			case orm.Type_TYPE_ENUM:
				c := field.GoType(f.QualifiedGoIdent)
				return fmt.Sprintf("%s(%s)", c, n)
			case orm.Type_TYPE_UUID:
				return fmt.Sprintf("%s[:]", n)
			case orm.Type_TYPE_TIME:
				c := f.QualifiedGoIdent(import_pb_ts.Ident("New"))
				return fmt.Sprintf("%s(%s)", c, n)
			default:
				return n
			}
		},
	})

	t, err := t.ParseFS(template_files, "*")
	if err != nil {
		panic(err)
	}

	return t
}
