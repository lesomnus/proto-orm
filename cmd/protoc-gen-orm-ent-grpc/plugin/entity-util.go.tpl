{{ $entity_pkg := entity $ }}
{{ $pred_ent := predicate $.Name -}}
{{ $req_name := pb (print $.Name "GetRequest") -}}
func {{ $.Name }}Pick(req *{{ $req_name }}) ({{ $pred_ent }}, error) {
	switch req.WhichKey() {
	{{ range .KeyLikes -}}
	{{ $n := print "req.Get" (pascal .Name) "()" -}}
	case {{ pb (print $.Name "GetRequest_" (pascal .Name) "_case") }}:
		{{/* printing field key */ -}}

		{{ if key_is_field . }}{{ with key_as_field . -}}
		{{ if is_edge . -}}
		{{/* TODO: edge key */ -}}
		{{ continue -}}
		{{ end -}}{{/* printing edge is done */ -}}

		{{ with as_attr . -}}
		{{ $t := .Type -}}
		{{ $p := $entity_pkg | ident (print (ent_pascal .Name) "EQ") -}}

		{{ if is_symmetric $t -}}
		{{/*   key is symmetric field */ -}}
		return {{ $p }}({{ $n }}), nil
		{{ else -}}
		{{/*   key is asymmetric field */ -}}
		if v, err := {{ convert_to_ent_field $n $t }}; err != nil {
			return nil, {{ grpc_errf "InvalidArgument" (print .Name ": %s" | quote) "\"err\"" }}
		} else {
			return {{ $p }}(v), nil
		}
		{{ end -}}

		{{ continue -}}
		{{ end }}{{/* attr is printed */ -}}
		{{ end }}{{ end }}{{/* field is printed */ -}}



		{{/* printing index key */ -}}

		{{ $v := key_as_index . -}}
		ps := make([]{{ $pred_ent }}, 0, {{ len $v.Refs }})
		{{ range $v.Refs }}{{/* for each refs in the index */ -}}
		{{ $ref_name := print $n ".Get" (pascal .Name) "()" -}}

		{{ if is_attr . }}{{ with as_attr . -}}
		{{/* ref is attribute */ -}}
		{{ $t := .Type -}}
		{{ $p := entity $ | ident (print (ent_pascal .Name) "EQ") -}}
		{{ if is_symmetric $t -}}
		{{/*   ref attribute is symmetric */ -}}
		ps = append(ps, {{ $p }}({{ $ref_name }}))
		{{ else -}}
		{{/*   ref attribute is asymmetric */ -}}
		if v, err := {{ convert_to_ent_field $ref_name $t }}; err != nil {
			return nil, {{ grpc_errf "InvalidArgument" (print $v.Name "." .Name ": %s" | quote) "\"err\"" }}
		} else {
			ps = append(ps, {{ $p }}(v))
		}
		{{ end -}}
		{{ continue -}}
		{{ end }}{{ end }}{{/* attribute ref is printed */ -}}

		{{ with as_edge . -}}
		{{/* ref is edge */ -}}
		if p, err := {{ .Target.Name }}Pick({{ $n }}.Get{{ pascal .Name }}()); err != nil {
			s := {{ grpc_status "Convert" }}(err)
			return nil, {{ grpc_errf "InvalidArgument" (print $v.Name "." .Name ": %s" | quote) "s.Message()" }}
		} else {
			ps = append(ps, {{ entity $ | ident (print "Has" (ent_pascal .Name) "With") }}(p))
		}
		{{ continue -}}
		{{ end -}}{{/* edge ref is printed */ -}}
		{{ end }}{{/* end of ref range */ -}}
		return {{ entity $ | ident "And" }}(ps...), nil
	{{ end -}}
	case {{ pb (print $.Name "GetRequest_Key_not_set_case") }}:
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
	if req.Has{{ $n }}() {
		{{ if is_symmetric $t -}}
		return req.Get{{ $n }}(), nil
		{{ else -}}
		if v, err := {{ convert_to_ent_field (print "req.Get" $n "()") $t }}; err != nil {
			return z, {{ grpc_errf "InvalidArgument" (print "key." $k.Name ": %s" | quote) "err" }}
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

func {{ .Name }}SelectedFields(m *{{ pb (print $.Name "Select") }}) []string {
	if m.GetAll() {
		return {{ $entity_pkg | ident "Columns" }}
	}

	vs := []string{ {{ $entity_pkg | ident "FieldID" }} }
	{{ range (slice .FieldsSortByNumber 1) -}}
	{{ with as_attr . -}}
	if m.Get{{ pascal .Name }}() {
		vs = append(vs, {{ $entity_pkg | ident (print "Field" (ent_pascal .Name)) }})
	}
	{{ end -}}
	{{ end }}
	return vs
}

func {{ .Name }}Select(q *{{ ent (print .Name "Query") }}, m *{{ pb (print $.Name "Select") }}) {
	if !m.GetAll() {
		fields := {{ .Name }}SelectedFields(m)
		q.Select(fields...)
	}
	{{ range (slice .FieldsSortByNumber 1) -}}
	{{ with as_edge . -}}
	{{ $n := pascal .Name -}}
	if m.Has{{ $n }}() {
		q.With{{ $n }}(func(q *{{ ent (print .Target.Name "Query") }}) {
			{{ print .Target.Name }}Select(q, m.Get{{ $n }}())
		})
	}
	{{ end -}}
	{{ end -}}
}
