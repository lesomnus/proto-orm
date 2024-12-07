{{ $server := print .Name "ServiceServer" }}
type {{ $server }} struct {
	db *{{ ent "Client" }}
	{{ pb (print "Unimplemented" .Name "ServiceServer") }}
}

func New{{ $server }}(db *{{ ent "Client" }}) *{{ $server }} {
	return &{{ $server }}{db: db}
}
