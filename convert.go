package bmstruct

import (
	"fmt"
	"reflect"
)

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
