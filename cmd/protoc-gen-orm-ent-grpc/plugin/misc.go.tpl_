func StatusFromEntError(err error) error {
	switch {
	case {{ pkg "entgo.io/ent/dialect/sql/sqlgraph" | ident "IsUniqueConstraintError" }}(err):
		return {{ grpc_errf "AlreadyExists" (quote "already exists: %s") "err" }}
	case {{ ent "IsConstraintError" }}(err):
		return {{ grpc_errf "InvalidArgument" (quote "invalid arguments: %s") "err" }}
	case {{ ent "IsNotFound" }}(err):
		return {{ grpc_errf "NotFound" (quote "not found: %s") "err" }}
	default:
		return {{ grpc_errf "Internal" (quote "ent: %s") "err" }}
	}
}
