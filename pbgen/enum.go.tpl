enum {{ .Name }} {
{{- with .Options }}
	{{ include "options" . | indent 1 }}
{{- end }}
{{- range .Body }}
	{{ include .TemplateName . | indent 1 }}
{{- end }}
}
{{- /**/ -}}
