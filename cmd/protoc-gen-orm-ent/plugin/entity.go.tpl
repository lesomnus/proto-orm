type {{ $.Name }} struct {
	{{ ent "Schema" }}
}

{{ if .HasScalarFields -}}
func ({{ $.Name }}) Fields() []{{ ent "Field" }} {
	return []{{ ent "Field" }} {
		{{ range $.FieldsSortByNumber -}}
		{{ if .Ignored -}}
			{{ continue -}}
		{{ end -}}
		{{ if .IsEdge -}}
			{{ continue -}}
		{{ end -}}
		{{ def_field . -}}
		{{ if .HasDefault -}}.
		Default({{ def_field_default . }})
		{{- end -}}
		{{ if not .Required | and (not .IsList) -}}.
		Optional()
		{{- end -}}
		{{ if .Unique -}}.
		Unique()
		{{- end -}}
		{{ if .Immutable -}}.
		Immutable()
		{{- end -}}
		{{ if .Nullable -}}.
		Nillable()
		{{- end -}},
		{{ end }}
	}
}
{{ end }}

{{ if .HasEdges -}}
func ({{ $.Name }}) Edges() []{{ ent "Edge" }} {
	return []{{ ent "Edge" }}{
		{{ range $.FieldsSortByNumber -}}
			{{ if .Ignored -}}
				{{ continue -}}
			{{ end -}}
			{{ if not .IsEdge -}}
				{{ continue -}}
			{{ end -}}
			{{ def_edge .Edge -}}
			{{ if .Edge.HasBind -}}.
			Field({{ .Edge.Bind.Name | printf "%q" }})
			{{- end -}}
			{{ if not .Edge.HasMany -}}.
			Unique()
			{{- end -}}
			{{ if .Required -}}.
			Required()
			{{- end -}}
			{{ if .Immutable -}}.
			Immutable()
			{{- end -}},
		{{ end }}
	}
}
{{ end }}

{{ if .HasIndexes -}}
func ({{ $.Name }}) Indexes() []{{ ent "Index" }} {
	return []{{ ent "Index" }}{
		{{ range $.Indexes -}}
			{{ def_index . -}}
			{{ if .Unique -}}.
			Unique()
			{{- end -}},
		{{ end }}
	}
}
{{ end }}
