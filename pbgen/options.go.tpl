{{ range $i, $_ := . }}{{ if ne $i 0 }}
{{ end -}}
option {{ .NameString }} = {{ with .Value }}{{ include .TemplateName . }};{{ end -}}
{{- end -}}
