{{ .Name }} = {{ .Number -}}
{{- if .Options }} {{ template "field-options.go.tpl" .Options -}}{{- end -}}
;
{{- /**/ -}}
