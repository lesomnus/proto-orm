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

	{{ with as_attr . -}}
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
	{{ end }}{{/* attribute is printed */ -}}


	{{/* printing edges */ -}}

	{{ with as_edge . -}}

	{{ $target := .Target -}}
	{{ $k := $target.Key -}}
	{{ $t := $k.Type -}}
	{{ $n := pascal .Name -}}

	{{ if not .IsList -}}
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
	{{ else -}}
	for _, p := range req.Get{{ $n }}() {
		switch p.GetOp() {
		case {{ ormpb "PatchOp_ADD" }}:
		case {{ ormpb "PatchOp_ERASE" }}:
		case {{ ormpb "PatchOp_CLEAR" }}:

		case {{ ormpb "PatchOp_UNSPECIFIED" }}:
			continue
		default:
			return nil, {{ grpc_errf "Unimplemented" (quote "unknown op") }}
		}

		ids := []{{ ent_type $t }}{}
		for _, item := range p.GetItems() {
			id, err := {{ $target.Name }}GetId(ctx, s.Db, item)
			if err != nil {
				return nil, err
			}
			ids = append(ids, id)
		}

		switch p.GetOp() {
		case {{ ormpb "PatchOp_ADD" }}:
			q.Add{{ ent_singular $n }}IDs(ids...)
		case {{ ormpb "PatchOp_ERASE" }}:
			q.Remove{{ ent_singular $n }}IDs(ids...)
		case {{ ormpb "PatchOp_CLEAR" }}:
			q.Clear{{ ent_plural $n }}()
		}
	}
	{{ end -}}
	{{ end }}{{/* edge field is printed */ -}}

	{{ end }}

	if _, err := q.Save(ctx); err != nil {
		return nil, StatusFromEntError(err)
	}

	return nil, nil
}
