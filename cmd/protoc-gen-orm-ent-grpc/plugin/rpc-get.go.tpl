func (s {{ $.Name }}ServiceServer) Get(ctx {{ pkg "context" | ident "Context" }}, req *{{ pb (print $.Name "GetRequest") }}) (*{{ pb $.Name }}, error) {
	q := s.Db.{{ $.Name }}.Query()
	if req.HasSelect() {
		{{ $.Name }}Select(q, req.GetSelect())
	} else {
		{{ $.Name }}SelectEdges(q)
	}

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
