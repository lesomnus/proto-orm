func (s *{{ $.Name }}ServiceServer) Patch(ctx {{ pkg "context" | ident "Context" }}, req *{{ pb (print $.Name "PatchRequest") }}) (*{{ pb $.Name }}, error) {
	id, err := {{ $.Name }}GetId(ctx, s.db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.db.{{ $.Name }}.UpdateOneID(id)
	{{ range .FieldsSortByNumber -}}
	{{ if .Immutable -}}
		{{/* skip: field cannot be patched */ -}}
		{{ continue -}}
	{{ end -}}
	{{ if .IsEdge -}}
		{{/* skip: edge is edited by Add or Erase */ -}}
		{{ continue -}}
	{{ end -}}
	{{ if .IsBound -}}
		{{/* skip: field is bound one */ -}}
		{{ continue -}}
	{{ end -}}

	{{ $t := .Type -}}
	{{ $v := print "req." (pascal .Name) -}}
	{{ $n := ent_pascal .Name -}}
	{{ if .Nullable -}}
	if {{ $v }}Null {
		q.Clear{{ $n }}()
	} else if {{ $v }} != nil {
		{{ if .IsScalar -}}
		q.Set{{ $n }}(*{{ $v }})
		{{ else if is_symmetric $t -}}
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

	if {{ $v }} != nil {
		{{ if .IsScalar -}}
		q.Set{{ $n }}(*{{ $v }})
		{{ else if is_symmetric $t -}}
		q.Set{{ $n }}({{ to_symmetric_ent $v $t }})
		{{ else -}}
		if v, err := {{ convert_to_ent_field $v $t }}; err != nil {
			return nil, {{ grpc_errf "InvalidArgument" (print .Name ": %s" | quote) "err" }}
		} else {
			q.Set{{ $n }}(v)
		}
		{{ end -}}
	}
	{{/* end of range */ -}}
	{{ end }}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}
