package orm

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func ResolveRpcOptions(f *protogen.File, m *protogen.Message) *RpcOptions {
	o := &RpcOptions{}
	if o_ := proto.GetExtension(f.Desc.Options(), E_All).(*FileOptions); o_ != nil && o_.Rpcs != nil && !o_.Rpcs.Disabled {
		o.Crud = &o_.Rpcs.Crud
		o.Add = o_.Rpcs.Add
		o.Get = o_.Rpcs.Get
		o.List = o_.Rpcs.List
		o.Patch = o_.Rpcs.Patch
		o.Erase = o_.Rpcs.Erase
	}
	if o_ := proto.GetExtension(m.Desc.Options(), E_Rpcs).(*RpcOptions); o_ != nil && o_.Disabled != nil {
		o.Crud = o_.Crud
		if *o_.Disabled {
			o = nil
		} else {
			if o_.Add != nil && o_.Add.Disabled {
				o.Add = nil
			}
			if o_.Get != nil && o_.Get.Disabled {
				o.Get = nil
			}
			if o_.List != nil && o_.List.Disabled {
				o.List = nil
			}
			if o_.Patch != nil && o_.Patch.Disabled {
				o.Patch = nil
			}
			if o_.Erase != nil && o_.Erase.Disabled {
				o.Erase = nil
			}
		}
	}

	return o
}

func (t Type) IsPrimitive() bool {
	switch t {
	case Type_TYPE_BOOL:
	case Type_TYPE_INT32:
	case Type_TYPE_SINT32:
	case Type_TYPE_UINT32:
	case Type_TYPE_INT64:
	case Type_TYPE_SINT64:
	case Type_TYPE_UINT64:
	case Type_TYPE_SFIXED32:
	case Type_TYPE_FIXED32:
	case Type_TYPE_FLOAT:
	case Type_TYPE_SFIXED64:
	case Type_TYPE_FIXED64:
	case Type_TYPE_DOUBLE:
	case Type_TYPE_STRING:
	case Type_TYPE_BYTES:

	default:
		return false
	}

	return true
}

// Decay convert the type into less specific type.
// int <- int32, int64, ...
// uint <- uint32, uint64, ...
// float <- float, double
// X <- X (rest)
func (t Type) Decay() Type {
	switch t {
	case Type_TYPE_INT32:
		fallthrough
	case Type_TYPE_INT64:
		fallthrough
	case Type_TYPE_SINT32:
		fallthrough
	case Type_TYPE_SINT64:
		fallthrough
	case Type_TYPE_SFIXED32:
		fallthrough
	case Type_TYPE_SFIXED64:
		return Type_TYPE_INT

	case Type_TYPE_UINT32:
		fallthrough
	case Type_TYPE_UINT64:
		fallthrough
	case Type_TYPE_FIXED32:
		fallthrough
	case Type_TYPE_FIXED64:
		return Type_TYPE_UINT

	case Type_TYPE_FLOAT:
		fallthrough
	case Type_TYPE_DOUBLE:
		return Type_TYPE_FLOAT
	}

	return t
}
