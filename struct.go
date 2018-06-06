package bmstruct

import (
	"fmt"
	"sort"
	"unsafe"
)

type Struct struct {
	*Template
	Data []byte
}

func (t *Template) New(data []byte) *Struct {
	if t.Size != len(data) {
		panic("data bytes does not match the template size")
	}
	return &Struct{
		Template: t,
		Data:     data,
	}
}

func (t *Template) Empty() *Struct {
	return &Struct{
		Template: t,
		Data:     make([]byte, t.Size),
	}
}

func (s *Struct) Pointer() unsafe.Pointer {
	return unsafe.Pointer(&s.Data[0])
}

func (s *Struct) Uintptr() uintptr {
	return uintptr(s.Pointer())
}

// func (s *Struct) Data() []byte {
// 	return s.data
// }

func (s *Struct) Clone() *Struct {
	clone := &Struct{
		Template: s.Template,
		Data:     make([]byte, s.Template.Size),
	}
	copy(clone.Data, s.Data)
	return clone
}

// Lookup method of Struct returns the value if the field indicated by
// fieldName. A clone of the field is returned so modifying the returned value
// does not impact the Struct. For modifying the Struct object use the Update
// method.
func (s *Struct) Lookup(fieldName string) Value {
	field, found := s.Template.Fields[fieldName]
	if !found {
		panic(fmt.Sprintf("field name %s not found in template", fieldName))
	}
	return field.copySlice(s.Data)
}

func (s *Struct) Update(fieldName string, value Value) {
	field, found := s.Template.Fields[fieldName]
	if !found {
		panic(fmt.Sprintf("field name %s not found in template", fieldName))
	}
	if uint64(len(value)) != field.Len {
		panic(fmt.Sprintf("new value size (%d bytes) and field length (%d bytes) mismatch",
			uint64(len(value)), field.Len))
	}
	field.updateSlice(s.Data, value)
}

type Structs struct {
	Structs map[uint64]*Struct
	Data    []byte
}

func (t *Template) Slice(data []byte) *Structs {
	if len(data)%t.Size != 0 {
		panic("data bytes does not align")
	}
	// structs := make(map[uint64]*Struct)
	structs := &Structs{
		Structs: make(map[uint64]*Struct),
		Data:    data,
	}
	for offset := uint64(0); offset < uint64(len(data)); offset += uint64(t.Size) {
		structs.Structs[offset] = t.New(data[offset : offset+uint64(t.Size)])
	}
	return structs
}

func (ss *Structs) Count() uint32 {
	return uint32(len(ss.Structs))
}

func (ss *Structs) At(offset uint64) *Struct {
	return ss.Structs[offset]
}

func (ss *Structs) Update(offset uint64, s *Struct) {
	oldS, found := ss.Structs[offset]
	if !found {
		panic("invalid offset, no instruction found")
	}
	if !oldS.Template.Equal(s.Template) {
		panic("structs cannot be updated with different kind of struct")
	}
	copy(oldS.Data, s.Data)
}

func (ss *Structs) Clone() *Structs {
	clone := &Structs{
		Structs: make(map[uint64]*Struct),
		Data:    make([]byte, len(ss.Data)),
	}
	copy(clone.Data, ss.Data)
	for offset, s := range ss.Structs {
		clone.Structs[offset] = s.Template.New(clone.Data[offset : offset+uint64(s.Template.Size)])
	}
	return clone
}

func (ss *Structs) Pointer() unsafe.Pointer {
	return unsafe.Pointer(&ss.Data[0])
}

func (ss *Structs) Uintptr() uintptr {
	return uintptr(ss.Pointer())
}

type UInt64Slice []uint64

func (s UInt64Slice) Len() int {
	return len(s)
}

func (s UInt64Slice) Less(i, j int) bool {
	return s[i] < s[j]

}

func (s UInt64Slice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (ss *Structs) Offsets() []uint64 {
	offsets := make([]uint64, ss.Count())
	i := 0
	for offset := range ss.Structs {
		offsets[i] = offset
		i++
	}
	sort.Sort(UInt64Slice(offsets))
	return offsets
}
