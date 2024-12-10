package plugin

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strings"

	"github.com/go-openapi/inflect"
	orm "github.com/lesomnus/proto-orm"
	"github.com/lesomnus/proto-orm/graph"
	"github.com/lesomnus/proto-orm/pbgen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Printer struct {
}

func NewPrinter() *Printer {
	return &Printer{}
}

func (*Printer) Print(f *File) error {
	if len(f.Entities) == 0 {
		return os.ErrNotExist
	}

	f.P(`// Code generated by "proto-orm-gen-ent". DO NOT EDIT`)
	f.P("")

	pf := &pbgen.ProtoFile{
		Edition: pbgen.SyntaxProto3,
		Imports: []pbgen.Import{
			{Name: f.Entities[0].File.Desc.Path()},
		},
	}
	if d := f.Source().Proto; d != nil {
		pf.Package = protoreflect.FullName(*d.Package)
		for _, s := range d.Dependency {
			if s == "orm.proto" {
				continue
			}
			pf.AddImport(pbgen.Import{Name: s})
		}
		pf.Options = []pbgen.Option{
			pbgen.OptionGoPackage(*d.Options.GoPackage),
		}
	}

	w := printWork{
		file:     pf,
		services: map[protoreflect.FullName]*pbgen.Service{},
		messages: map[protoreflect.FullName]*pbgen.Message{},
	}
	for _, e := range f.Entities {
		o := orm.ResolveRpcOptions(f.Source(), e.Source)
		if o == nil {
			continue
		}

		w.Do(e, o)
	}

	ss := slices.Collect(maps.Values(w.services))
	slices.SortFunc(ss, func(a *pbgen.Service, b *pbgen.Service) int {
		return strings.Compare(string(a.Name), string(b.Name))
	})
	for _, v := range ss {
		pf.TopLevelDefinitions = append(pf.TopLevelDefinitions, v)
	}

	// Processed messages.
	ms := map[pbgen.Type]bool{}

	// Print RPC messages for each services.
	for _, s := range ss {
		for _, e := range s.Body {
			r, ok := e.(pbgen.Rpc)
			if !ok {
				continue
			}
			if _, ok := ms[r.Request.Type]; ok {
				continue
			} else {
				ms[r.Request.Type] = true
			}

			m, ok := w.messages[protoreflect.FullName(r.Request.Type)]
			if !ok {
				panic(fmt.Sprintf("message must be exist: %s", r.Request.Type))
			} else {
				pf.TopLevelDefinitions = append(pf.TopLevelDefinitions, m)
			}

			// XXXGetRequest may contains XXXGetByXXX.
			if strings.HasSuffix(string(m.FullName), "GetRequest") {
				o := m.Body[0].(pbgen.MessageOneof)
				for _, e := range o.Body {
					f, ok := e.(pbgen.MessageOneofField)
					if !ok {
						continue
					}
					if _, ok := ms[f.Type]; ok {
						continue
					} else {
						ms[f.Type] = true
					}

					m, ok := w.messages[protoreflect.FullName(f.Type)]
					if !ok {
						// `f.Type` maybe a primitive types: int or string.
						continue
					} else {
						pf.TopLevelDefinitions = append(pf.TopLevelDefinitions, m)
					}
				}
			}

			// Note that messages for response are defined in the another files:
			// e.g. entity messages, well known messages.
		}
	}

	return pbgen.PrintFile(f, pf)
}

type printWork struct {
	file     *pbgen.ProtoFile
	services map[protoreflect.FullName]*pbgen.Service
	messages map[protoreflect.FullName]*pbgen.Message
}

func (w *printWork) rpc(r *graph.Rpc) pbgen.Rpc {
	return pbgen.Rpc{
		Name: protoreflect.Name(r.Name),
		Request: pbgen.RpcType{
			Type: pbgen.Type(r.Req.FullName),
		},
		Response: pbgen.RpcType{
			Type: pbgen.Type(r.Res.FullName),
		},
	}
}

func (w *printWork) Do(e *graph.Entity, o *orm.RpcOptions) {
	s := &pbgen.Service{
		Name: protoreflect.Name(fmt.Sprintf("%sService", e.Name())),
	}
	if o.Add != nil && !o.Add.Disabled {
		r := e.Rpcs[graph.RpcOpAdd]
		s.Body = append(s.Body, w.rpcAdd(r))
	}
	if o.Get != nil && !o.Get.Disabled {
		r := e.Rpcs[graph.RpcOpGet]
		s.Body = append(s.Body, w.rpcGet(r))
	}
	// if o.List != nil && !o.List.Disabled {
	// }
	if o.Patch != nil && !o.Patch.Disabled {
		r := e.Rpcs[graph.RpcOpPatch]
		s.Body = append(s.Body, w.rpcPatch(r))
	}
	if o.Erase != nil && !o.Erase.Disabled {
		r := e.Rpcs[graph.RpcOpErase]
		s.Body = append(s.Body, w.rpcErase(r))
	}

	w.services[e.FullName()] = s
}

func (w *printWork) msgAddReq(r *graph.Rpc) *pbgen.Message {
	full := r.Req.FullName
	if m, ok := w.messages[full]; ok {
		return m
	}

	m := &pbgen.Message{FullName: full}
	w.messages[full] = m

	body := []pbgen.MessageBody{}
	for _, f := range r.Entity.FieldsSortByNumber() {
		v := pbgen.MessageField{
			Name:   protoreflect.Name(f.Name()),
			Number: int(f.Number()),
		}

		switch u := f.(type) {
		case (*graph.Attr):
			if u.IsBound() {
				// Skip since edge (which is accessed by `f.Bound`) will be added.
				continue
			}

			v.Type = pbgen.Type(u.ProtoType())

			d := u.Source().Desc
			if d.IsMap() {
				// Map cannot have label.
			} else if d.IsList() {
				v.Label = pbgen.LabelRepeated
			} else if d.HasOptionalKeyword() || u.HasDefault() {
				v.Label = pbgen.LabelOptional
			}

		case (*graph.Edge):
			if !u.HasInverse() {
				// TODO: Should the set of User[pet] be supported at the time of User creation?
				continue
			}

			target := u.Target
			r := target.Rpcs[graph.RpcOpGet]
			if r == nil {
				panic("TODO: target entity does not enables Get rpc")
				// To resolve this, graph should parse disabled rpcs also
				// or make one on demand.
			}
			m := w.msgGetReq(r)
			v.Type = pbgen.Type(m.FullName)

			d := u.Source().Desc
			if d.IsList() {
				v.Label = pbgen.LabelRepeated
			}
		}

		body = append(body, v)
	}

	m.Body = body
	return m
}

func (w *printWork) rpcAdd(r *graph.Rpc) pbgen.Rpc {
	w.msgAddReq(r)
	return w.rpc(r)
}

func (w *printWork) nameIndexGet(e *graph.Entity, i *graph.Index) protoreflect.FullName {
	n := fmt.Sprintf("%sGetBy%s", inflect.Camelize(e.Name()), inflect.Camelize(string(i.Name())))
	return e.FullName().Parent().Append(protoreflect.Name(n))
}

func (w *printWork) indexGet(e *graph.Entity, i *graph.Index) *pbgen.Message {
	full := w.nameIndexGet(e, i)
	if m, ok := w.messages[full]; ok {
		return m
	}

	m := &pbgen.Message{FullName: full}
	w.messages[full] = m

	body := []pbgen.MessageBody{}
	for _, r := range i.Refs {
		d := r.Source().Desc
		switch u := r.(type) {
		case (*graph.Attr):
			v := pbgen.MessageField{
				Name:   d.Name(),
				Type:   pbgen.Type(d.Kind().String()),
				Number: int(d.Number()),
			}
			body = append(body, v)

		case (*graph.Edge):
			m := w.msgGetReq(u.Target.Rpcs[graph.RpcOpGet])
			v := pbgen.MessageField{
				Type:   pbgen.Type(m.FullName),
				Name:   protoreflect.Name(d.Name()),
				Number: int(d.Number()),
			}
			body = append(body, v)
		}
	}

	m.Body = body
	return m
}

func (w *printWork) msgGetReq(r *graph.Rpc) *pbgen.Message {
	full := r.Req.FullName
	if m, ok := w.messages[full]; ok {
		return m
	}

	m := &pbgen.Message{FullName: full}
	w.messages[full] = m

	oneof := pbgen.MessageOneof{Name: "key"}
	for _, k := range r.Entity.KeyLikes() {
		switch v := k.(type) {
		case *graph.Attr:
			oneof.Body = append(oneof.Body, pbgen.MessageOneofField{
				Type:   pbgen.Type(v.ProtoType()),
				Name:   protoreflect.Name(v.Name()),
				Number: int(v.Number()),
			})

		case *graph.Edge:
			// TODO:
			continue
		}
	}
	for _, i := range r.Entity.Indexes {
		if !i.Unique {
			continue
		}

		m := w.indexGet(r.Entity, i)
		oneof.Body = append(oneof.Body, pbgen.MessageOneofField{
			Type:   pbgen.Type(m.FullName),
			Name:   protoreflect.Name(i.Name()),
			Number: int(i.Refs[0].Number()),
		})
	}

	m.Body = []pbgen.MessageBody{oneof}
	return m
}

func (w *printWork) rpcGet(r *graph.Rpc) pbgen.Rpc {
	w.msgGetReq(r)
	return w.rpc(r)
}

func (w *printWork) msgPatchReq(r *graph.Rpc) *pbgen.Message {
	full := r.Req.FullName
	if m, ok := w.messages[full]; ok {
		return m
	}

	m := &pbgen.Message{FullName: full}
	w.messages[full] = m

	k := w.msgGetReq(r.Entity.Rpcs[graph.RpcOpGet])
	k_f := k.Body[0].(pbgen.MessageOneof).Body[0].(pbgen.MessageOneofField)
	body := []pbgen.MessageBody{pbgen.MessageField{
		Type:   pbgen.Type(k.FullName),
		Name:   "key",
		Number: k_f.Number*2 - 1,
	}}

	for _, f := range r.Entity.FieldsSortByNumber() {
		if f.IsImmutable() {
			// Key must be immutable
			continue
		}

		n := int(f.Number())*2 - 1
		if n == body[0].(pbgen.MessageField).Number {
			continue
		}

		v := pbgen.MessageField{
			Type:   pbgen.Type(f.ProtoType()),
			Name:   protoreflect.Name(f.Name()),
			Number: n,
		}
		if f.IsList() {
			v.Label = pbgen.LabelRepeated
		} else if !f.Source().Desc.IsMap() {
			v.Label = pbgen.LabelOptional
		}

		body = append(body, v)
		if f.IsNullable() {
			body = append(body, pbgen.MessageField{
				Type:   pbgen.TypeBool,
				Name:   protoreflect.Name(fmt.Sprintf("%s_null", f.Name())),
				Number: n + 1,
			})
		}
	}

	m.Body = body
	return m
}

func (w *printWork) rpcPatch(r *graph.Rpc) pbgen.Rpc {
	w.msgPatchReq(r)
	w.file.AddImport(pbgen.Import{
		Name: pbgen.ImportWellKnownEmpty,
	})
	return w.rpc(r)
}

func (w *printWork) rpcErase(r *graph.Rpc) pbgen.Rpc {
	w.msgGetReq(r)
	w.file.AddImport(pbgen.Import{
		Name: pbgen.ImportWellKnownEmpty,
	})
	return w.rpc(r)
}
