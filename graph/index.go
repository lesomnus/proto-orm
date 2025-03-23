package graph

type Index struct {
	name   string
	Entity *Entity // Message that defines this index.
	Refs   []Field

	Unique    bool
	Immutable bool
	Hidden    bool
}

func (i *Index) Name() string {
	return i.name
}

// Split splits references into attributes and edges.
func (i *Index) Split() ([]*Attr, []*Edge) {
	ss := []*Attr{}
	es := []*Edge{}
	for _, r := range i.Refs {
		switch f := r.(type) {
		case (*Attr):
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

func (i *Index) IsHidden() bool {
	return i.Hidden
}
