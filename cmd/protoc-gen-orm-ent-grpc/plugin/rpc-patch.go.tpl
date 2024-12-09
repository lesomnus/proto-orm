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


	{{/* printing scalar field */ -}}

	{{ if is_scalar . }}{{ with as_scalar . -}}
	{{ if .IsBound -}}
		{{/* skip: field is bounded one */ -}}
		{{ continue -}}
	{{ end -}}

	{{ $t := .Type -}}
	{{ $v := print "req." (pascal .Name) -}}
	{{ $n := ent_pascal .Name -}}

	{{ if .IsNullable -}}
	if {{ $v }}Null {
		q.Clear{{ $n }}()
	} else {{ end -}}

	if {{ $v }} != nil {
		{{ if $t.IsPrimitive -}}
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
	{{ end }}{{ end }}{{/* scalar field is printed */ -}}


	{{/* printing edge field */ -}}

	{{ with as_edge . -}}
	{{ if or .IsList .HasInverse -}}
		{{/* skip: patch for list is not supported */ -}}
		{{/* skip: ownership (which is defined by .Inverse) should be patched by subject (at .Reverse) if it is bidirectional. */ -}}
		{{ continue -}}
	{{ end -}}

	{{ $target := .Target -}}
	{{ $k := $target.Key -}}
	{{ $v := print "req." (pascal .Name) -}}
	{{ $n := ent_pascal $target.Name -}}

	{{ if .IsNullable -}}
	if {{ $v }}Null {
		q.Clear{{ $n }}()
	} else {{ end -}}

	if {{ $v }} != nil {
		q.Set{{ $n }}($v)
	}
	{{ continue -}}
	{{ end }}{{/* edge field is printed */ -}}

	{{ end }}

	v, err := q.Save(ctx)
	if err != nil {
		return nil, StatusFromEntError(err)
	}

	return v.Proto(), nil
}
