func (s *{{ $.Name }}ServiceServer) Get(ctx {{ pkg "context" | ident "Context" }}, req *{{ pb (print $.Name "GetRequest") }}) (*{{ pb $.Name }}, error) {
	q := s.db.{{ $.Name }}.Query()
	if p, err := {{ $.Name }}Pick(req); err != nil {
		return nil, err
	} else {
		q.Where(p)
	}

	{{ range .FieldsSortByNumber -}}
	{{ with as_edge . -}}
	{{ if not (.IsUnidirectional | or .IsExclusive) -}}
		{{ continue -}}
	{{ end -}}
	{{ $n := ent_pascal .Name -}}
	q.With{{ $n }}(func(q *{{ ent (print .Target.Name "Query") }}) {
		q.Select({{ entity .Target | ident "FieldID" }})
	})
	{{ end -}}
	{{ end }}

	v, err := q.Only(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}
