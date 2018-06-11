package bmstruct

import (
	// "encoding/base64"
	"fmt"
	"unsafe"
)

// type binary []byte

// func (b binary) MarshalText() (text []byte, err error) {
// 	dst := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
// 	base64.StdEncoding.Encode(dst, b)
// 	return dst, nil
// }

// func (b binary) UnmarshalText(text []byte) error {
// 	b = make([]byte, base64.StdEncoding.DecodedLen(len(text)))
// 	base64.StdEncoding.Decode(b, text)
// 	return nil
// }

type Struct struct {
	*Template `json:"template"`
	Data      []byte `json:"data"`
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

func (s *Struct) Value() Value {
	return Value(s.Data)
}

func (s *Struct) Pointer() unsafe.Pointer {
	return unsafe.Pointer(&s.Data[0])
}

func (s *Struct) Uintptr() uintptr {
	return uintptr(s.Pointer())
}

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

func (s *Struct) Update(fieldName string, valuable Valuable) {
	value := valuable.Value()
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
	*Template
	Data []byte
}

func (t *Template) Slice(data []byte) *Structs {
	if len(data)%t.Size != 0 {
		panic("data bytes does not align")
	}
	structs := &Structs{
		Template: t,
		Data:     data,
	}
	return structs
}

func (ss *Structs) Count() uint32 {
	return uint32(len(ss.Data) / ss.Template.Size)
}

func (ss *Structs) At(offset uint64) *Struct {
	if offset%uint64(ss.Template.Size) != 0 {
		return nil
	}
	return ss.Template.New(ss.Data[offset : offset+uint64(ss.Template.Size)])
}

func (ss *Structs) Nth(n int) *Struct {
	offset := uint64(n * ss.Template.Size)
	if offset+uint64(ss.Template.Size) > uint64(len(ss.Data)) {
		panic("index out of bounds")
	}
	return ss.Template.New(ss.Data[offset : offset+uint64(ss.Template.Size)])
}

func (ss *Structs) Update(offset uint64, s *Struct) {
	if offset%uint64(ss.Template.Size) != 0 {
		panic("invalid offset, no struct found")
	}
	if !ss.Template.Equal(s.Template) {
		panic("structs cannot be updated with different kind of struct")
	}
	copy(ss.Data[offset:offset+(uint64(ss.Template.Size))], s.Data)
}

func (ss *Structs) Clone() *Structs {
	clone := &Structs{
		Data:     make([]byte, len(ss.Data)),
		Template: ss.Template,
	}
	copy(clone.Data, ss.Data)
	return clone
}

func (ss *Structs) Pointer() unsafe.Pointer {
	return unsafe.Pointer(&ss.Data[0])
}

func (ss *Structs) Uintptr() uintptr {
	return uintptr(ss.Pointer())
}

func (ss *Structs) Offsets() []uint64 {
	offsets := make([]uint64, ss.Count())
	i := 0
	for offset := uint64(0); offset < uint64(len(ss.Data)); offset += uint64(ss.Template.Size) {
		offsets[i] = offset
		i++
	}
	return offsets
}

func (ss *Structs) Value() Value {
	return Value(ss.Data)
}
