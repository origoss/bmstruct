package bmstruct

import (
	"fmt"
	"reflect"
)

// func interfaceToBytes(i interface{}) ([]byte, error) {
// 	iType := reflect.TypeOf(i)
// 	// iValue := reflect.ValueOf(i)
// 	switch iType.Kind() {
// 	case reflect.Int, reflect.Uint, reflect.Int8, reflect.Uint8,
// 		reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32,
// 		reflect.Int64, reflect.Uint64:
// 		bs := make([]byte, iType.Size())
// 		for n := range bs {
// 			var err error
// 			bs[n], err = nthByteOfInt(i, n)
// 			if err != nil {
// 				return nil, errors.Wrap(err, "nthByteOfInt failed")
// 			}
// 		}
// 		return bs, nil
// 	case reflect.Slice:
// 		iValue := reflect.ValueOf(i)
// 		elemType := iType.Elem()
// 		elemSize := int(elemType.Size())
// 		bs := make([]byte, iValue.Len()*elemSize)
// 		for n := 0; n < iValue.Len(); n++ {
// 			elemBytes, err := interfaceToBytes(iValue.Index(n).Interface())
// 			if err != nil {
// 				return nil, errors.Wrap(err, "processing Slice in interfaceToBytes")
// 			}
// 			copy(bs[n*elemSize:(n+1)*elemSize], elemBytes)
// 		}
// 		return bs, nil
// 	default:
// 		return nil, errors.Errorf("interfaceToBytes does not support %s", iType.Kind())
// 	}
// }

// bitmap                                         = 0b11001010 0b10101010
// offset                                         = 4          0
// len                                            = 2          1
// value                                          = 0b00000010 0b00000000
//
// mask          := (1 << len) - 1                = 0b00000011 0b00000001
// masked_value  := value & mask                  = 0b00000010 0b00000000
// shifted_mask  := mask << offset                = 0b00110000 0b00000001
// shifted_value := masked_value << offset        = 0b00100000 0b00000000
// erased_bitmap := bitmap &^ shifted_mask        = 0b11001010 0b10101010
// value         := erased_bitmap | shifted_value = 0b11101010 0b10101010
func setBitFieldOfByte(b *byte, offset, length uint8, value byte) {
	if length == 0 {
		panic("length cannot be less than 1")
	}
	if offset+length > 8 {
		panic("offset+length cannot be larger than 8")
	}
	mask := byte((1 << length) - 1)
	maskedValue := value & mask
	shiftedMask := mask << uint(offset)
	shiftedValue := maskedValue << uint(offset)
	erasedBitMap := *b &^ shiftedMask
	*b = erasedBitMap | shiftedValue
}

//   76543210
// 0b11001010, 4, 2 => 0b11(00)1010
//
// 0b00000001 < 2(len)
// 0b00000100 - 1                   = 0b00000011
//
// 0b11001010 >>  offset            = 0b00001100
//
// 0b00001100 & 0b00000011          = 0b00000000

func bitFieldOfByte(b byte, offset, length uint8) byte {
	if length == 0 {
		panic("length cannot be less than 1")
	}
	if offset+length > 8 {
		panic("offset+length cannot be larger than 8")
	}
	// return (b << uint(8-offset-length)) & ((1 << uint(length)) - 1)
	return (b >> uint(offset)) & ((1 << uint(length)) - 1)
}

// nthByteOfInt returns the nth byte of an int type that is represented as
// []byte. A mutable slice is returned
func nthByteOfInt(kindOfInt interface{}, n int) byte {
	if n < 0 {
		panic("n shall be larger than 0")
	}
	switch iValue := kindOfInt.(type) {
	case int:
		if n > 7 {
			panic("n is out of bounds")
		}
		return byte(iValue & (0xff << uint(n*8)) >> (uint(n * 8)))
	case uint:
		if n > 7 {
			panic("n is out of bounds")
		}
		return byte(iValue & (0xff << uint(n*8)) >> (uint(n * 8)))
	case int8:
		if n > 0 {
			panic("n is out of bounds")
		}
		return byte(iValue)
	case uint8:
		if n > 0 {
			panic("n is out of bounds")
		}
		return iValue & (0xff << uint(n*8)) >> (uint(n * 8))
	case int16:
		if n > 1 {
			panic("n is out of bounds")
		}
		return byte(iValue & (0xff << uint(n*8)) >> (uint(n * 8)))
	case uint16:
		if n > 1 {
			panic("n is out of bounds")
		}
		return byte(iValue & (0xff << uint(n*8)) >> (uint(n * 8)))
	case int32:
		if n > 3 {
			panic("n is out of bounds")
		}
		return byte(iValue & (0xff << uint(n*8)) >> (uint(n * 8)))
	case uint32:
		if n > 3 {
			panic("n is out of bounds")
		}
		return byte(iValue & (0xff << uint(n*8)) >> (uint(n * 8)))
	case int64:
		if n > 7 {
			panic("n is out of bounds")
		}
		return byte(iValue & (0xff << uint(n*8)) >> (uint(n * 8)))
	case uint64:
		if n > 7 {
			panic("n is out of bounds")
		}
		return byte(iValue & (0xff << uint(n*8)) >> (uint(n * 8)))
	case uintptr:
		if n > 7 {
			panic("n is out of bounds")
		}
		return byte(iValue & (0xff << uint(n*8)) >> (uint(n * 8)))
	default:
		panic(fmt.Sprintf("nthByteOfInt does not handle %s", reflect.TypeOf(kindOfInt).Kind()))
	}
}

// func bytesToInterface(t reflect.Type, bs []byte) (interface{}, error) {
// 	switch t.Kind() {
// 	// case reflect.Int, reflect.Uint, reflect.Int8, reflect.Uint8,
// 	// 	reflect.Int16, reflect.Uint16, reflect.Int32, reflect.Uint32,
// 	// 	reflect.Int64, reflect.Uint64:
// 	case reflect.Int:
// 		var v int
// 		for n := 0; n < int(t.Size()); n++ {
// 			v |= int(bs[n]) << uint(n*8)
// 		}
// 		return v, nil
// 	case reflect.Uint:
// 		var v uint
// 		for n := 0; n < int(t.Size()); n++ {
// 			v |= uint(bs[n]) << uint(n*8)
// 		}
// 		return v, nil
// 	case reflect.Int8:
// 		var v int8
// 		for n := 0; n < int(t.Size()); n++ {
// 			v |= int8(bs[n]) << uint(n*8)
// 		}
// 		return v, nil
// 	case reflect.Uint8:
// 		var v uint8
// 		for n := 0; n < int(t.Size()); n++ {
// 			v |= bs[n] << uint(n*8)
// 		}
// 		return v, nil
// 	case reflect.Int16:
// 		var v int16
// 		for n := 0; n < int(t.Size()); n++ {
// 			v |= int16(bs[n]) << uint(n*8)
// 		}
// 		return v, nil
// 	case reflect.Uint16:
// 		var v uint16
// 		for n := 0; n < int(t.Size()); n++ {
// 			v |= uint16(bs[n]) << uint(n*8)
// 		}
// 		return v, nil
// 	case reflect.Int32:
// 		var v int32
// 		for n := 0; n < int(t.Size()); n++ {
// 			v |= int32(bs[n]) << uint(n*8)
// 		}
// 		return v, nil
// 	case reflect.Uint32:
// 		var v uint32
// 		for n := 0; n < int(t.Size()); n++ {
// 			v |= uint32(bs[n]) << uint(n*8)
// 		}
// 		return v, nil
// 	case reflect.Int64:
// 		var v int64
// 		for n := 0; n < int(t.Size()); n++ {
// 			v |= int64(bs[n]) << uint(n*8)
// 		}
// 		return v, nil
// 	case reflect.Uint64:
// 		var v uint64
// 		for n := 0; n < int(t.Size()); n++ {
// 			v |= uint64(bs[n]) << uint(n*8)
// 		}
// 		return v, nil
// 		// TODO
// 	// case reflect.Slice:
// 	// 	elemType := t.Elem()
// 	// 	elemSize := int(elemType.Size())
// 	// 	if len(bs)%elemSize != 0 {
// 	// 		return nil, errors.Errorf("bytes in byteToInterface does not")
// 	// 	}

// 	default:
// 		return nil, errors.Errorf("bytesToInterface does not support %s", t.Kind())
// 	}
// }
