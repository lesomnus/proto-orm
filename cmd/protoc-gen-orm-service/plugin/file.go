package plugin

import (
	"github.com/lesomnus/proto-orm/graph"
	orm "github.com/lesomnus/proto-orm"
	"google.golang.org/protobuf/compiler/protogen"
)

type File struct {
	*protogen.GeneratedFile
	Entities   []*graph.Entity
	RpcOptions orm.RpcOptions
}

func (f *File) Source() *protogen.File {
	return f.Entities[0].File
}
