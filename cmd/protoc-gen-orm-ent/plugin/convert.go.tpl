func (e *{{ $.Name }}) Proto() *{{ pb $.Name }} {
	m := &{{ pb $.Name }}{}
	{{ range .FieldsSortByNumber -}}

	{{ if is_attr . }}{{ with as_attr . -}}
	{{ if .Nullable -}}
	{{/* TODO: maybe scalar array does not need to be dereferenced. */ -}}
	if v := e.{{ ent_pascal .Name }}; v != nil {
		m.{{ .GoName }} = {{ ent_value_to_proto "*e" . }}
	}
	{{ else -}}
	m.{{ .GoName }} = {{ ent_value_to_proto "e" . }}
	{{ end -}}
	{{ continue -}}
	{{ end }}{{ end -}}

	{{ if is_edge . }}{{ with as_edge . -}}
	{{ if .IsList -}}
	for _, v := range e.Edges.{{ ent_pascal .Name }} {
		m.{{ .GoName }} = append(m.{{ .GoName }}, v.Proto())
	}
	{{ else -}}
	if v := e.Edges.{{ ent_pascal .Name }}; v != nil {
		m.{{ .GoName }} = v.Proto()
	}
	{{ end -}}
	{{ continue -}}
	{{ end }}{{ end -}}

	{{ end }}

	return m
}
