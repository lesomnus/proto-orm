rpc {{ .Name }} (
{{- with .Request -}}
	{{ if .Stream }}stream {{ end -}}
	{{ .Type -}}
{{- end -}}
) returns (
{{- with .Response -}}
	{{ if .Stream }}stream {{ end -}}
	{{ .Type -}}
{{- end -}}
)
{{- if len .Options | eq 0 -}}
;
{{- else }} {
	{{ include "options" .Options | indent 1 }}
}
{{- end -}}
