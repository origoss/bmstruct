package bmstruct

import (
	"reflect"
)

//Field represents a smaller fraction of a byte slice or Value that in itself
//contains a meaningful value, like an int or string.
//
//Let's define 'f1' as the followings
//
//  f1 := Field {
//      Name:   "f1",
//      Offset: 1,
//      Len:    2,
//  }
//
//'f1' is a 2 bytes long slice of any byte slice starting at the offset 1.
//
//  0:  01001001
//  1: |11100110|
//  2: |00010100|
//  3:  00001000
//      ...
//
//Field can represent a bit field that is smaller than a byte. 'f2' is such a
//Field. Len of a bit field must be exactly 1.
//
//  f2 := Field{
//      Name:           "f2",
//      Offset:         1,
//      Len:            1,
//      BitFieldOffset: 3,
//      BitFieldLen:    4,
//  }
//
//  0:  01001001
//  1: |11100110|
//       ^^^^
//  2:  00010100
//  3:  00001000
type Field struct {
	Name           string `json:"name"`
	Offset         uint64 `json:"offset"`
	Len            uint64 `json:"length"`
	BitFieldOffset uint8  `json:"bf-offset,omitempty"`
	BitFieldLen    uint8  `json:"bf-len,omitempty"`
}

func newField(t reflect.Type, name string, offset uint64) *Field {
	return &Field{
		Name:   name,
		Offset: offset,
		Len:    uint64(t.Size()),
	}
}

func (f *Field) slice(data []byte) []byte {
	if f.BitFieldLen != 0 {
		panic("slice shouldn't be used for bit fields")
	}
	return data[f.Offset : f.Offset+f.Len]
}

func (f *Field) copySlice(data []byte) []byte {
	if f.BitFieldLen != 0 {
		return []byte{
			bitFieldOfByte(data[f.Offset],
				f.BitFieldOffset,
				f.BitFieldLen)}
	}
	b := make([]byte, f.Len)
	copy(b, f.slice(data))
	return b
}

func (f *Field) updateSlice(data []byte, value Value) {
	if f.BitFieldLen != 0 {
		if len(value) != 1 {
			panic("BitField value shall contain a single byte only")
		}
		setBitFieldOfByte(
			&data[f.Offset],
			f.BitFieldOffset,
			f.BitFieldLen,
			value[0])
	} else {
		if uint64(len(value)) != f.Len {
			panic("input value has incorrect length")
		}
		copy(f.slice(data), value)
	}
}

//BitField function creates a new Field with the given name, offset,
//bitFieldOffset and bitFieldLen.
func BitField(name string, offset uint64,
	bitFieldOffset, bitFieldLen uint8) *Field {
	if bitFieldLen == 0 {
		panic("invalid bitfield: length cannot be 0")
	}
	if bitFieldOffset+bitFieldLen > 8 {
		panic("invalid bitfield: offset+length cannot be larger than 8")
	}
	f := newField(reflect.TypeOf(uint8(0)), name, offset)
	f.BitFieldOffset = bitFieldOffset
	f.BitFieldLen = bitFieldLen
	return f
}

//Uint8Field creates a new Field with the given name and offset. Len is
//calculated to fit a uint8 value.
func Uint8Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(uint8(0)),
		name,
		offset,
	)
}

//Int8Field creates a new Field with the given name and offset. Len is
//calculated to fit an int8 value.
func Int8Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(int8(0)),
		name,
		offset,
	)
}

//Uint16Field creates a new Field with the given name and offset. Len is
//calculated to fit a uint16 value.
func Uint16Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(uint16(0)),
		name,
		offset,
	)
}

//Int16Field creates a new Field with the given name and offset. Len is
//calculated to fit an int16 value.
func Int16Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(int16(0)),
		name,
		offset,
	)
}

//Uint32Field creates a new Field with the given name and offset. Len is
//calculated to fit a uint32 value.
func Uint32Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(uint32(0)),
		name,
		offset,
	)

}

//Int32Field creates a new Field with the given name and offset. Len is
//calculated to fit an int32 value.
func Int32Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(int32(0)),
		name,
		offset,
	)
}

//Uint64Field creates a new Field with the given name and offset. Len is
//calculated to fit a uint64 value.
func Uint64Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(uint64(0)),
		name,
		offset,
	)

}

//Int64Field creates a new Field with the given name and offset. Len is
//calculated to fit an int64 value.
func Int64Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(int64(0)),
		name,
		offset,
	)
}

//UintField creates a new Field with the given name and offset. Len is
//calculated to fit a uint value.
func UintField(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(uint(0)),
		name,
		offset,
	)
}

//IntField creates a new Field with the given name and offset. Len is
//calculated to fit an int value.
func IntField(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(int(0)),
		name,
		offset,
	)
}
