package bmstruct

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Value", func() {
	Describe("Nil type", func() {
		It("should panic for GetValue()", func() {
			Expect(func() {
				Nil.GetValue()
			}).To(Panic())
		})
		It("should return 0 for Address()", func() {
			Expect(Nil.Address()).To(Equal(uintptr(0)))
		})
	})
	Describe("IsNil", func() {
		It("should return true for Nil", func() {
			Expect(IsNil(Nil)).To(BeTrue())
		})
		It("should return false for other types", func() {
			Expect(IsNil(Uint16(4234))).NotTo(BeTrue())
		})
	})
	Describe("Value of Value", func() {
		It("should always return itself", func() {
			v := Value{1, 2, 3, 4}
			Expect(v.GetValue()).To(Equal(v))
		})
	})
	Describe("Uint8", func() {
		It("should generate the expected Value", func() {
			Expect(Uint8(0)).To(Equal(Value{0}))
			Expect(Uint8(132)).To(Equal(Value{132}))
		})
		It("should convert in both direction properly", func() {
			Expect(Uint8(uint8(0)).Uint8()).To(Equal(uint8(0)))
			Expect(Uint8(uint8(132)).Uint8()).To(Equal(uint8(132)))
		})
		It("should panic when Value's length is incorrect", func() {
			Expect(func() {
				Value{1, 2}.Uint8()
			}).To(Panic())
			Expect(func() {
				Value{1, 2, 3, 4}.Uint8()
			}).To(Panic())
		})
	})
	Describe("Int8", func() {
		It("should generate the expected Value", func() {
			Expect(Int8(0)).To(Equal(Value{0}))
			Expect(Int8(43)).To(Equal(Value{43}))
			Expect(Int8(-42)).To(Equal(Value{214}))
		})
		It("should convert in both direction properly", func() {
			Expect(Int8(int8(0)).Int8()).To(Equal(int8(0)))
			Expect(Int8(int8(43)).Int8()).To(Equal(int8(43)))
			Expect(Int8(int8(-42)).Int8()).To(Equal(int8(-42)))
		})
		It("should panic when Value's length is incorrect", func() {
			Expect(func() {
				Value{1, 2}.Int8()
			}).To(Panic())
			Expect(func() {
				Value{1, 2, 3, 4}.Int8()
			}).To(Panic())
		})
	})
	Describe("Uint16", func() {
		It("should generate the expected Value", func() {
			Expect(Uint16(0)).To(Equal(Value{0, 0}))
			Expect(Uint16(132)).To(Equal(Value{132, 0}))
			Expect(Uint16(512)).To(Equal(Value{0, 2}))
		})
		It("should convert in both direction properly", func() {
			Expect(Uint16(uint16(0)).Uint16()).To(Equal(uint16(0)))
			Expect(Uint16(uint16(132)).Uint16()).To(Equal(uint16(132)))
			Expect(Uint16(uint16(512)).Uint16()).To(Equal(uint16(512)))
		})
		It("should panic when Value's length is incorrect", func() {
			Expect(func() {
				Value{1}.Uint16()
			}).To(Panic())
			Expect(func() {
				Value{1, 2, 3, 4}.Uint16()
			}).To(Panic())
		})
	})
	Describe("Int16", func() {
		It("should generate the expected Value", func() {
			Expect(Int16(0)).To(Equal(Value{0, 0}))
			Expect(Int16(132)).To(Equal(Value{132, 0}))
			Expect(Int16(512)).To(Equal(Value{0, 2}))
			Expect(Int16(-42)).To(Equal(Value{214, 255}))
		})
		It("should convert in both direction properly", func() {
			Expect(Int16(int16(0)).Int16()).To(Equal(int16(0)))
			Expect(Int16(int16(132)).Int16()).To(Equal(int16(132)))
			Expect(Int16(int16(512)).Int16()).To(Equal(int16(512)))
			Expect(Int16(int16(-42)).Int16()).To(Equal(int16(-42)))
		})
		It("should panic when Value's length is incorrect", func() {
			Expect(func() {
				Value{1}.Int16()
			}).To(Panic())
			Expect(func() {
				Value{1, 2, 3, 4}.Int16()
			}).To(Panic())
		})
	})
	Describe("Uint32", func() {
		It("should generate the expected Value", func() {
			Expect(Uint32(0)).To(Equal(Value{0, 0, 0, 0}))
			Expect(Uint32(132)).To(Equal(Value{132, 0, 0, 0}))
			Expect(Uint32(512)).To(Equal(Value{0, 2, 0, 0}))
		})
		It("should convert in both direction properly", func() {
			Expect(Uint32(uint32(0)).Uint32()).To(Equal(uint32(0)))
			Expect(Uint32(uint32(132)).Uint32()).To(Equal(uint32(132)))
			Expect(Uint32(uint32(512)).Uint32()).To(Equal(uint32(512)))
		})
		It("should panic when Value's length is incorrect", func() {
			Expect(func() {
				Value{1}.Uint32()
			}).To(Panic())
			Expect(func() {
				Value{1, 2}.Uint32()
			}).To(Panic())
			Expect(func() {
				Value{1, 2, 3, 4, 5}.Uint32()
			}).To(Panic())
		})
	})
	Describe("Int32", func() {
		It("should generate the expected Value", func() {
			Expect(Int32(0)).To(Equal(Value{0, 0, 0, 0}))
			Expect(Int32(132)).To(Equal(Value{132, 0, 0, 0}))
			Expect(Int32(512)).To(Equal(Value{0, 2, 0, 0}))
			Expect(Int32(-42)).To(Equal(Value{214, 255, 255, 255}))
		})
		It("should convert in both direction properly", func() {
			Expect(Int32(int32(0)).Int32()).To(Equal(int32(0)))
			Expect(Int32(int32(132)).Int32()).To(Equal(int32(132)))
			Expect(Int32(int32(512)).Int32()).To(Equal(int32(512)))
			Expect(Int32(int32(-42)).Int32()).To(Equal(int32(-42)))
		})
		It("should panic when Value's length is incorrect", func() {
			Expect(func() {
				Value{1}.Uint32()
			}).To(Panic())
			Expect(func() {
				Value{1, 2}.Uint32()
			}).To(Panic())
			Expect(func() {
				Value{1, 2, 3, 4, 5}.Uint32()
			}).To(Panic())
		})
	})
	Describe("Uint64", func() {
		It("should generate the expected Value", func() {
			Expect(Uint64(0)).To(Equal(Value{0, 0, 0, 0, 0, 0, 0, 0}))
			Expect(Uint64(132)).To(Equal(Value{132, 0, 0, 0, 0, 0, 0, 0}))
			Expect(Uint64(512)).To(Equal(Value{0, 2, 0, 0, 0, 0, 0, 0}))
		})
		It("should convert in both direction properly", func() {
			Expect(Uint64(uint64(0)).Uint64()).To(Equal(uint64(0)))
			Expect(Uint64(uint64(132)).Uint64()).To(Equal(uint64(132)))
			Expect(Uint64(uint64(512)).Uint64()).To(Equal(uint64(512)))
		})
		It("should panic when Value's length is incorrect", func() {
			Expect(func() {
				Value{1}.Uint64()
			}).To(Panic())
			Expect(func() {
				Value{1, 2}.Uint64()
			}).To(Panic())
			Expect(func() {
				Value{1, 2, 3, 4}.Uint64()
			}).To(Panic())
		})
	})
	Describe("Int64", func() {
		It("should generate the expected Value", func() {
			Expect(Int64(0)).To(Equal(Value{0, 0, 0, 0, 0, 0, 0, 0}))
			Expect(Int64(132)).To(Equal(Value{132, 0, 0, 0, 0, 0, 0, 0}))
			Expect(Int64(512)).To(Equal(Value{0, 2, 0, 0, 0, 0, 0, 0}))
			Expect(Int64(-42)).To(Equal(Value{214, 255, 255, 255, 255, 255, 255, 255}))
		})
		It("should convert in both direction properly", func() {
			Expect(Int64(int64(0)).Int64()).To(Equal(int64(0)))
			Expect(Int64(int64(132)).Int64()).To(Equal(int64(132)))
			Expect(Int64(int64(512)).Int64()).To(Equal(int64(512)))
			Expect(Int64(int64(-42)).Int64()).To(Equal(int64(-42)))
		})
		It("should panic when Value's length is incorrect", func() {
			Expect(func() {
				Value{1}.Uint64()
			}).To(Panic())
			Expect(func() {
				Value{1, 2}.Uint64()
			}).To(Panic())
			Expect(func() {
				Value{1, 2, 3, 4}.Uint64()
			}).To(Panic())
		})
	})
	Describe("Uintptr", func() {
		It("should generate the expected Value", func() {
			Expect(Uintptr(0)).To(Equal(Value{0, 0, 0, 0, 0, 0, 0, 0}))
			Expect(Uintptr(132)).To(Equal(Value{132, 0, 0, 0, 0, 0, 0, 0}))
			Expect(Uintptr(512)).To(Equal(Value{0, 2, 0, 0, 0, 0, 0, 0}))
		})
		It("should convert in both direction properly", func() {
			Expect(Uintptr(uintptr(0)).Uintptr()).To(Equal(uintptr(0)))
			Expect(Uintptr(uintptr(132)).Uintptr()).To(Equal(uintptr(132)))
			Expect(Uintptr(uintptr(512)).Uintptr()).To(Equal(uintptr(512)))
		})
		It("should panic when Value's length is incorrect", func() {
			Expect(func() {
				Value{1}.Uintptr()
			}).To(Panic())
			Expect(func() {
				Value{1, 2}.Uintptr()
			}).To(Panic())
			Expect(func() {
				Value{1, 2, 3, 4}.Uintptr()
			}).To(Panic())
		})
	})
	Describe("ByteSlice", func() {
		It("should generate the expected Value", func() {
			Expect(ByteSlice([]byte{1, 2, 3, 4, 5})).To(Equal(Value{1, 2, 3, 4, 5}))
			Expect(ByteSlice([]byte{})).To(Equal(Value{}))
		})
		It("should convert in both direction properly", func() {
			Expect(ByteSlice([]byte{1, 2, 3, 4, 5}).ByteSlice()).To(Equal([]byte{1, 2, 3, 4, 5}))
			Expect(ByteSlice([]byte{}).ByteSlice()).To(Equal([]byte{}))
		})
	})
	Describe("ZeroTermString", func() {
		It("should generate the expected Value", func() {
			Expect(ZeroTermString("test")).To(Equal(Value{'t', 'e', 's', 't', 0}))
			Expect(ZeroTermString("")).To(Equal(Value{0}))
		})
		It("should convert in both direction properly", func() {
			Expect(ZeroTermString("test").ZeroTermString()).To(Equal("test"))
			Expect(ZeroTermString("").ZeroTermString()).To(Equal(""))
		})
	})
})

func ExampleInt16() {
	intValue := Int16(-42)
	uintValue := intValue.Uint16()
	fmt.Println("uint64 representation of -42:", uintValue)
	// Output:
	// uint64 representation of -42: 65494
}

func ExampleZeroTermString() {
	value := ZeroTermString("Hello")
	fmt.Println("Value:", value)
	fmt.Println("Original:", value.ZeroTermString())
	// Output:
	// Value: [72 101 108 108 111 0]
	// Original: Hello
}
