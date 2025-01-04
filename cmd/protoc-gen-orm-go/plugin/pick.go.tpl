{{ $r := pb (print $.Name "GetRequest" ) -}}
func (x *{{ $.Name }}) Pick() *{{ $r }} {
	return {{ $.Name }}ById(x.GetId())
}
