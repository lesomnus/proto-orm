package graph

type Index struct {
	name   string
	Entity *Entity // Message that defines this index.
	Refs   []Field

	Unique    bool
	Immutable bool
}

func (i *Index) Name() string {
	return i.name
}

// Split splits references into scalar fields and edges.
func (i *Index) Split() ([]*ScalarField, []*Edge) {
	ss := []*ScalarField{}
	es := []*Edge{}
	for _, r := range i.Refs {
		switch f := r.(type) {
		case (*ScalarField):
			ss = append(ss, f)
		case (*Edge):
			es = append(es, f)
		}
	}

	return ss, es
}

func (i *Index) IsUnique() bool {
	return i.Unique
}

func (i *Index) IsImmutable() bool {
	return i.Immutable
}
