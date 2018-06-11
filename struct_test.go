package bmstruct

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Struct", func() {
	var tmpl *Template
	var data []byte

	BeforeEach(func() {
		tmpl = NewTemplate(20,
			IntField("field1", 2),
			IntField("field2", 10),
		)
		data = make([]byte, 20)
		for i := range data {
			data[i] = byte(i)
		}
	})
	Describe("Creating structs", func() {
		var sliceData []byte

		BeforeEach(func() {
			sliceData = make([]byte, 4*20)
		})
		Context("with mismatching data size", func() {
			It("should panic", func() {
				Expect(func() {
					tmpl.Slice(sliceData[0:75])
				}).To(Panic())
			})
		})
		Context("with proper data size", func() {
			var ss *Structs

			BeforeEach(func() {
				ss = tmpl.Slice(sliceData)
			})

			It("should return the correct Offsets in order", func() {
				Expect(ss.Offsets()).To(Equal([]uint64{
					0, 20, 40, 60,
				}))
			})
			It("should report correct number of structs with Count", func() {
				Expect(ss.Count()).To(Equal(uint32(4)))
			})
			Describe("the At method call", func() {
				Context("with proper offsets", func() {
					It("shall return structs", func() {
						Expect(ss.At(0)).ToNot(BeNil())
						Expect(ss.At(20)).ToNot(BeNil())
						Expect(ss.At(40)).ToNot(BeNil())
						Expect(ss.At(60)).ToNot(BeNil())
					})
				})
				Context("with invalid offsets", func() {
					It("shall return nil", func() {
						Expect(ss.At(1)).To(BeNil())
						Expect(ss.At(26)).To(BeNil())
						Expect(ss.At(50)).To(BeNil())
						Expect(ss.At(144)).To(BeNil())
					})
				})
			})
			Describe("when cloned", func() {
				var clone *Structs

				BeforeEach(func() {
					clone = ss.Clone()
				})
				It("shall generate a clone with the same number of structs", func() {
					Expect(ss.Count()).To(Equal(clone.Count()))
				})
				It("shall generate a clone with independent data", func() {
					Expect(func() []byte {
						clone.At(0).Update("field1", Value{
							42, 42, 42, 42,
							42, 42, 42, 42})
						return sliceData
					}()).NotTo(ContainElement(byte(42)))
					Expect(func() []byte {
						clone.At(0).Update("field1", Value{
							42, 42, 42, 42,
							42, 42, 42, 42})
						return clone.Data
					}()).To(ContainElement(byte(42)))
				})
			})
			Describe("when updated", func() {
				var newS *Struct

				BeforeEach(func() {
					newS = tmpl.New(data)
				})
				Context("with invalid offset", func() {
					It("should panic", func() {
						Expect(func() {
							ss.Update(3, newS)
						}).To(Panic())
					})
				})
				Context("with invalid struct template", func() {
					var tmpl2 *Template
					var s2 *Struct

					BeforeEach(func() {
						tmpl2 = NewTemplate(25,
							IntField("f1", 0),
						)
						s2 = tmpl2.Empty()
					})

					It("should panic", func() {
						Expect(func() {
							ss.Update(0, s2)
						}).To(Panic())
					})
				})
				Context("with valid struct template", func() {
					BeforeEach(func() {
						ss.Update(0, newS)
					})
					It("should change the data", func() {
						Expect(sliceData[0:25]).To(Equal([]byte{
							0, 1, 2, 3, 4, 5, 6, 7,
							8, 9, 10, 11, 12, 13, 14, 15,
							16, 17, 18, 19, 0, 0, 0, 0,
							0,
						}))
					})
				})
			})
		})
	})
	Describe("Creating struct", func() {
		Context("with too small data", func() {
			It("should panic", func() {
				Expect(func() {
					tmpl.New(data[0:15])
				}).To(Panic())
			})
		})
		Context("with too large data", func() {
			It("should panic", func() {
				data = make([]byte, 50)
				Expect(func() {
					tmpl.New(data)
				}).To(Panic())
			})
		})
		Context("with matching data size", func() {
			It("should not panic", func() {
				Expect(func() {
					tmpl.New(data)
				}).ToNot(Panic())
			})
		})
	})
	Describe("Struct operations", func() {
		var s *Struct

		BeforeEach(func() {
			s = tmpl.New(data)
		})
		Describe("Data method", func() {
			It("should return the stored data", func() {
				Expect(s.Data).To(Equal(data))
			})
		})
		Describe("Clone method", func() {
			var clone *Struct

			BeforeEach(func() {
				clone = s.Clone()
			})
			It("should create a struct with the same Template", func() {
				Expect(s.Template.Equal(clone.Template)).To(BeTrue())
			})
			It("should create a struct with the same data", func() {
				Expect(s.Data).To(Equal(clone.Data))
			})
			It("should create a struct with cloned data", func() {
				s.Data[0] = 0
				clone.Data[0] = 42
				Expect(s.Data[0]).To(Equal(byte(0)))
				Expect(clone.Data[0]).To(Equal(byte(42)))
			})
		})
		Describe("Lookup operation", func() {

			Context("for non-existing key", func() {
				It("should panic", func() {
					Expect(func() {
						s.Lookup("no-such-field")
					}).To(Panic())
				})
			})
			Context("for an existing key", func() {
				It("should return the correct value", func() {
					Expect(s.Lookup("field1")).To(Equal(Value([]byte{2, 3, 4, 5, 6, 7, 8, 9})))
				})
				It("should return a cloned value", func() {
					v := s.Lookup("field1")
					v[0] = 4
					Expect(data).To(Equal([]byte{
						0, 1, 2, 3, 4, 5, 6, 7,
						8, 9, 10, 11, 12, 13, 14, 15,
						16, 17, 18, 19,
					}))
				})
			})
		})
		Describe("Update operation", func() {
			var v Value

			BeforeEach(func() {
				v = make([]byte, 8)
				for i := range v {
					v[i] = byte(i)
				}
			})

			Context("for non-existing key", func() {
				It("should panic", func() {
					Expect(func() {
						s.Update("no-such-field", v)
					}).To(Panic())
				})
			})
			Context("for too small key", func() {
				It("should panic", func() {
					Expect(func() {
						s.Update("field1", v[0:7])
					}).To(Panic())
				})
			})
			Context("for too large key", func() {
				It("should panic", func() {
					Expect(func() {
						s.Update("field1", Value(make([]byte, 9)))
					}).To(Panic())
				})
			})
			Context("with a proper key and value", func() {
				It("shall change the data", func() {
					s.Update("field1", v)
					Expect(data).To(Equal([]byte{
						0, 1, 0, 1, 2, 3, 4, 5,
						6, 7, 10, 11, 12, 13, 14, 15,
						16, 17, 18, 19,
					}))
				})
			})

		})
	})
})
