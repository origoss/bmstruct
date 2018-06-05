package bmstruct

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Field", func() {
	Describe("IntField", func() {
		var f *Field
		var data []byte

		BeforeEach(func() {
			f = IntField("test", 4)
			data = make([]byte, 16)
			for i := range data {
				data[i] = byte(i)
			}
		})

		Context("when calling copyslice", func() {
			It("returns the correct value", func() {
				Expect(f.copySlice(data)).To(Equal([]byte{
					4, 5, 6, 7, 8, 9, 10, 11,
				}))

			})
			It("returns the copy of the data", func() {
				Expect(func() []byte {
					v := f.copySlice(data)
					v[0] = 64
					return data
				}()).To(Equal([]byte{
					0, 1, 2, 3, 4, 5, 6, 7,
					8, 9, 10, 11, 12, 13, 14, 15,
				}))

			})
		})
		Context("when calling updateSlice", func() {
			It("updates data as expected", func() {
				Expect(func() []byte {
					f.updateSlice(data,
						[]byte{42, 42, 42, 42, 42, 42, 42, 42})
					return data
				}()).To(Equal([]byte{
					0, 1, 2, 3, 42, 42, 42, 42,
					42, 42, 42, 42, 12, 13, 14, 15,
				}))
			})
		})
	})
	Describe("BitField", func() {
		var f *Field
		var data []byte
		Describe("creation", func() {
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

		Describe("usage", func() {
			BeforeEach(func() {
				f = BitField("test", 4, 2, 4)
				data = make([]byte, 16)
				for i := range data {
					data[i] = byte(i)
				}
				// data[5] = 4 = 0b00000100
				// bitfield is here  ^^^^
			})

			Context("when calling copyslice", func() {
				It("returns the correct value", func() {
					Expect(f.copySlice(data)).To(Equal([]byte{
						1,
					}))

				})
				It("returns the copy of the data", func() {
					Expect(func() []byte {
						v := f.copySlice(data)
						v[0] = 64
						return data
					}()).To(Equal([]byte{
						0, 1, 2, 3, 4, 5, 6, 7,
						8, 9, 10, 11, 12, 13, 14, 15,
					}))

				})
			})
			Context("when calling updateSlice", func() {
				It("updates data as expected", func() {
					Expect(func() []byte {
						// was 0b00000100
						//         ^^^^
						f.updateSlice(data,
							[]byte{10})
						// becomes 0b00101000 = 40
						//             ^^^^
						return data
					}()).To(Equal([]byte{
						0, 1, 2, 3, 40, 5, 6, 7,
						8, 9, 10, 11, 12, 13, 14, 15,
					}))
				})
				It("panics for multi-byte values", func() {
					Expect(func() {
						f.updateSlice(data,
							[]byte{10, 10})
					}).To(Panic())
				})
			})
		})
	})
})
