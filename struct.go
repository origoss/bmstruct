package bmstruct

import (
	"fmt"
)

//Struct is a Template with an associated Value. You can think of a Struct as an
//"object" where the Template is the "class".
type Struct struct {
	*Template `json:"template"`
	Value     `json:"data"`
}

//New method instantiates a Struct object by mapping the given data to the
//Template.
//
//The size of the given data shall exactly match the Template size otherwise it
//will panic.
func (t *Template) New(data Valuable) *Struct {
	value := data.GetValue()
	if t.Size != len(value) {
		panic("data bytes does not match the template size")
	}
	return &Struct{
		Template: t,
		Value:    value,
	}
}

//Empty method instantiates a Struct object with empty data, i.e. all bytes
//are zeroed out.
func (t *Template) Empty() *Struct {
	return &Struct{
		Template: t,
		Value:    make([]byte, t.Size),
	}
}

//Clone method creates a new Struct object with the same template and the same
//cloned Value as the original Struct.
func (s *Struct) Clone() *Struct {
	clone := &Struct{
		Template: s.Template,
		Value:    s.Value.Clone(),
	}
	return clone
}

// Lookup method of Struct returns the Value if the field indicated by
// fieldName. A clone of the field is returned so modifying the returned value
// does not impact the Struct. For modifying the Struct object use the Update
// method.
//
// Lookup operation for a non-existing field name will panic.
func (s *Struct) Lookup(fieldName string) Value {
	field, found := s.Template.Fields[fieldName]
	if !found {
		panic(fmt.Sprintf("field name %s not found in template", fieldName))
	}
	return field.copySlice(s.Value)
}

// Update method of Struct changes the field indicated by fieldName to the given
// Value.
//
// Update operation will panic for a non-existing field name or incorrect Value
// size.
func (s *Struct) Update(fieldName string, valuable Valuable) {
	value := valuable.GetValue()
	field, found := s.Template.Fields[fieldName]
	if !found {
		panic(fmt.Sprintf("field name %s not found in template", fieldName))
	}
	if uint64(len(value)) != field.Len {
		panic(fmt.Sprintf("new value size (%d bytes) and field length (%d bytes) mismatch",
			uint64(len(value)), field.Len))
	}
	field.updateSlice(s.Value, value)
}

//Structs represents an array of Struct objects over a Value.
type Structs struct {
	*Template `json:"template"`
	Value     `json:"data"`
}

//Slice method creates a new Structs object. Slice will panic when the length of
//the given data does not align with the size of the Template.
func (t *Template) Slice(data Valuable) *Structs {
	value := data.GetValue()
	if len(value)%t.Size != 0 {
		panic("data bytes does not align")
	}
	structs := &Structs{
		Template: t,
		Value:    value,
	}
	return structs
}

//Count method returns the number of Struct objects in a Structs.
func (ss *Structs) Count() uint32 {
	return uint32(len(ss.Value) / ss.Template.Size)
}

//At method returns the Struct object that starts at the given offset. For
//returning the nth Struct object use the 'Nth' method call.
//
//At method panics when the offset is invalid (too large or not aligned
//with the Template size).
//
//At method returns a copy of data. Any modification on the returned Struct does
//not impact the Structs object.
func (ss *Structs) At(offset uint64) *Struct {
	if offset+uint64(ss.Template.Size) > uint64(len(ss.Value)) {
		panic("offset out of bounds")
	}
	if offset%uint64(ss.Template.Size) != 0 {
		panic("offset does not align with the Template size")
	}
	return ss.Template.New(ss.Value[offset : offset+uint64(ss.Template.Size)]).Clone()
}

//Nth method returns the nth Struct object. For returning the Struct object at a
//given offset use the 'At' method call.
//
//Nth method panics when n is invalid (too large or negative).
//
//Nth method returns a copy of data. Any modification on the returned Struct
//does not impact the Structs object.
func (ss *Structs) Nth(n int) *Struct {
	if n < 0 {
		panic("index out of bounds")
	}
	offset := uint64(n * ss.Template.Size)
	if offset+uint64(ss.Template.Size) > uint64(len(ss.Value)) {
		panic("index out of bounds")
	}
	return ss.Template.New(ss.Value[offset : offset+uint64(ss.Template.Size)]).Clone()
}

//Update method updates a Struct at the specified offset with the given Struct.
//
//Update panics if the offset is invalid (i.e. too large or does not align with
//the Template size) or the new Struct has a different kind of a Template.
func (ss *Structs) Update(offset uint64, s *Struct) {
	if offset+uint64(ss.Template.Size) > uint64(len(ss.Value)) {
		panic("offset out of bounds")
	}
	if offset%uint64(ss.Template.Size) != 0 {
		panic("invalid offset, no struct found")
	}
	if !ss.Template.Equal(s.Template) {
		panic("structs cannot be updated with different kind of struct")
	}
	copy(ss.Value[offset:offset+(uint64(ss.Template.Size))], s.Value)
}

//Clone method returns a new Structs that is the clone of the original Struct,
//i.e. the template is the same and the Value is cloned.
func (ss *Structs) Clone() *Structs {
	clone := &Structs{
		Value:    ss.Value.Clone(),
		Template: ss.Template,
	}
	return clone
}

//StructsIterFn is a function definition that is used by the Iter method if
//Structs.
type StructsIterFn func(offset uint64, s *Struct)

//Iter method iterates over the Struct objects stored in Structs. For each
//Struct the given 'fn' method is called with the offset and the Struct object.
func (ss *Structs) Iter(fn StructsIterFn) {
	for offset := uint64(0); offset < uint64(len(ss.Value)); offset += uint64(ss.Template.Size) {
		fn(offset, ss.At(offset))
	}
}
