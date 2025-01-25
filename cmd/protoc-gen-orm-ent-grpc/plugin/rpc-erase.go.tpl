func (s {{ $.Name }}ServiceServer) Erase(ctx {{ pkg "context" | ident "Context" }}, req *{{ pb (print $.Name "GetRequest") }}) (*{{ pkg "google.golang.org/protobuf/types/known/emptypb" | ident "Empty" }}, error) {
	p, err := {{ $.Name }}Pick(req)
	if err != nil {
		return nil, err
	}
	if _, err := s.Db.{{ $.Name }}.Delete().Where(p).Exec(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}
