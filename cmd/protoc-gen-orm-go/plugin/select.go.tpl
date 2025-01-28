{{ $r := pb (print $.Name "GetRequest" ) -}}
func (x *{{ $r }}) Select(v *{{ print $.Name "Select" }}) *{{ $r }} {
	x.SetSelect(v)
	return x
}
