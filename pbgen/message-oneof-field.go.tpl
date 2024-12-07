{{ .Type }} {{ .Name }} = {{ .Number -}}
{{- with .Options }} {{ template "field-options.go.tpl" . -}}{{- end -}}
;
{{- /**/ -}}
