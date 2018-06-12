package bmstruct

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Struct", func() {
	var tmpl *Template
	var data Value

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
	Describe("creating struct with New", func() {
		Context("with improper data size", func() {
			It("should panic", func() {
				Expect(func() {
					tmpl.New(data[0:16])
				}).To(Panic())
			})
		})
		Context("with proper data size", func() {
			It("should succeed", func() {
				Expect(func() {
					tmpl.New(data)
				}).NotTo(Panic())
			})
		})
	})
	Describe("for an existing struct", func() {
		var s *Struct

		BeforeEach(func() {
			s = tmpl.New(data)
		})
		Describe("the Clone method", func() {
			var clone *Struct

			BeforeEach(func() {
				clone = s.Clone()
			})
			It("should create a struct with the same Template", func() {
				Expect(s.Template.Equal(clone.Template)).To(BeTrue())
			})
			It("should create a struct with the same data", func() {
				Expect(s.Value).To(Equal(clone.Value))
			})
			It("should create a struct with cloned data", func() {
				s.Value[0] = 0
				clone.Value[0] = 42
				Expect(s.Value[0]).To(Equal(byte(0)))
				Expect(clone.Value[0]).To(Equal(byte(42)))
			})
		})
		Describe("the Lookup method", func() {
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
					Expect(data).To(Equal(Value{
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
					Expect(data).To(Equal(Value{
						0, 1, 0, 1, 2, 3, 4, 5,
						6, 7, 10, 11, 12, 13, 14, 15,
						16, 17, 18, 19,
					}))
				})
			})

		})
	})
	Describe("Creating Structs with Slice", func() {
		var sliceData Value

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
			It("should not panic", func() {
				Expect(func() {
					tmpl.Slice(sliceData)
				}).NotTo(Panic())
			})
		})
	})
	Describe("for an existing Structs", func() {
		var ss *Structs
		var sliceData Value

		BeforeEach(func() {
			sliceData = make([]byte, 4*20)
			ss = tmpl.Slice(sliceData)
		})

		// Describe("the Offsets method", func() {
		// 	It("should return the correct offsets in order", func() {
		// 		Expect(ss.Offsets()).To(Equal([]uint64{
		// 			0, 20, 40, 60,
		// 		}))
		// 	})
		// })
		Describe("the Count method", func() {
			It("should report correct number of structs with Count", func() {
				Expect(ss.Count()).To(Equal(uint32(4)))
			})
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
				It("shall panic", func() {
					Expect(func() {
						ss.At(1)
					}).To(Panic())
					Expect(func() {
						ss.At(26)
					}).To(Panic())
					Expect(func() {
						ss.At(50)
					}).To(Panic())
					Expect(func() {
						ss.At(144)
					}).To(Panic())
					Expect(func() {
						ss.At(100)
					}).To(Panic())
				})
			})
		})
		Describe("the Nth method call", func() {
			Context("with proper ns", func() {
				It("shall return structs", func() {
					Expect(ss.Nth(0)).ToNot(BeNil())
					Expect(ss.Nth(1)).ToNot(BeNil())
					Expect(ss.Nth(2)).ToNot(BeNil())
					Expect(ss.Nth(3)).ToNot(BeNil())
				})
			})
			Context("with invalid n", func() {
				It("shall panic", func() {
					Expect(func() {
						ss.Nth(-1)
					}).To(Panic())
					Expect(func() {
						ss.Nth(4)
					}).To(Panic())
				})
			})
		})
		Describe("the Update method call", func() {
			var newS *Struct

			BeforeEach(func() {
				newS = tmpl.New(data)
			})
			Context("with invalid offset", func() {
				It("should panic", func() {
					Expect(func() {
						ss.Update(3, newS)
					}).To(Panic())
					Expect(func() {
						ss.Update(100, newS)
					}).To(Panic())
				})
			})
			Context("with invalid struct template", func() {
				It("should panic", func() {
					s2 := NewTemplate(25,
						IntField("f1", 0),
					).Empty()
					Expect(func() {
						ss.Update(0, s2)
					}).To(Panic())
				})
			})
			Context("with valid struct template", func() {
				It("should change the data", func() {
					ss.Update(0, newS)
					Expect(sliceData[0:25]).To(Equal(Value{
						0, 1, 2, 3, 4, 5, 6, 7,
						8, 9, 10, 11, 12, 13, 14, 15,
						16, 17, 18, 19, 0, 0, 0, 0,
						0,
					}))
				})
			})
		})
		Describe("the Clone method call", func() {
			var clone *Structs

			BeforeEach(func() {
				clone = ss.Clone()
			})
			It("shall generate a clone with the same template", func() {
				Expect(ss.Template.Equal(clone.Template)).To(BeTrue())
			})
			It("shall generate a clone with the same number of structs", func() {
				Expect(ss.Count()).To(Equal(clone.Count()))
			})
			It("shall generate a clone with independent data", func() {
				clone.Update(0, tmpl.New(Value{
					42, 42, 42, 42, 42, 42, 42, 42,
					42, 42, 42, 42, 42, 42, 42, 42,
					42, 42, 42, 42,
				}))
				Expect(clone.At(0).Value).To(Equal(Value{
					42, 42, 42, 42, 42, 42, 42, 42,
					42, 42, 42, 42, 42, 42, 42, 42,
					42, 42, 42, 42,
				}))
				Expect(sliceData).NotTo(ContainElement(byte(42)))
			})
		})
		Describe("the Iter method call", func() {
			It("shall iterate over all structs", func() {
				offsets := make([]uint64, ss.Count())
				structs := make([]*Struct, ss.Count())
				i := 0
				ss.Iter(func(offset uint64, s *Struct) {
					offsets[i] = offset
					structs[i] = s
					i++
				})
				Expect(offsets).To(Equal([]uint64{0, 20, 40, 60}))
				Expect(structs).NotTo(ContainElement(nil))
			})
		})
	})
})
