{{ range $i, $_ := split "\n" .Value }}{{ if ne $i 0 }}
{{ end -}}// {{ . }}
{{- end -}}
