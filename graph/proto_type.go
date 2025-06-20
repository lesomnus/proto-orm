package graph

type ProtoTyped interface {
	// Returns valid name of type in proto.
	// e.g. string, map<string, string>, google.protobuf.Timestamp, etc...
	ProtoType() string

	// Path to proto file where this type is declared in.
	// Returns empty string if this type is primitive one.
	ImportPath() string
}

type protoType struct {
	name string
	path string
}

func (t protoType) ProtoType() string {
	return t.name
}

func (t protoType) ImportPath() string {
	return t.path
}

func ProtoType(name, path string) ProtoTyped {
	return protoType{name, path}
}
