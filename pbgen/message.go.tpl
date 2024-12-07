message {{ .FullName.Name }} {
{{- range .Body }}
	{{ include .TemplateName . | indent 1 }}
{{- end }}
}
{{- /**/ -}}
