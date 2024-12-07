{
{{- range .Fields }}
	{{ .Name }}: {{ with .Value }}{{ include .TemplateName . | indent 1 }}{{ end }}
{{- end }}
}
{{- /**/ -}}
