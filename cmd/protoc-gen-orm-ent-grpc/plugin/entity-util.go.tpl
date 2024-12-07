{{ $pred_ent := predicate $.Name -}}
{{ $req_name := pb (print $.Name "GetRequest") -}}
func {{ $.Name }}Pick(req *{{ $req_name }}) ({{ $pred_ent }}, error) {
	switch k := req.GetKey().(type) {
	{{ range .KeyLikes -}}
	case *{{ pb (print $.Name "GetRequest_" (pascal .Name)) }}:
		{{ if key_is_field . -}}
		{{/* key is field */ -}}
		{{ $v := key_as_field . -}}
		{{ $n := pascal $v.Name -}}
		{{ $t := $v.Type -}}
		{{ $p := entity $ | ident (print (ent_pascal $v.Name) "EQ") -}}
		{{ if is_symmetric $t -}}
		{{/*   key is symmetric field */ -}}
		return {{ $p }}(k.{{ $n }}), nil
		{{ else -}}
		{{/*   key is asymmetric field */ -}}
		if v, err := {{ convert_to_ent_field (print "k." $n) $t }}; err != nil {
			return nil, {{ grpc_errf "InvalidArgument" (print $v.Name ": %s" | quote) "\"err\"" }}
		} else {
			return {{ $p }}(v), nil
		}
		{{ end -}}
		{{ else -}}
		{{/*   key is index */ -}}
		{{ $v := key_as_index . -}}
		{{ $n := pascal $v.Name -}}
		ps := make([]{{ $pred_ent }}, 0, {{ len $v.Refs }})
		{{ range $v.Refs }}{{/* for each refs in index */ -}}
		
		{{ $ref_name := pascal .Name -}}
		{{ if .IsEdge -}}
		{{/*       ref is edge */ -}}
		if p, err := {{ $ref_name }}Pick(k.{{ $n }}.Get{{ $ref_name }}()); err != nil {
			s, _ := {{ grpc_status "FromError" }}(err)
			return nil, {{ grpc_errf "InvalidArgument" (print $v.Name ".%s" | quote) "s.Message()" }}
		} else {
			ps = append(ps, {{ entity $ | ident (print "Has" (ent_pascal .Name) "With") }}(p))
		}
		{{ else -}}
		{{/*       ref is field */ -}}
		{{ $t := .Type -}}
		{{ $p := entity $ | ident (print (ent_pascal $v.Name) "EQ") -}}
		{{ if is_symmetric $t -}}
		{{/*         ref field is symmetric */ -}}
		ps = append(ps, {{ $p }}(k.{{ $ref_name }}))
		{{ else -}}
		{{/*         ref field is asymmetric */ -}}
		if v, err := {{ convert_to_ent_field (print "k." $ref_name) $t }}; err != nil {
			return nil, {{ grpc_errf "InvalidArgument" (print $v.Name "." $ref_name ": %s" | quote) "\"err\"" }}
		} else {
			ps = append(ps, {{ $p }}(v))
		}
		{{ end -}}
		{{ end -}}

		{{ end }}{{/* end of range */ -}}
		return {{ entity $ | ident "And" }}(ps...), nil
		{{ end -}}
	{{ end -}}
	case nil:
		return nil, {{ grpc_errf "InvalidArgument" "\"key not provided\"" }}
	default:
		return nil, {{ grpc_errf "Unimplemented" "\"unknown type of key\"" }}
	}
}

{{ $k := $.Key }}
{{ $t := $k.Type -}}
{{ $n := pascal $k.Name }}
func {{ .Name }}GetId(ctx {{ pkg "context" | ident "Context" }}, db *{{ ent "Client" }}, req *{{ $req_name }}) ({{ ent_type $t }}, error) {
	var z {{ ent_type $t }}
	k := req.GetKey()
	if r, ok := k.(*{{ pb (print $.Name "GetRequest_" $n) }}); ok {
		{{ if is_symmetric $t -}}
		return r.{{ $n }}
		{{ else -}}
		if v, err := {{ convert_to_ent_field (print "r." $n) $t }}; err != nil {
			return z, {{ grpc_errf "InvalidArgument" (print $n ": %s" | quote) "\"err\"" }}
		} else {
			return v, nil
		}
		{{ end -}}
	}

	p, err := {{ $.Name }}Pick(req)
	if err != nil {
		return z, err
	}

	v, err := db.{{ $.Name }}.Query().Where(p).OnlyID(ctx)
	if err != nil {
		return z, {{ grpc_errf "Internal" (quote "query: %s") "err" }}
	}

	return v, nil
}
