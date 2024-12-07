package pbgen

import (
	"fmt"

	"google.golang.org/protobuf/reflect/protoreflect"
)

func OptionGoPackage(v string) Option {
	o := UnsafeOptionLiteral("go_package", fmt.Sprintf(`%q`, v))
	o.Standard = true
	return o
}

var (
	FeatureEnumTypeClosed = Option{Name: "features.enum_type", Value: &UnsafeLiteral{Value: "CLOSED"}, Standard: true}
	FeatureEnumTypeOpen   = Option{Name: "features.enum_type", Value: &UnsafeLiteral{Value: "OPEN"}, Standard: true}

	FeatureFieldPresenceLegacyRequired = Option{Name: "features.field_presence", Value: &UnsafeLiteral{Value: "LEGACY_REQUIRED"}, Standard: true}
	FeatureFieldPresenceExplicit       = Option{Name: "features.field_presence", Value: &UnsafeLiteral{Value: "EXPLICIT"}, Standard: true}
	FeatureFieldPresenceImplicit       = Option{Name: "features.field_presence", Value: &UnsafeLiteral{Value: "IMPLICIT"}, Standard: true}
)

func UnsafeOptionLiteral(name protoreflect.FullName, v string) Option {
	return Option{
		Name:  name,
		Value: &UnsafeLiteral{Value: v},
	}
}
