{{ with .Edition }}{{ .Keyword }} = "{{ .Value }}"{{ end }};

{{- with .Package }}

package {{ . }};
{{- end }}

{{- range $i, $_ := .Imports }}{{ if eq $i 0 }}
{{ end }}
import {{ with .Visibility }}{{ . }} {{ end }}"{{ .Name }}";
{{- end }}

{{- with .Options }}

{{ template "options.go.tpl" . }}
{{- end }}


{{- $empty_line := true }}
{{- range .TopLevelDefinitions }}{{ if $empty_line }}
{{ end }}
{{- $empty_line = ne .TemplateName "comment" }}
{{ include .TemplateName . }}
{{- end }}
