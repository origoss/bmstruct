package bmstruct

import (
	"reflect"
)

//Template is a set of fields and is similar to the struct data structure. The
//Fields of a Template may overlap which results in a C union like data
//structure.
//
//Besides the Fields, a Template specifies its size too.
type Template struct {
	Fields map[string]*Field `json:"fields"`
	Size   int               `json:"size"`
}

//NewTemplate creates a new Template object. It checks the validity of size and
//whether the given fields fit into the given size. If the size parameter is
//less than 0, NewTemplate will calculate the size based on the given fields.
//
//NewTemplate panics when the given size is too small or when no fields were
//specified.
//
//It is valid to specify a larger Template size than the fields require.
func NewTemplate(size int, fields ...*Field) *Template {
	if len(fields) == 0 {
		panic("at least 1 field shall be specified")
	}
	t := &Template{
		Fields: make(map[string]*Field),
		Size:   size,
	}
	for _, field := range fields {
		t.Fields[field.Name] = field
	}
	if size < 0 {
		t.Size = int(t.minLen())
	} else if t.minLen() > uint64(size) {
		panic("Template size too small")
	}
	return t
}

func (t *Template) minLen() uint64 {
	l := uint64(0)
	for _, field := range t.Fields {
		if fMax := field.Offset + field.Len; fMax > l {
			l = fMax
		}
	}
	return l
}

//Equal method compares the Template to another template. It returns true if the
//Size parameter is true and if the Fields are the same.
//
//Note, that it returns true only when the order of the Fields is the same in
//the other Template.
func (t *Template) Equal(other *Template) bool {
	return reflect.DeepEqual(t, other) || reflect.DeepEqual(*t, *other)
}

//Field method turns the Template object into a Field object. This can be used
//to define hierarchical templates, i.e. a Template that contains another
//Template.
func (t *Template) Field(name string, offset uint64) *Field {
	return &Field{
		Name:   name,
		Offset: offset,
		Len:    uint64(t.Size),
	}
}

//FieldAt method returns the Field that is to be found at the given offset. The
//method panics if the offset is invalid.
func (t *Template) FieldAt(offset uint64) *Field {
	for _, f := range t.Fields {
		if f.Offset == offset {
			return f
		}
	}
	panic("invalid offset")
}
