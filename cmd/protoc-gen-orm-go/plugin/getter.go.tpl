{{ $r := pb (print $.Name "GetRequest" ) -}}
{{ $b := pb (print $.Name "GetRequest_builder" ) -}}
{{ range .KeyLikes -}}

{{ with key_as_field . -}}

{{ if is_edge . -}}
{{/* TODO: edge key */ -}}
{{ continue -}}
{{ end -}}

{{ with as_attr . }}

func {{ $.Name }}By{{ pascal .Name }}(v {{ proto_go_type . }}) *{{ $r }} {
	return {{ $b }}{ {{ pascal .Name }}: {{ if .IsScalar }}&{{ end }}v }.Build()
}
{{ end -}}

{{ end -}}
{{ with key_as_index . }}

func {{ $.Name }}By{{ pascal .Name }}({{ range .Refs -}}

	{{ with as_attr . -}}
	{{ .Name }} {{ proto_go_type . }},
	{{- end -}}

	{{ with as_edge . -}}
	{{ .Name }} *{{ pascal .Target.Name }}GetRequest,
	{{- end -}}

	{{ end -}}
) *{{ $r }} {
	v := &{{ pb (print $.Name "GetBy" (pascal .Name) ) }}{}
	{{ range .Refs -}}
	v.Set{{ pascal .Name }}({{ .Name }})
	{{ end -}}
	return {{ $b }}{ {{ pascal .Name }}: v }.Build()
}
{{ end -}}

{{ end -}}
