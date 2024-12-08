package plugin

import (
	"fmt"

	orm "github.com/lesomnus/proto-orm"
)

func toEntIdent(t orm.Type) string {
	v := ""
	switch t {
	case orm.Type_TYPE_BOOL:
		v = "Bool"
	case orm.Type_TYPE_ENUM:
		panic("todo")
	case orm.Type_TYPE_INT32:
		fallthrough
	case orm.Type_TYPE_SINT32:
		fallthrough
	case orm.Type_TYPE_SFIXED32:
		v = "Int32"
	case orm.Type_TYPE_UINT32:
		fallthrough
	case orm.Type_TYPE_FIXED32:
		v = "Uint32"
	case orm.Type_TYPE_INT64:
		fallthrough
	case orm.Type_TYPE_SINT64:
		fallthrough
	case orm.Type_TYPE_SFIXED64:
		v = "Int64"
	case orm.Type_TYPE_UINT64:
		fallthrough
	case orm.Type_TYPE_FIXED64:
		v = "Uint64"
	case orm.Type_TYPE_FLOAT:
		v = "Float32"
	case orm.Type_TYPE_DOUBLE:
		v = "Float"
	case orm.Type_TYPE_STRING:
		v = "String"
	case orm.Type_TYPE_BYTES:
		v = "Bytes"
	// case orm.Type_TYPE_MESSAGE:
	// case orm.Type_TYPE_GROUP:
	case orm.Type_TYPE_UUID:
		v = "UUID"
	case orm.Type_TYPE_TIME:
		v = "Time"

	default:
		panic(fmt.Sprintf("unknown type or type not supported: %s", t.String()))
	}

	return v
}
