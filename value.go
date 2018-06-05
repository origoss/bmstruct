package bmstruct

import (
	"reflect"
	"unsafe"
)

type Value []byte

func intValue(v interface{}) Value {
	t := reflect.TypeOf(v)
	bs := make([]byte, t.Size())
	for n := range bs {
		bs[n] = nthByteOfInt(v, n)
	}
	return bs
}

func Uint8(v uint8) Value {
	return intValue(v)
}

func Int16(v int16) Value {
	return intValue(v)
}

func Int32(v int32) Value {
	return intValue(v)
}

func Uint32(v uint32) Value {
	return intValue(v)
}

func Uint64(v uint64) Value {
	return intValue(v)
}

func Pointer(v unsafe.Pointer) Value {
	return Uintptr(uintptr(v))
}

func Uintptr(v uintptr) Value {
	return intValue(v)
}

func PtrToBytes(v []byte) Value {
	return Pointer(unsafe.Pointer(&v[0]))
}

func (v Value) Uint8() uint8 {
	var i uint8
	t := reflect.TypeOf(i)
	if uint64(t.Size()) != uint64(len(v)) {
		panic("value cannot be converted to Uint8, size mismatch")
	}
	for n := 0; n < int(t.Size()); n++ {
		i |= v[n] << uint(n*8)
	}
	return i
}

func (v Value) Int16() int16 {
	var i int16
	t := reflect.TypeOf(i)
	if uint64(t.Size()) != uint64(len(v)) {
		panic("value cannot be converted to Int16, size mismatch")
	}
	for n := 0; n < int(t.Size()); n++ {
		i |= int16(v[n]) << uint(n*8)
	}
	return i
}

func (v Value) Int32() int32 {
	var i int32
	t := reflect.TypeOf(i)
	if uint64(t.Size()) != uint64(len(v)) {
		panic("value cannot be converted to Int32, size mismatch")
	}
	for n := 0; n < int(t.Size()); n++ {
		i |= int32(v[n]) << uint(n*8)
	}
	return i
}

func (v Value) Uint32() uint32 {
	var i uint32
	t := reflect.TypeOf(i)
	if uint64(t.Size()) != uint64(len(v)) {
		panic("value cannot be converted to Uint32, size mismatch")
	}
	for n := 0; n < int(t.Size()); n++ {
		i |= uint32(v[n]) << uint(n*8)
	}
	return i
}

func (v Value) Uint64() uint64 {
	var i uint64
	t := reflect.TypeOf(i)
	if uint64(t.Size()) != uint64(len(v)) {
		panic("value cannot be converted to Uint64, size mismatch")
	}
	for n := 0; n < int(t.Size()); n++ {
		i |= uint64(v[n]) << uint(n*8)
	}
	return i
}
