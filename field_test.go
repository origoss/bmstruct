package bmstruct

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Field", func() {
	var f *Field
	var bf *Field
	var data []byte

	BeforeEach(func() {
		f = IntField("test", 4)
		bf = BitField("testbf", 3, 1, 2)
		data = make([]byte, 16)
		for i := range data {
			data[i] = byte(i)
		}
		// data: 0: 00000000
		//       1: 00000001
		//       2: 00000010
		//       3: 00000011
		//               ^^    => bf
		//       4: 00000100 | => f
		//       5: 00000101 | => f
		//       6: 00000110 | => f
		//       ...         | => f
		//      11: 00001011 | => f
		//      12: 00001100
		//       ...
	})
	Describe("copySlice", func() {
		Context("for a regular Field", func() {
			It("should return the expected value", func() {
				Expect(f.copySlice(data)).To(Equal([]byte{
					4, 5, 6, 7, 8, 9, 10, 11,
				}))
			})
			It("should return a copy of the value", func() {
				v := f.copySlice(data)
				data[6] = 42
				Expect(v).To(Equal([]byte{
					4, 5, 6, 7, 8, 9, 10, 11,
				}))
			})
		})
		Context("for a bitfield", func() {
			It("should return the expected value", func() {
				Expect(bf.copySlice(data)).To(Equal([]byte{
					1,
				}))
			})
			It("should return a copy of the value", func() {
				v := bf.copySlice(data)
				data[3] = 255
				Expect(v).To(Equal([]byte{
					1,
				}))
			})
		})
	})
	Describe("updateSlice", func() {
		Context("for a regular field", func() {
			It("should modify the data at the expected slice", func() {
				f.updateSlice(data, Value{42, 42, 42, 42, 42, 42, 42, 42})
				Expect(data).To(Equal([]byte{
					0, 1, 2, 3, 42, 42, 42, 42,
					42, 42, 42, 42, 12, 13, 14, 15,
				}))
			})
			It("should panic when the input size is incorrect", func() {
				Expect(func() {
					f.updateSlice(data, Value{42, 42, 42, 42})
				}).To(Panic())
			})
		})
		Context("for a bitfield", func() {
			It("should modify the data at the expected byte", func() {
				bf.updateSlice(data, Value{3})
				Expect(data).To(Equal([]byte{
					0, 1, 2, 7, 4, 5, 6, 7,
					8, 9, 10, 11, 12, 13, 14, 15,
				}))
			})
			It("should panic when the input size is incorrect", func() {
				Expect(func() {
					bf.updateSlice(data, Value{42, 42})
				}).To(Panic())
			})
		})
	})
	Describe("BitField", func() {
		Describe("when created", func() {
			Context("with offset + length > 8", func() {
				It("should panic", func() {
					Expect(func() {
						BitField("test", 0, 3, 6)
					}).To(Panic())
				})
			})
			Context("with length == 0", func() {
				It("should panic", func() {
					Expect(func() {
						BitField("test", 0, 3, 0)
					}).To(Panic())
				})
			})
		})
	})
})
