package plugin

import (
	"embed"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/go-openapi/inflect"
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

	import_grpc_status protogen.GoImportPath = "google.golang.org/grpc/status"
	import_grpc_codes  protogen.GoImportPath = "google.golang.org/grpc/codes"
)

var (
	//go:embed *.go.tpl
	template_files embed.FS
)

type Printer struct {
	EntPkg protogen.GoImportPath
	PbPkg  protogen.GoImportPath
}

func (p *Printer) Print(g graph.Graph, e *graph.Entity, f *protogen.GeneratedFile) error {
	t := p.newTemplate(f)
	if err := t.ExecuteTemplate(f, "service-server.go.tpl", e); err != nil {
		return err
	}
	f.P("")
	if _, ok := e.Rpcs[graph.RpcOpAdd]; ok {
		if err := t.ExecuteTemplate(f, "rpc-add.go.tpl", e); err != nil {
			return err
		}
		f.P("")
	}
	if _, ok := e.Rpcs[graph.RpcOpGet]; ok {
		if err := t.ExecuteTemplate(f, "rpc-get.go.tpl", e); err != nil {
			return err
		}
		f.P("")
	}
	if _, ok := e.Rpcs[graph.RpcOpPatch]; ok {
		if err := t.ExecuteTemplate(f, "rpc-patch.go.tpl", e); err != nil {
			return err
		}
		f.P("")
	}
	if _, ok := e.Rpcs[graph.RpcOpErase]; ok {
		if err := t.ExecuteTemplate(f, "rpc-erase.go.tpl", e); err != nil {
			return err
		}
		f.P("")
	}
	if err := t.ExecuteTemplate(f, "entity-util.go.tpl", e); err != nil {
		return err
	}

	return nil
}

func (p *Printer) newTemplate(f *protogen.GeneratedFile) *template.Template {
	t := template.New("")
	t.Funcs(template.FuncMap{
		"debug": func(v any) string {
			fmt.Fprintln(os.Stderr, v)
			return ""
		},
		"quote": func(v string) string {
			return fmt.Sprintf("%q", v)
		},
		"ent_pascal": func(v string) string {
			return rules.EntPascal(v)
		},
		"pascal": func(v string) string {
			return inflect.Camelize(v)
		},
		"plural": func(v string) string {
			return inflect.Pluralize(v)
		},

		"pkg": func(name string) protogen.GoImportPath {
			return protogen.GoImportPath(name)
		},
		"ident": func(name string, pkg protogen.GoImportPath) string {
			return f.QualifiedGoIdent(pkg.Ident(name))
		},
		"ent": func(name string) string {
			return f.QualifiedGoIdent(p.EntPkg.Ident(name))
		},
		"entity": func(e *graph.Entity) protogen.GoImportPath {
			// I think it should converted to package name according to convention.
			name := strings.ToLower(e.Name())
			return protogen.GoImportPath(path.Join(string(p.EntPkg), name))
		},
		"pb": func(name string) string {
			return f.QualifiedGoIdent(p.PbPkg.Ident(name))
		},
		"predicate": func(name string) string {
			i := protogen.GoImportPath(path.Join(string(p.EntPkg), "predicate"))
			return f.QualifiedGoIdent(i.Ident(name))
		},
		"grpc_status": func(name string) string {
			return f.QualifiedGoIdent(import_grpc_status.Ident(name))
		},
		"grpc_errf": func(code string, args ...string) string {
			s := f.QualifiedGoIdent(import_grpc_status.Ident("Errorf"))
			c := f.QualifiedGoIdent(import_grpc_codes.Ident(code))
			a := strings.Join(args, ", ")
			return fmt.Sprintf("%s(%s, %s)", s, c, a)
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

		"convert_to_ent_field": func(ident string, t orm.Type) string {
			switch t {
			case orm.Type_TYPE_UUID:
				c := f.QualifiedGoIdent(import_uuid.Ident("FromBytes"))
				return fmt.Sprintf("%s(%s)", c, ident)

			default:
				panic("convert not supported")
			}
		},
		"ent_type": func(t orm.Type) string {
			switch t {
			case orm.Type_TYPE_UNSPECIFIED:
			case orm.Type_TYPE_BOOL:
				return "bool"
			case orm.Type_TYPE_ENUM:
			case orm.Type_TYPE_INT32:
				fallthrough
			case orm.Type_TYPE_SINT32:
				fallthrough
			case orm.Type_TYPE_SFIXED32:
				return "int32"
			case orm.Type_TYPE_UINT32:
				fallthrough
			case orm.Type_TYPE_FIXED32:
				return "uint32"
			case orm.Type_TYPE_INT64:
				fallthrough
			case orm.Type_TYPE_SINT64:
				fallthrough
			case orm.Type_TYPE_SFIXED64:
				return "int64"
			case orm.Type_TYPE_UINT64:
				fallthrough
			case orm.Type_TYPE_FIXED64:
				return "uint64"
			case orm.Type_TYPE_FLOAT:
				return "float32"
			case orm.Type_TYPE_DOUBLE:
				return "float"
			case orm.Type_TYPE_STRING:
				return "string"
			case orm.Type_TYPE_BYTES:
				return "[]byte"
			case orm.Type_TYPE_MESSAGE:
			case orm.Type_TYPE_GROUP:
			case orm.Type_TYPE_UUID:
				return f.QualifiedGoIdent(import_uuid.Ident("UUID"))
			case orm.Type_TYPE_TIME:
				return f.QualifiedGoIdent(import_pb_ts.Ident("Timestamp"))
			}
			panic("type must be scalar mapped")
		},

		// Test if `t` is same for Proto and Ent.
		"is_symmetric": func(t orm.Type) bool {
			switch t {
			case orm.Type_TYPE_UNSPECIFIED,
				orm.Type_TYPE_MESSAGE,
				orm.Type_TYPE_GROUP,
				orm.Type_TYPE_UUID:
				return false

			default:
				return true
			}
		},
		"to_symmetric_ent": func(ident string, t orm.Type) string {
			i, _ := strings.CutPrefix(ident, "*")
			switch t {
			case orm.Type_TYPE_ENUM:
				return fmt.Sprintf("int(%s)", ident)
			case orm.Type_TYPE_TIME:
				return fmt.Sprintf("%s.AsTime()", i)
			case orm.Type_TYPE_MESSAGE,
				orm.Type_TYPE_JSON:
				return i
			default:
				return ident
			}
		},
	})

	t, err := t.ParseFS(template_files, "*")
	if err != nil {
		panic(err)
	}

	return t
}
