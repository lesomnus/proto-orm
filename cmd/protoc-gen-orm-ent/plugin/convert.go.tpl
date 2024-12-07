func (e *{{ .Name }}) Proto() *{{ pb .Name }} {
	m := &{{ pb .Name }}{}
	{{ range .FieldsSortByNumber -}}
	{{ if .IsEdge -}}
	{{ if .IsList -}}
	for _, v := range e.Edges.{{ ent_name . }} {
		m.{{ .GoName }} = append(m.{{ .GoName }}, v.Proto())
	}
	{{ else -}}
	if v := e.Edges.{{ ent_name . }}; v != nil {
		m.{{ .GoName }} = v.Proto()
	}
	{{ end -}}
	{{ else if not .Required -}}
	if v := e.{{ ent_name . }}; v != nil {
		m.{{ .GoName }} = {{ ent_value_to_proto "*e" . }}
	}
	{{ else -}}
	m.{{ .GoName }} = {{ ent_value_to_proto "e" . }}
	{{ end -}}
	{{ end }}

	return m
}
