package bmstruct

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Convert", func() {
	Describe("setBitFieldOfByte", func() {
		var b byte

		BeforeEach(func() {
			b = 0xAA // 0b10101010
		})
		Context("when length > 8", func() {
			It("should panic", func() {
				Expect(func() {
					setBitFieldOfByte(&b, 0, 9, 1)
				}).To(Panic())
			})
		})
		Context("when length =< 0", func() {
			It("should panic", func() {
				Expect(func() {
					setBitFieldOfByte(&b, 0, 0, 1)
				}).To(Panic())
			})
		})
		Context("when offset + length > 8", func() {
			It("should panic", func() {
				Expect(func() {
					setBitFieldOfByte(&b, 6, 3, 1)
				}).To(Panic())
			})
		})
		Context("when setting a single bit to 1", func() {
			It("should return the correct value", func() {
				bs := make([]byte, 8)
				for i := range bs {
					bClone := b
					setBitFieldOfByte(&bClone, uint8(i), 1, 1)
					bs[i] = bClone
				}
				Expect(bs).To(Equal([]byte{
					0xAB, // 0b10101011
					0xAA, // 0b10101010
					0xAE, // 0b10101110
					0xAA, // 0b10101010
					0xBA, // 0b10111010
					0xAA, // 0b10101010
					0xEA, // 0b11101010
					0xAA, // 0b10101010
				}))
			})
		})
		Context("when setting a single bit to 0", func() {
			It("should return the correct value", func() {
				bs := make([]byte, 8)
				for i := range bs {
					bClone := b
					setBitFieldOfByte(&bClone, uint8(i), 1, 0)
					bs[i] = bClone
				}
				Expect(bs).To(Equal([]byte{
					0xAA, // 0b10101010
					0xA8, // 0b10101000
					0xAA, // 0b10101010
					0xA2, // 0b10100010
					0xAA, // 0b10101010
					0x8A, // 0b10001010
					0xAA, // 0b10101010
					0x2A, // 0b00101010
				}))
			})
		})
		Context("when setting bits to 11", func() {
			It("should return the correct value", func() {
				bs := make([]byte, 7)
				for i := range bs {
					bClone := b
					setBitFieldOfByte(&bClone, uint8(i), 2, 3)
					bs[i] = bClone
				}
				Expect(bs).To(Equal([]byte{
					0xAB, // 0b10101011
					0xAE, // 0b10101110
					0xAE, // 0b10101110
					0xBA, // 0b10111010
					0xBA, // 0b10111010
					0xEA, // 0b11101010
					0xEA, // 0b11101010
				}))
			})
		})
		Context("when setting bits to 01", func() {
			It("should return the correct value", func() {
				bs := make([]byte, 7)
				for i := range bs {
					bClone := b
					setBitFieldOfByte(&bClone, uint8(i), 2, 1)
					bs[i] = bClone
				}
				Expect(bs).To(Equal([]byte{
					0xA9, // 0b10101001
					0xAA, // 0b10101010
					0xA6, // 0b10100110
					0xAA, // 0b10101010
					0x9A, // 0b10011010
					0xAA, // 0b10101010
					0x6A, // 0b01101010
				}))
			})
		})
	})
	Describe("bitFieldOfByte", func() {
		var b byte

		BeforeEach(func() {
			// 0b11001010 = 0xCA
			b = 0xCA
		})
		Context("when length > 8", func() {
			It("should panic", func() {
				Expect(func() {
					bitFieldOfByte(b, 0, 9)
				}).To(Panic())
			})
		})
		Context("when length =< 0", func() {
			It("should panic", func() {
				Expect(func() {
					bitFieldOfByte(b, 0, 0)
				}).To(Panic())
			})
		})
		Context("when offset + length > 8", func() {
			It("should panic", func() {
				Expect(func() {
					bitFieldOfByte(b, 6, 3)
				}).To(Panic())
			})
		})
		Context("when length = 1", func() {
			It("should return correct values", func() {
				bs := make([]byte, 8)
				for i := range bs {
					bs[i] = bitFieldOfByte(b, uint8(i), 1)
				}
				Expect(bs).To(Equal([]byte{
					0, 1, 0, 1, 0, 0, 1, 1,
				}))
			})
		})
		Context("when length = 2", func() {
			It("should return correct values", func() {
				bs := make([]byte, 7)
				for i := range bs {
					bs[i] = bitFieldOfByte(b, uint8(i), 2)
				}
				Expect(bs).To(Equal([]byte{
					// 10, 01, 10, 01, 00, 10, 11
					2, 1, 2, 1, 0, 2, 3,
				}))
			})
		})
		Context("when length = 3", func() {
			It("should return correct values", func() {
				bs := make([]byte, 6)
				for i := range bs {
					bs[i] = bitFieldOfByte(b, uint8(i), 3)
				}
				Expect(bs).To(Equal([]byte{
					// 010, 101, 010, 001, 100, 110
					2, 5, 2, 1, 4, 6,
				}))
			})
		})
		Context("when length = 8", func() {
			It("should return correct values", func() {
				bs := make([]byte, 1)
				for i := range bs {
					bs[i] = bitFieldOfByte(b, uint8(i), 8)
				}
				Expect(bs).To(Equal([]byte{
					// 0b11001010 = 0xCA
					0xCA,
				}))
			})
		})
	})
	Describe("nthByteOfInt", func() {
		Context("for string", func() {
			It("shall panic", func() {
				Expect(func() {
					nthByteOfInt("is a string", 0)
				}).To(Panic())
			})
		})
		Context("for int", func() {
			var v int

			BeforeEach(func() {
				v = 0x0123456789ABCDEF
			})
			It("shall return the correct values for n = 0-7", func() {
				b := make([]byte, 8)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0xEF, 0xCD, 0xAB, 0x89,
					0x67, 0x45, 0x23, 0x01,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 7", func() {
				Expect(func() {
					nthByteOfInt(v, 8)
				}).To(Panic())
			})
		})
		Context("for negative int", func() {
			var v int

			BeforeEach(func() {
				v = -0x0123456789ABCDEF
			})
			It("shall return the correct values for n = 0-7", func() {
				b := make([]byte, 8)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0x11, 0x32, 0x54, 0x76,
					0x98, 0xBA, 0xDC, 0xFE,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 7", func() {
				Expect(func() {
					nthByteOfInt(v, 8)
				}).To(Panic())
			})
		})
		Context("for uint", func() {
			var v uint

			BeforeEach(func() {
				v = 0x0123456789ABCDEF
			})
			It("shall return the correct values for n = 0-7", func() {
				b := make([]byte, 8)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0xEF, 0xCD, 0xAB, 0x89,
					0x67, 0x45, 0x23, 0x01,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 7", func() {
				Expect(func() {
					nthByteOfInt(v, 8)
				}).To(Panic())
			})
		})
		Context("for int8", func() {
			var v int8

			BeforeEach(func() {
				v = 0x01
			})
			It("shall return the correct values for n = 0", func() {
				b := make([]byte, 1)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0x01,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 0", func() {
				Expect(func() {
					nthByteOfInt(v, 1)
				}).To(Panic())
			})
		})
		Context("for negative int8", func() {
			var v int8

			BeforeEach(func() {
				v = -0x01
			})
			It("shall return the correct values for n = 0", func() {
				b := make([]byte, 1)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0xFF,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 0", func() {
				Expect(func() {
					nthByteOfInt(v, 1)
				}).To(Panic())
			})
		})
		Context("for uint8", func() {
			var v uint8

			BeforeEach(func() {
				v = 0x01
			})
			It("shall return the correct values for n = 0", func() {
				b := make([]byte, 1)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0x01,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 0", func() {
				Expect(func() {
					nthByteOfInt(v, 1)
				}).To(Panic())
			})
		})
		Context("for int16", func() {
			var v int16

			BeforeEach(func() {
				v = 0x0123
			})
			It("shall return the correct values for n = 0-1", func() {
				b := make([]byte, 2)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0x23, 0x01,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 1", func() {
				Expect(func() {
					nthByteOfInt(v, 2)
				}).To(Panic())
			})
		})
		Context("for negative int16", func() {
			var v int16

			BeforeEach(func() {
				v = -0x0123
			})
			It("shall return the correct values for n = 0-1", func() {
				b := make([]byte, 2)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0xDD, 0xFE,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 1", func() {
				Expect(func() {
					nthByteOfInt(v, 2)
				}).To(Panic())
			})
		})
		Context("for uint16", func() {
			var v uint16

			BeforeEach(func() {
				v = 0x0123
			})
			It("shall return the correct values for n = 0-1", func() {
				b := make([]byte, 2)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0x23, 0x01,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 1", func() {
				Expect(func() {
					nthByteOfInt(v, 2)
				}).To(Panic())
			})
		})
		Context("for int32", func() {
			var v int32

			BeforeEach(func() {
				v = 0x01234567
			})
			It("shall return the correct values for n = 0-3", func() {
				b := make([]byte, 4)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0x67, 0x45, 0x23, 0x01,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 3", func() {
				Expect(func() {
					nthByteOfInt(v, 4)
				}).To(Panic())
			})
		})
		Context("for negative int32", func() {
			var v int32

			BeforeEach(func() {
				v = -0x01234567
			})
			It("shall return the correct values for n = 0-3", func() {
				b := make([]byte, 4)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0x99, 0xBA, 0xDC, 0xFE,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 3", func() {
				Expect(func() {
					nthByteOfInt(v, 4)
				}).To(Panic())
			})
		})
		Context("for uint32", func() {
			var v uint32

			BeforeEach(func() {
				v = 0x01234567
			})
			It("shall return the correct values for n = 0-3", func() {
				b := make([]byte, 4)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0x67, 0x45, 0x23, 0x01,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 3", func() {
				Expect(func() {
					nthByteOfInt(v, 4)
				}).To(Panic())
			})
		})
		Context("for int64", func() {
			var v int64

			BeforeEach(func() {
				v = 0x0123456789ABCDEF
			})
			It("shall return the correct values for n = 0-7", func() {
				b := make([]byte, 8)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0xEF, 0xCD, 0xAB, 0x89,
					0x67, 0x45, 0x23, 0x01,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 7", func() {
				Expect(func() {
					nthByteOfInt(v, 8)
				}).To(Panic())
			})
		})
		Context("for negative int64", func() {
			var v int64

			BeforeEach(func() {
				v = -0x0123456789ABCDEF
			})
			It("shall return the correct values for n = 0-7", func() {
				b := make([]byte, 8)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0x11, 0x32, 0x54, 0x76,
					0x98, 0xBA, 0xDC, 0xFE,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 7", func() {
				Expect(func() {
					nthByteOfInt(v, 8)
				}).To(Panic())
			})
		})
		Context("for uint64", func() {
			var v uint64

			BeforeEach(func() {
				v = 0x0123456789ABCDEF
			})
			It("shall return the correct values for n = 0-7", func() {
				b := make([]byte, 8)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0xEF, 0xCD, 0xAB, 0x89,
					0x67, 0x45, 0x23, 0x01,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 7", func() {
				Expect(func() {
					nthByteOfInt(v, 8)
				}).To(Panic())
			})
		})
		Context("for uintptr", func() {
			var v uintptr

			BeforeEach(func() {
				v = 0x0123456789ABCDEF
			})
			It("shall return the correct values for n = 0-7", func() {
				b := make([]byte, 8)
				for i := range b {
					b[i] = nthByteOfInt(v, i)
				}
				Expect(b).To(Equal([]byte{
					0xEF, 0xCD, 0xAB, 0x89,
					0x67, 0x45, 0x23, 0x01,
				}))
			})
			It("shall panic for n < 0", func() {
				Expect(func() {
					nthByteOfInt(v, -42)
				}).To(Panic())
			})
			It("shall panic for n > 7", func() {
				Expect(func() {
					nthByteOfInt(v, 8)
				}).To(Panic())
			})
		})
	})
})
