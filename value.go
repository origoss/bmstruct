package bmstruct

import (
	"reflect"
	"unsafe"
)

//Valuable interface defines the Value() method that returns a Value. If an
//object implements the Valueable interface, it can be used where Values can be
//used.
type Valuable interface {
	GetValue() Value
	Address() uintptr
}

//NilType is a special type. The Nil object with this type should be used in
//programs.
type NilType struct{}

//GetValue implements the Valuable interface for NilType
func (n *NilType) GetValue() Value {
	panic("Getting Value of Nil type")
}

//Address implements the Valuable interface for NilType
func (n *NilType) Address() uintptr {
	return 0
}

//Nil represents a Value for which the Address() method returns 0 and the
//GetValue() panics.
var Nil = NilType{}

//Value is a byte slice with extra capabilities. A Value can convert itself into
//the most common go types (int, []byte, string, etc.) and vice versa the most
//common go types can be converted easily to Value.
type Value []byte

//Value method of a Value returns itself
func (v Value) GetValue() Value {
	return v
}
func intValue(v interface{}) Value {
	t := reflect.TypeOf(v)
	bs := make([]byte, t.Size())
	for n := range bs {
		bs[n] = nthByteOfInt(v, n)
	}
	return bs
}

//Uint8 function converts an uint8 value to Value type.
func Uint8(v uint8) Value {
	return intValue(v)
}

//Uint8 method returns the uint8 representation of a Value. The method will
//panic if the Value's length is incorrect.
func (v Value) Uint8() uint8 {
	if len(v) != 1 {
		panic("value cannot be converted to Uint8, size mismatch")
	}
	return v[0]
}

//Int8 function converts an int8 value to Value type.
func Int8(v int8) Value {
	return intValue(v)
}

//Int8 method returns the int8 representation of a Value. The method will
//panic if the Value's length is incorrect.
func (v Value) Int8() int8 {
	if len(v) != 1 {
		panic("value cannot be converted to Int8, size mismatch")
	}
	return int8(v[0])
}

//Uint16 function converts an uint16 value to Value type.
func Uint16(v uint16) Value {
	return intValue(v)
}

//Uint16 method returns the uint16 representation of a Value. The method will
//panic if the Value's length is incorrect.
func (v Value) Uint16() uint16 {
	var i uint16
	t := reflect.TypeOf(i)
	if uint64(t.Size()) != uint64(len(v)) {
		panic("value cannot be converted to Uint16, size mismatch")
	}
	for n := 0; n < int(t.Size()); n++ {
		i |= uint16(v[n]) << uint(n*8)
	}
	return i
}

//Int16 function converts an int16 value to Value type.
func Int16(v int16) Value {
	return intValue(v)
}

//Int16 method returns the int16 representation of a Value. The method will
//panic if the Value's length is incorrect.
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

//Uint32 function converts an uint32 value to Value type.
func Uint32(v uint32) Value {
	return intValue(v)
}

//Uint32 method returns the uint32 representation of a Value. The method will
//panic if the Value's length is incorrect.
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

//Int32 function converts an int32 value to Value type.
func Int32(v int32) Value {
	return intValue(v)
}

//Int32 method returns the int32 representation of a Value. The method will
//panic if the Value's length is incorrect.
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

//Uint64 function converts an uint64 value to Value type.
func Uint64(v uint64) Value {
	return intValue(v)
}

//Uint64 method returns the uint64 representation of a Value. The method will
//panic if the Value's length is incorrect.
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

//Int64 function converts an int64 value to Value type.
func Int64(v int64) Value {
	return intValue(v)
}

//Int64 method returns the int64 representation of a Value. The method will
//panic if the Value's length is incorrect.
func (v Value) Int64() int64 {
	var i int64
	t := reflect.TypeOf(i)
	if uint64(t.Size()) != uint64(len(v)) {
		panic("value cannot be converted to Uint64, size mismatch")
	}
	for n := 0; n < int(t.Size()); n++ {
		i |= int64(v[n]) << uint(n*8)
	}
	return i
}

//Uintptr function converts an uintptr value to Value type.
func Uintptr(v uintptr) Value {
	return intValue(v)
}

//Uintptr method returns the uintptr representation of a Value. The method will
//panic if the Value's length is incorrect.
//
//Note, that Uintptr() does not return the address of the byte slice but an 8
//byte long integer that can be used for storing any address in memory.
//
//If you want to get the pointer to the byte slice backing the Value use the
//Address() method instead.
func (v Value) Uintptr() uintptr {
	var i uintptr
	t := reflect.TypeOf(i)
	if uint64(t.Size()) != uint64(len(v)) {
		panic("value cannot be converted to Uint64, size mismatch")
	}
	for n := 0; n < int(t.Size()); n++ {
		i |= uintptr(v[n]) << uint(n*8)
	}
	return i
}

//ByteSlice function converts an []byte value to Value type. The value of 'b'
//will not be copied. Value will just refer to the byte slice.
func ByteSlice(b []byte) Value {
	return Value(b)
}

//ByteSlice method returns the []byte representation of a Value. The returned
//value is a reference to the byte slice of the Value, not a copy.
func (v Value) ByteSlice() []byte {
	return []byte(v)
}

//ZeroTermString function converts a Go string to a value so that the internal representation will be a zero-terminated string.
//
//Example:
//  ZeroTermString("Hello") == Value{'H', 'e', 'l', 'l', 'o', 0}
func ZeroTermString(s string) Value {
	b := make([]byte, len(s)+1)
	copy(b, []byte(s))
	return b
}

//ZeroTermString method returns the string representation of a Value. The Value
//is expected to store a zero terminated string thus the last character will be
//trimmed from the result.
func (v Value) ZeroTermString() string {
	return string(v[0 : len(v)-1])
}

//Address method returns the memory address as uintptr of the byte slice that
//backs to Value.
func (v Value) Address() uintptr {
	return uintptr(unsafe.Pointer(&v[0]))
}

//Clone method returns a new Value with a copy of the original byte slice.
func (v Value) Clone() Value {
	clone := make([]byte, len(v))
	copy(clone, v)
	return clone
}
