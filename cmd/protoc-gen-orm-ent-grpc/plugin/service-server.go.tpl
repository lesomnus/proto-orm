{{ $server := print $.Name "ServiceServer" }}
type {{ $server }} struct {
	Db *{{ ent "Client" }}
	{{ pb (print "Unimplemented" $server) }}
}

func New{{ $server }}(db *{{ ent "Client" }}) {{ $server }} {
	return {{ $server }}{Db: db}
}
