package plugin

import "github.com/lesomnus/proto-orm/pbgen"

type generatedMessage struct {
	*pbgen.Message
	p string
}

func (m *generatedMessage) ProtoType() string {
	return string(m.Message.FullName)
}

func (m *generatedMessage) ImportPath() string {
	return m.p
}
