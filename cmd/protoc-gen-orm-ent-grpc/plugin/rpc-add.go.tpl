func (s {{ $.Name }}ServiceServer) Add(ctx {{ pkg "context" | ident "Context" }}, req *{{ pb (print $.Name "AddRequest") }}) (*{{ pb $.Name }}, error) {
	q := s.Db.{{ $.Name }}.Create()
	{{ range .FieldsSortByNumber -}}
	{{ if .IsVirtual -}}
		{{ continue -}}
	{{ end -}}

	{{/* printing attributes */ -}}

	{{ if is_attr . }}{{ with as_attr . -}}
	{{ if .IsBound -}}
		{{/* skip: attribute is bounded one */ -}}
		{{ continue -}}
	{{ end -}}

	{{ $t := .Type -}}
	{{ $n := ent_pascal .Name -}}
	{{ $v := print "req.Get" (pascal .Name) "()" -}}

	{{ if .IsSoft -}}
	if v := {{ $v }}; v != nil {
		q.Set{{ $n }}(v)
	}
	{{ continue -}}
	{{ end -}}

	{{ if .IsNullable | or .HasDefault -}}
	if {{ print "req.Has" (pascal .Name) "()" }} {
		{{ if is_symmetric $t -}}
		q.Set{{ $n }}({{ to_symmetric_ent $v $t }})
		{{ else -}}
		if v, err := {{ convert_to_ent_field $v $t }}; err != nil {
			return nil, {{ grpc_errf "InvalidArgument" (print .Name ": %s" | quote) "err" }}
		} else {
			q.Set{{ $n }}(v)
		}
		{{ end -}}
	}
	{{ continue -}}
	{{ end -}}
	
	{{ if is_symmetric $t -}}
	{{/*   type is symmetric */ -}}
	q.Set{{ $n }}({{ to_symmetric_ent $v $t }})
	{{ else -}}
	{{/*   type is asymmetric */ -}}
	if v, err := {{ convert_to_ent_field $v $t }}; err != nil {
		return nil, {{ grpc_errf "InvalidArgument" (print .Name ": %s" | quote) "err" }}
	} else {
		q.Set{{ $n }}(v)
	}
	{{ end -}}
	{{ continue -}}
	{{ end }}{{ end }}{{/* attribute is printed */ -}}


	{{/* printing edges */ -}}

	{{ with as_edge . -}}
	{{ if not (or .IsUnidirectional .IsExclusive) -}}
		{{ continue -}}
	{{ end -}}

	{{ $target := .Target -}}
	{{ $n := pascal .Name -}}
	{{ $k := $target.Key -}}

	{{ if .IsList -}}
	{
		ids := []{{ ent_type $k.Type }}{}
		for _, r := range req.Get{{ $n }}() {
			{{/* TODO: optimize: get multiple in single query */ -}}
			if id, err := {{ $target.Name }}GetId(ctx, s.Db, r); err != nil {
				return nil, err
			} else {
				ids = append(ids, id)
			}
		}
		q.Add{{ $n }}IDs(ids...)
	}
	{{ else if .IsOptional | or .IsNullable -}}
	if v := req.Get{{ $n }}(); v != nil {
		if id, err := {{ $target.Name }}GetId(ctx, s.Db, v); err != nil {
			return nil, err
		} else {
			q.Set{{ $n }}ID(id)
		}
	}
	{{ else -}}
	if id, err := {{ $target.Name }}GetId(ctx, s.Db, req.Get{{ $n }}()); err != nil {
		s := {{ grpc_status "Convert" }}(err)
		s = {{ grpc_status "Newf" }}(s.Code(), "%s: %s", "{{ .Name }}", s.Message())
		return nil, s.Err()
	} else {
		q.Set{{ $n }}ID(id)
	}
	{{ end -}}
	{{ continue -}}
	{{ end }}{{/* edge field is printed */ -}}

	{{ end }}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}
