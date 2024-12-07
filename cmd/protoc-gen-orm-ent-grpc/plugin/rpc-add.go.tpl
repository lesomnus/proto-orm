func (s *{{ .Name }}ServiceServer) Add(ctx {{ pkg "context" | ident "Context" }}, req *{{ pb (print .Name "AddRequest") }}) (*{{ pb .Name }}, error) {
	q := s.db.{{ .Name }}.Create()
	{{ range .FieldsSortByNumber -}}
	{{ if .IsEdge -}}
	{{/* field is edge */ -}}
	{{ $target := .Edge.Target -}}
	{{ $k := $target.Key -}}
	{{ $n := $target.Name -}}
	{{ $v := key_as_field $k -}}
	{{ if .Edge.HasMany -}}
	{{/*     has many edges */ -}}
	{
		ids := []{{ ent_type $v.Type }}{}
		for _, r := range req.Get{{ plural $n }}() {
			{{/* TODO: optimize: get multiple in single query */ -}}
			if id, err := {{ $n }}GetId(ctx, s.db, r); err != nil {
				return nil, err
			} else {
				ids = append(ids, id)
			}
		}
		q.Add{{ $n }}IDs(ids...)
	}
	{{ else -}}
	{{/*     has single edge */ -}}
	if id, err := {{ $n }}GetId(ctx, s.db, req.Get{{ $n }}()); err != nil {
		return nil, err
	} else {
		q.Set{{ $n }}ID(id)
	}
	{{ end -}}
	{{ else -}}
	{{/* field is not an edge */ -}}
	{{ if .IsBound -}}
		{{/* skip: field is bounded one */ -}}
		{{ continue -}}
	{{ end -}}
	{{ $t := .Type -}}
	{{ $v := print "req." (pascal .Name) -}}
	{{ $n := ent_pascal .Name -}}
	{{ if is_symmetric $t -}}
	{{/*   field is symmetric */ -}}
	q.Set{{ $n }}({{ to_symmetric_ent $v $t }})
	{{ else -}}
	{{/*   field is asymmetric */ -}}
	if v, err := {{ convert_to_ent_field $v $t }}; err != nil {
		return nil, {{ grpc_errf "InvalidArgument" (print .Name ": %s" | quote) "err" }}
	} else {
		q.Set{{ $n }}(v)
	}
	{{ end -}}
	{{ end -}}
	{{ end }}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}
