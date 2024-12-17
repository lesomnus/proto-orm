func (s *{{ $.Name }}ServiceServer) Patch(ctx {{ pkg "context" | ident "Context" }}, req *{{ pb (print $.Name "PatchRequest") }}) (*{{ pkg "google.golang.org/protobuf/types/known/emptypb" | ident "Empty" }}, error) {
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


	{{/* printing attributes */ -}}

	{{ if is_attr . }}{{ with as_attr . -}}
	{{ if .IsBound -}}
		{{/* skip: field is bounded one */ -}}
		{{ continue -}}
	{{ end -}}

	{{ $t := .Type -}}
	{{ $v := print "req.Get" (pascal .Name) "()" -}}
	{{ $n := ent_pascal .Name -}}

	{{ if .IsNullable -}}
	if {{ print "req.Get" (pascal .Name) "Null()" }} {
		q.Clear{{ $n }}()
	} else {{ end -}}

	{{ if .IsSoft -}}
	if v := {{ $v }}; v != nil {
		q.Set{{ $n }}(v)
	}
	{{ else -}}
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
	{{ end -}}
	
	{{ continue -}}
	{{ end }}{{ end }}{{/* attribute is printed */ -}}


	{{/* printing edges */ -}}

	{{ with as_edge . -}}
	{{ if not (or .IsUnidirectional .IsExclusive) -}}
		{{ continue -}}
	{{ end -}}

	{{ $target := .Target -}}
	{{ $k := $target.Key -}}
	{{ $t := $k.Type -}}
	{{ $v := print "req." (pascal .Name) -}}
	{{ $n := ent_pascal $target.Name -}}

	{{ if .IsNullable -}}
	if {{ print "req.Get" (pascal .Name) "Null()" }} {
		q.Clear{{ pascal .Name }}()
	} else {{ end -}}

	if id, err := {{ $.Name }}GetId(ctx, s.db, req.GetKey()); err != nil {
		return nil, err
	} else {
		q.Set{{ pascal .Name }}ID(id)
	}
	{{ continue -}}
	{{ end }}{{/* edge field is printed */ -}}

	{{ end }}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}
