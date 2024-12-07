[
{{- range $i, $_ := .Values }}
	{{- if ne $i 0 }},{{ end }}
	{{ include .TemplateName . | indent 1 }}
{{- end}}
]
{{- /**/ -}}
