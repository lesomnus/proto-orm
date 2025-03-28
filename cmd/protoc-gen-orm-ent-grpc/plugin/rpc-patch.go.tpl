func (s {{ $.Name }}ServiceServer) Patch(ctx {{ pkg "context" | ident "Context" }}, req *{{ pb (print $.Name "PatchRequest") }}) (*{{ pkg "google.golang.org/protobuf/types/known/emptypb" | ident "Empty" }}, error) {
	id, err := {{ $.Name }}GetId(ctx, s.Db, req.GetKey())
	if err != nil {
		return nil, err
	}

	q := s.Db.{{ $.Name }}.UpdateOneID(id)
	{{ range .FieldsSortByNumber -}}
	{{ if .IsVirtual -}}
		{{ continue -}}
	{{ end -}}
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
	{{ $n := pascal .Name -}}
	{{ $v := print "req." $n -}}

	{{ if .IsNullable -}}
	if req.Get{{ $n }}Null() {
		q.Clear{{ $n }}()
	} else {{ end -}}
	if req.Has{{ $n }}() {
		if id, err := {{ $target.Name  }}GetId(ctx, s.Db, req.Get{{ $n }}()); err != nil {
			return nil, err
		} else {
			q.Set{{ $n }}ID(id)
		}
	}
	{{ continue -}}
	{{ end }}{{/* edge field is printed */ -}}

	{{ end }}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}
