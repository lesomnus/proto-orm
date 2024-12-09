{{ $server := print $.Name "ServiceServer" }}
type {{ $server }} struct {
	db *{{ ent "Client" }}
	{{ pb (print "Unimplemented" $server) }}
}

func New{{ $server }}(db *{{ ent "Client" }}) *{{ $server }} {
	return &{{ $server }}{db: db}
}
