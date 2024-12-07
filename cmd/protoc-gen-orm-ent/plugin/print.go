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
		"ent": func(name string) string {
			return f.QualifiedGoIdent(import_ent.Ident(name))
		},
		"pb": func(name string) string {
			return f.QualifiedGoIdent(e.File.GoImportPath.Ident(name))
		},
		"def_field": func(field *graph.Field) string {
			if field.IsEdge() {
				panic("it must be a scalar field")
			}

			t := field.Type()
			ident := f.QualifiedGoIdent(import_ent_field.Ident(toEntIdent(t)))
			name := field.Name()
			if t != orm.Type_TYPE_UUID {
				return fmt.Sprintf(`%s(%q)`, ident, name)
			} else {
				v := f.QualifiedGoIdent(import_uuid.Ident("New"))
				return fmt.Sprintf(`%s(%q, %s())`, ident, name, v)
			}
		},
		"def_field_default": func(field *graph.Field) string {
			if !field.HasDefault() {
				panic("it must have a default")
			}

			switch field.Type() {
			case orm.Type_TYPE_UUID:
				return f.QualifiedGoIdent(import_uuid.Ident("New"))
			case orm.Type_TYPE_TIME:
				return f.QualifiedGoIdent(protogen.GoImportPath("time").Ident("Now"))
			default:
				panic("TODO: HasDefault() should prevent this")
			}
		},
		"def_edge": func(edge *graph.Edge) string {
			source_name := edge.Source.Name()
			if edge.From != nil {
				return fmt.Sprintf(
					"%s(%q, %s.Type).Ref(%q)",
					f.QualifiedGoIdent(import_ent_edge.Ident("From")),
					source_name,
					edge.Target.Source.GoIdent.GoName,
					edge.From.Name(),
				)
			}

			v := fmt.Sprintf(
				`%s(%q, %s.Type)`,
				f.QualifiedGoIdent(import_ent_edge.Ident("To")),
				source_name,
				edge.Target.Source.GoIdent.GoName,
			)
			if edge.IsSelfLoop() {
				v += fmt.Sprintf(".From(%q)", edge.From.Name())
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
		"ent_name": func(field *graph.Field) string {
			return rules.EntPascal(field.Name())
		},
		"ent_value_to_proto": func(ident string, field *graph.Field) string {
			t := field.Type()
			n := rules.EntPascal(field.Name())
			if field.IsEdge() {
				panic("is it needed? or handled in template?")
				return fmt.Sprintf("%s.Proto()", ident)
			}
			switch t {
			case orm.Type_TYPE_UUID:
				return fmt.Sprintf("%s.%s[:]", ident, n)
			case orm.Type_TYPE_TIME:
				c := f.QualifiedGoIdent(import_pb_ts.Ident("New"))
				return fmt.Sprintf("%s(%s.%s)", c, ident, n)
			default:
				return fmt.Sprintf("%s.%s", ident, n)
			}
		},
	})

	t, err := t.ParseFS(template_files, "*")
	if err != nil {
		panic(err)
	}

	return t
}
