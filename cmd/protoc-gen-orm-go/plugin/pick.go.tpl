{{ $r := pb (print $.Name "GetRequest" ) -}}
func (x *{{ $.Name }}) Pick() *{{ $r }} {
	return {{ $.Name }}ById(x.GetId())
}

func (x *{{ $r }}) Picks(v *{{ $.Name }}) bool {
	switch x.WhichKey() {
	{{ range .KeyLikes -}}
	{{ $n := print "req.Get" (pascal .Name) "()" -}}
	case {{ pb (print $.Name "GetRequest_" (pascal .Name) "_case") }}:
		{{ with key_as_field . -}}
		{{ if is_edge . -}}
		{{/* TODO: edge key */ -}}
		{{ continue -}}
		{{ end -}}{{/* printing edge is done */ -}}

		{{ with as_attr . -}}
		{{ $getter := print "Get" (pascal .Name) "()" -}}
		return {{ def_attr_eq . (print "x." $getter) (print "v." $getter) }}

		{{ continue -}}
		{{ end }}{{/* attr is printed */ -}}
		{{ end }}{{/* field is printed */ -}}
		


		{{/* printing index key */ -}}
		return true {{ with key_as_index . -}}
		{{ $index := . -}}
		{{ range .Refs -}}
		
		{{ $x_get := print "x.Get" (pascal $index.Name) "().Get" (pascal .Name) "()" -}}
		{{ $v_get := print "v.Get" (pascal .Name) "()" -}}
		{{ with as_attr . }} &&
			{{ def_attr_eq . $x_get $v_get -}}
		{{ continue -}}
		{{ end }}{{/* attr is printed */ -}}

		{{ with as_edge . -}} &&
			{{ print $x_get ".Picks(" $v_get ")" -}}
		{{ continue -}}
		{{ end }}{{/* edge is printed */ -}}

		{{ end }}{{/* attribute ref is printed */ -}}
		{{ end }}{{/* printing index is done */}}
	{{ end -}}

	default:
		return false
	}
}
