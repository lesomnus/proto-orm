func (s *{{ $.Name }}ServiceServer) Get(ctx {{ pkg "context" | ident "Context" }}, req *{{ pb (print $.Name "GetRequest") }}) (*{{ pb $.Name }}, error) {
	q := s.db.{{ $.Name }}.Query()
	if p, err := {{ $.Name }}Pick(req); err != nil {
		return nil, err
	} else {
		q.Where(p)
	}

	v, err := q.Only(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}
