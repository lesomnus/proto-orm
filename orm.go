package orm

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
)

func ResolveRpcOptions(f *protogen.File, m *protogen.Message) *RpcOptions {
	o := &RpcOptions{}
	if o_ := proto.GetExtension(f.Desc.Options(), E_All).(*FileOptions); o_ != nil && o_.Rpcs != nil && !o_.Rpcs.Disabled {
		o.Add = o_.Rpcs.Add
		o.Get = o_.Rpcs.Get
		o.List = o_.Rpcs.List
		o.Patch = o_.Rpcs.Patch
		o.Erase = o_.Rpcs.Erase
	}
	if o_ := proto.GetExtension(m.Desc.Options(), E_Rpcs).(*RpcOptions); o_ != nil && o_.Disabled != nil {
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
