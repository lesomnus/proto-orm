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

// XxxGetRequest has `select` field before fields from entity starts:
//
//	{
//		XxxSelect select = 1
//		oneof key {
//			bytes id = 2 // <- note that it starts from 2.
//		}
//	}
const msgGetEntityFieldOffset = 1

type Printer struct {
	// `p` should be `/path/to/file.proto` and
	// it will returns something like `/path/to/file_svc.prot`.
	namer func(p string) string

	messages map[protoreflect.FullName]*generatedMessage
}

func NewPrinter() *Printer {
	return &Printer{}
}

func (p *Printer) Print(f *File) error {
	if len(f.Entities) == 0 {
		return os.ErrNotExist
	}

	f.P(`// Code generated by "protoc-gen-orm-service", DO NOT EDIT.`)
	f.P("")

	pf := &pbgen.ProtoFile{
		Path:    f.Path,
		Edition: pbgen.SyntaxProto3,
		Imports: []pbgen.Import{
			{Name: f.Entities[0].File.Desc.Path()},
		},
	}
	if d := f.Source().Proto; d != nil {
		pf.Package = protoreflect.FullName(*d.Package)
		pf.Options = []pbgen.Option{
			pbgen.OptionGoPackage(*d.Options.GoPackage),
		}
	}

	w := printWork{
		Printer:  p,
		file:     pf,
		services: map[protoreflect.FullName]*pbgen.Service{},
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

			// XxxGetRequest
			// - contains XxxSelect.
			// - may contains XxxGetByXxx.
			if strings.HasSuffix(string(m.FullName), "GetRequest") {
				msg_select := m.Body[0].(pbgen.MessageField)
				if _, ok := ms[msg_select.Type]; !ok {
					ms[msg_select.Type] = true
					if m, ok := w.messages[protoreflect.FullName(msg_select.Type)]; !ok {
						panic("XxxSelect not generated")
					} else {
						pf.TopLevelDefinitions = append(pf.TopLevelDefinitions, m)
					}
				}

				msg_key := m.Body[msgGetEntityFieldOffset].(pbgen.MessageOneof)
				for _, e := range msg_key.Body {
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

	slices.SortFunc(pf.Imports, func(a, b pbgen.Import) int {
		return strings.Compare(a.Name, b.Name)
	})
	return pbgen.PrintFile(f, pf)
}

type printWork struct {
	*Printer
	file     *pbgen.ProtoFile
	services map[protoreflect.FullName]*pbgen.Service
}

func (w *printWork) newMessage(m *graph.RpcMessage) (*generatedMessage, bool) {
	n := m.FullName
	if v, ok := w.messages[n]; ok {
		return v, false
	}

	v := &generatedMessage{
		Message: &pbgen.Message{FullName: n},
		p:       w.namer(m.Entity.File.Desc.Path()),
	}
	w.messages[n] = v
	return v, true
}

func (w *printWork) importType(t graph.ProtoTyped) {
	if p := t.ImportPath(); p != "" {
		w.file.AddImport(pbgen.Import{Name: p})
	}
}

func (w *printWork) typeMessage(t graph.ProtoTyped) pbgen.Type {
	w.importType(t)
	return pbgen.Type(t.ProtoType())
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
	if r, ok := e.Rpcs[graph.RpcOpAdd]; ok {
		s.Body = append(s.Body, w.rpcAdd(r))
	}
	if r, ok := e.Rpcs[graph.RpcOpGet]; ok {
		s.Body = append(s.Body, w.rpcGet(r))
	}
	if r, ok := e.Rpcs[graph.RpcOpPatch]; ok {
		s.Body = append(s.Body, w.rpcPatch(r))
	}
	if r, ok := e.Rpcs[graph.RpcOpErase]; ok {
		s.Body = append(s.Body, w.rpcErase(r))
	}

	w.services[e.FullName()] = s
}

func (w *printWork) msgAddReq(r *graph.Rpc) *generatedMessage {
	m, ok := w.newMessage(r.Req)
	if !ok {
		return m
	}

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

			v.Type = w.typeMessage(u)

			if u.IsMap() {
				// Map cannot have label.
			} else if u.IsList() {
				v.Label = pbgen.LabelRepeated
			} else if u.HasOptionalKeyword() || u.HasDefault() {
				v.Label = pbgen.LabelOptional
			}

		case (*graph.Edge):
			if u.IsVirtual() {
				continue
			}
			if !(u.IsUnidirectional() || u.IsExclusive()) {
				continue
			}

			target := u.Target
			r := target.Rpcs[graph.RpcOpGet]
			if r == nil {
				panic("TODO: target entity does not enables Get rpc")
				// To resolve this, graph should parse the disabled rpcs also
				// or make one on demand.
			}
			m := w.msgGetReq(r)
			v.Type = w.typeMessage(m)

			if u.IsList() {
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

func (w *printWork) indexGet(e *graph.Entity, i *graph.Index) *generatedMessage {
	full := w.nameIndexGet(e, i)
	if m, ok := w.messages[full]; ok {
		return m
	}

	m := &generatedMessage{
		Message: &pbgen.Message{FullName: full},
		p:       w.namer(e.File.Desc.Path()),
	}
	w.messages[full] = m

	body := []pbgen.MessageBody{}
	for _, r := range i.Refs {
		switch u := r.(type) {
		case (*graph.Attr):
			v := pbgen.MessageField{
				Type:   w.typeMessage(r),
				Name:   protoreflect.Name(r.Name()),
				Number: int(r.Number()),
			}
			body = append(body, v)

		case (*graph.Edge):
			m := w.msgGetReq(u.Target.Rpcs[graph.RpcOpGet])
			v := pbgen.MessageField{
				Type:   w.typeMessage(m),
				Name:   protoreflect.Name(r.Name()),
				Number: int(r.Number()),
			}
			body = append(body, v)
		}
	}

	m.Body = body
	return m
}

func (w *printWork) nameMsgSelect(e *graph.Entity) protoreflect.FullName {
	n := fmt.Sprintf("%sSelect", inflect.Camelize(e.Name()))
	return e.FullName().Parent().Append(protoreflect.Name(n))
}

func (w *printWork) msgSelect(e *graph.Entity) *generatedMessage {
	full := w.nameMsgSelect(e)
	if m, ok := w.messages[full]; ok {
		w.importType(m)
		return m
	}

	m := &generatedMessage{
		Message: &pbgen.Message{FullName: full},
		p:       w.namer(e.File.Desc.Path()),
	}
	w.messages[full] = m

	body := []pbgen.MessageBody{
		pbgen.MessageField{
			Type:   pbgen.TypeBool,
			Name:   "all",
			Number: 1,
			Label:  pbgen.LabelOptional,
		},
	}
	for _, r := range e.FieldsSortByNumber()[1:] {
		if r.IsIgnored() || r.IsVirtual() {
			continue
		}

		switch u := r.(type) {
		case (*graph.Attr):
			v := pbgen.MessageField{
				Type:   pbgen.TypeBool,
				Name:   protoreflect.Name(u.Name()),
				Number: int(u.Number()),
				Label:  pbgen.LabelOptional,
			}
			body = append(body, v)

		case (*graph.Edge):
			m := w.msgSelect(u.Target)
			v := pbgen.MessageField{
				Type:   w.typeMessage(m),
				Name:   protoreflect.Name(u.Name()),
				Number: int(u.Number()),
				Label:  pbgen.LabelOptional,
			}
			body = append(body, v)
		}
	}

	m.Body = body
	return m
}

func (w *printWork) msgGetReq(r *graph.Rpc) *generatedMessage {
	m, ok := w.newMessage(r.Req)
	if !ok {
		return m
	}

	msg_select := w.msgSelect(r.Entity)
	field_select := pbgen.MessageField{
		Type:   w.typeMessage(msg_select),
		Name:   "select",
		Number: 1,
	}

	field_key := pbgen.MessageOneof{Name: "key"}
	for _, k := range r.Entity.KeyLikes() {
		switch v := k.(type) {
		case *graph.Attr:
			field_key.Body = append(field_key.Body, pbgen.MessageOneofField{
				Type:   w.typeMessage(v),
				Name:   protoreflect.Name(v.Name()),
				Number: int(v.Number()) + msgGetEntityFieldOffset,
			})

		case *graph.Edge:
			// TODO:
			panic("todo")

		case *graph.Index:
			m := w.indexGet(r.Entity, v)
			field_key.Body = append(field_key.Body, pbgen.MessageOneofField{
				Type:   w.typeMessage(m),
				Name:   protoreflect.Name(v.Name()),
				Number: int(v.Refs[0].Number()) + msgGetEntityFieldOffset,
			})
		}
	}

	m.Body = []pbgen.MessageBody{
		field_select,
		field_key,
	}
	return m
}

func (w *printWork) rpcGet(r *graph.Rpc) pbgen.Rpc {
	w.msgGetReq(r)
	return w.rpc(r)
}

func (w *printWork) msgPatchReq(r *graph.Rpc) *generatedMessage {
	m, ok := w.newMessage(r.Req)
	if !ok {
		return m
	}

	k := w.msgGetReq(r.Entity.Rpcs[graph.RpcOpGet])
	body := []pbgen.MessageBody{pbgen.MessageField{
		Type:   w.typeMessage(k),
		Name:   "key",
		Number: 1,
	}}

	for _, f := range r.Entity.FieldsSortByNumber() {
		if f.IsImmutable() {
			// Key must be immutable.
			continue
		}

		n := int(f.Number())*2 - 1
		v := pbgen.MessageField{
			Name:   protoreflect.Name(f.Name()),
			Number: n,
		}
		switch u := f.(type) {
		case (*graph.Attr):
			if u.IsBound() {
				// Skip since edge (which is accessed by `f.Bound`) will be added.
				continue
			}

			v.Type = w.typeMessage(u)

			if u.IsMap() {
				// Map cannot have label.
			} else if u.IsList() {
				v.Label = pbgen.LabelRepeated
			} else if !f.IsMap() {
				v.Label = pbgen.LabelOptional
			}

		case (*graph.Edge):
			if u.IsVirtual() {
				continue
			}
			if !(u.IsUnidirectional() || u.IsExclusive()) {
				continue
			}

			target := u.Target
			r := target.Rpcs[graph.RpcOpGet]
			if r == nil {
				panic("TODO: target entity does not enables Get rpc")
				// To resolve this, graph should parse the disabled rpcs also
				// or make one on demand.
			}
			m := w.msgGetReq(r)
			v.Type = w.typeMessage(m)

			if u.IsList() {
				v.Label = pbgen.LabelRepeated
			}
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
