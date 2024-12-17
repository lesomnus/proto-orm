func (e *{{ $.Name }}) Proto() *{{ pb $.Name }} {
	m := &{{ pb $.Name }}{}
	{{ range .FieldsSortByNumber -}}

	{{ if is_attr . }}{{ with as_attr . -}}
	{{ if .Nullable -}}
	{{/* TODO: maybe scalar array does not need to be dereferenced. */ -}}
	if v := e.{{ ent_pascal .Name }}; v != nil {
		m.Set{{ .GoName }}({{ ent_value_to_proto "*e" . }})
	}
	{{ else -}}
	m.Set{{ .GoName }}({{ ent_value_to_proto "e" . }})
	{{ end -}}
	{{ continue -}}
	{{ end }}{{ end -}}

	{{ if is_edge . }}{{ with as_edge . -}}
	{{ if .IsList -}}
	m.Set{{ .GoName }}(func() []*{{ pb .Target.Name }} {
		ts := e.Edges.{{ ent_pascal .Name }}
		vs := make([]*{{ pb .Target.Name }}, len(ts))
		for i, t := range ts {
			vs[i] = t.Proto()
		}
		return vs
	}())
	{{ else -}}
	if v := e.Edges.{{ ent_pascal .Name }}; v != nil {
		m.Set{{ .GoName }}(v.Proto())
	}
	{{ end -}}
	{{ continue -}}
	{{ end }}{{ end -}}

	{{ end }}

	return m
}
