type {{ $.Name }} struct {
	{{ ent "Schema" }}
}

{{ if len $.Attrs | lt 0 -}}
func ({{ $.Name }}) Fields() []{{ ent "Field" }} {
	return []{{ ent "Field" }} {
		{{ range $.AttrsSortByNumber -}}
		{{ if .Ignored -}}
			{{ continue -}}
		{{ end -}}
		
		{{ def_field . -}}
		{{ if .HasDefault -}}.
		Default({{ def_field_default . }})
		{{- end -}}
		{{ if .IsOptional | and (not .HasDefault) -}}.
		Optional()
		{{- end -}}
		{{ if .IsUnique -}}.
		Unique()
		{{- end -}}
		{{ if .IsImmutable -}}.
		Immutable()
		{{- end -}}
		{{ if .IsNullable -}}.
		Nillable()
		{{- end -}},
		{{ end }}
	}
}
{{ end }}

{{ if len $.Edges | lt 0 -}}
func ({{ $.Name }}) Edges() []{{ ent "Edge" }} {
	return []{{ ent "Edge" }}{
		{{ range $.EdgesSortByNumber -}}
			{{ if .Ignored -}}
				{{ continue -}}
			{{ end -}}

			{{ def_edge . -}}
			{{ if .HasBind -}}.
			Field({{ .Bind.Name | printf "%q" }})
			{{- end -}}
			{{ if not .IsList -}}.
			Unique()
			{{- end -}}
			{{ if .IsRequired -}}.
			Required()
			{{- end -}}
			{{ if .IsImmutable -}}.
			Immutable()
			{{- end -}},
		{{ end }}
	}
}
{{ end }}

{{ if len $.Indexes | lt 0 -}}
func ({{ $.Name }}) Indexes() []{{ ent "Index" }} {
	return []{{ ent "Index" }}{
		{{ range $.Indexes -}}
			{{ def_index . -}}
			{{ if .IsUnique -}}.
			Unique()
			{{- end -}},
		{{ end }}
	}
}
{{ end }}
