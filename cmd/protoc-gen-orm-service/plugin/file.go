package plugin

import (
	orm "github.com/lesomnus/proto-orm"
	"github.com/lesomnus/proto-orm/graph"
	"google.golang.org/protobuf/compiler/protogen"
)

type File struct {
	*protogen.GeneratedFile
	Path       string
	Entities   []*graph.Entity
	RpcOptions orm.RpcOptions
}

func (f *File) Source() *protogen.File {
	return f.Entities[0].File
}
