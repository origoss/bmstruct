package bmstruct

import (
	"reflect"
)

type Template struct {
	Fields map[string]*Field `json:"fields"`
	Size   int               `json:"size"`
}

func NewTemplate(size int, fields ...*Field) *Template {
	if size < 1 {
		panic("size shall be greater than 0")
	}
	t := &Template{
		Fields: make(map[string]*Field),
		Size:   size,
	}
	for _, field := range fields {
		t.Fields[field.Name] = field
	}
	if t.minLen() > uint64(size) {
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

func (t *Template) Equal(other *Template) bool {
	return reflect.DeepEqual(t, other) || reflect.DeepEqual(*t, *other)
}
