package bmstruct

import (
	"reflect"
)

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
		// Type:   t,
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
		copy(f.slice(data), value)
	}
}

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

func IntField(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(int(0)),
		name,
		offset,
	)
}

func Uint8Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(uint8(0)),
		name,
		offset,
	)
}

func Int16Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(int16(0)),
		name,
		offset,
	)
}

func Int32Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(int32(0)),
		name,
		offset,
	)
}

func Uint32Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(uint32(0)),
		name,
		offset,
	)
}

func Uint64Field(name string, offset uint64) *Field {
	return newField(reflect.TypeOf(uint64(0)),
		name,
		offset,
	)
}
