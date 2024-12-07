package graph

type Index struct {
	name   string
	Entity *Entity // Message that defines this index.
	Refs   []*Field

	Key       bool
	Unique    bool
	Immutable bool
}

func (i *Index) Name() string {
	return i.name
}

// Split splits references into scalar fields and edges.
func (i *Index) Split() ([]*Field, []*Field) {
	fs := []*Field{}
	es := []*Field{}
	for _, r := range i.Refs {
		if r.IsEdge() {
			es = append(es, r)
		} else {
			fs = append(fs, r)
		}
	}

	return fs, es
}
