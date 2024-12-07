package graph

import (
	"fmt"

	"github.com/go-openapi/inflect"
	orm "github.com/lesomnus/proto-orm"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type Rpc struct {
	Entity *Entity
	Name   string

	Req *RpcMessage
	Res *RpcMessage
}

type RpcMessage struct {
	FullName protoreflect.FullName
	Stream   bool

	Entity *Entity
	Source *protogen.Message
}

type RpcOp int

const (
	RpcOpAdd RpcOp = iota
	RpcOpGet
	RpcOpList
	RpcOpPatch
	RpcOpErase
)

type RpcParser struct {
	Messages map[protoreflect.FullName]*RpcMessage
}

const (
	msg_empty protoreflect.FullName = "google.protobuf.Empty"
)

const (
	dir_req = "Request"
	dir_res = "Response"
)

func NewRpcParser() *RpcParser {
	p := &RpcParser{Messages: map[protoreflect.FullName]*RpcMessage{}}
	p.Messages[msg_empty] = &RpcMessage{FullName: msg_empty}
	return p
}

func (p *RpcParser) entityMsg(e *Entity) *RpcMessage {
	full := e.FullName()
	if m, ok := p.Messages[full]; ok {
		return m
	}

	m := &RpcMessage{
		FullName: full,
		Entity:   e,
		Source:   e.Source,
	}
	p.Messages[full] = m
	return m
}

func (*RpcParser) nameMsg(e *Entity, op string, dir string) protoreflect.FullName {
	n := fmt.Sprintf("%s%s%s", inflect.Camelize(e.Name()), op, dir)
	return e.FullName().Parent().Append(protoreflect.Name(n))
}

func (p *RpcParser) msg(e *Entity, op string, dir string) *RpcMessage {
	full := p.nameMsg(e, op, dir)
	if m, ok := p.Messages[full]; ok {
		return m
	}

	m := &RpcMessage{
		FullName: full,
		Entity:   e,
	}
	p.Messages[full] = m
	return m
}

func (w *RpcParser) rpcAdd(e *Entity) *Rpc {
	op := "Add"
	return &Rpc{
		Entity: e,
		Name:   op,

		Req: w.msg(e, op, dir_req),
		Res: w.entityMsg(e),
	}
}

func (w *RpcParser) rpcGet(e *Entity) *Rpc {
	op := "Get"
	return &Rpc{
		Entity: e,
		Name:   op,

		Req: w.msg(e, op, dir_req),
		Res: w.entityMsg(e),
	}
}

func (w *RpcParser) rpcPatch(e *Entity) *Rpc {
	op := "Patch"
	return &Rpc{
		Entity: e,
		Name:   op,

		Req: w.msg(e, op, dir_req),
		Res: w.Messages[msg_empty],
	}
}

func (w *RpcParser) rpcErase(e *Entity) *Rpc {
	return &Rpc{
		Entity: e,
		Name:   "Erase",

		Req: w.msg(e, "Get", dir_req),
		Res: w.Messages[msg_empty],
	}
}

func (w *RpcParser) Parse(e *Entity, o *orm.RpcOptions) map[RpcOp]*Rpc {
	vs := map[RpcOp]*Rpc{}
	if o.Add != nil && !o.Add.Disabled {
		vs[RpcOpAdd] = w.rpcAdd(e)
	}
	if o.Get != nil && !o.Get.Disabled {
		vs[RpcOpGet] = w.rpcGet(e)
	}
	if o.Patch != nil && !o.Patch.Disabled {
		vs[RpcOpPatch] = w.rpcPatch(e)
	}
	if o.Erase != nil && !o.Erase.Disabled {
		vs[RpcOpErase] = w.rpcErase(e)
	}

	return vs
}
