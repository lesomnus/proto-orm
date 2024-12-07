{{ if eq 1 (len .) }}
{{- with index . 0 -}}
[{{ .Name }} = {{ with .Value }}{{ include .TemplateName . | indent 1 }}{{ end }}]
{{- end -}}
{{- else -}}
[
{{- range $i, $_ := . }}{{ if ne $i 0 }},{{ end }}
	{{ .Name }} = {{ with .Value }}{{ include .TemplateName . | indent 1 }}{{ end }}
{{- end }}
]
{{- end -}}
