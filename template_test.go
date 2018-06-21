package bmstruct

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Template", func() {
	Describe("Creating template", func() {
		Context("with no fields", func() {
			It("should panic", func() {
				Expect(func() {
					NewTemplate(0)
				}).To(Panic())
			})
		})
		Context("with negative size", func() {
			It("should calculate the size from the fields", func() {
				t := NewTemplate(-1,
					IntField("f1", 0),
					IntField("f2", 8),
				)
				Expect(t.Size).To(Equal(16))
			})
		})
		Describe("with a single field", func() {
			Context("and too small size", func() {
				It("should panic", func() {
					Expect(func() {
						NewTemplate(4, IntField("f1", 0))
					}).To(Panic())
				})
			})
			Context("and exactly matching size", func() {
				It("should succeed", func() {
					Expect(func() {
						NewTemplate(8, IntField("f1", 0))
					}).ToNot(Panic())
				})
			})
			Context("and oversized", func() {
				It("should succeed", func() {
					Expect(func() {
						NewTemplate(12, IntField("f1", 0))
					}).ToNot(Panic())
				})
			})
		})
	})
	Describe("Comparing 2 templates", func() {
		var t1a, t1b, t2, t3 *Template

		BeforeEach(func() {
			t1a = NewTemplate(40,
				IntField("f1", 0),
				IntField("f2", 8),
			)
			t1b = NewTemplate(40,
				IntField("f1", 0),
				IntField("f2", 8),
			)
			t2 = NewTemplate(30,
				IntField("f1", 0),
				IntField("f2", 8),
			)
			t3 = NewTemplate(40,
				IntField("f1", 0),
			)
		})
		Context("which are equivalent", func() {
			It("should return true", func() {
				Expect(t1a.Equal(t1a)).To(BeTrue())
				Expect(t1a.Equal(t1b)).To(BeTrue())
			})
		})
		Context("which are different", func() {
			It("should return false", func() {
				Expect(t1a.Equal(t2)).NotTo(BeTrue())
				Expect(t1a.Equal(t3)).NotTo(BeTrue())
				Expect(t2.Equal(t3)).NotTo(BeTrue())
			})
		})
	})
	Describe("Getting the field at offset", func() {
		var t *Template

		BeforeEach(func() {
			t = NewTemplate(40,
				IntField("f1", 0),
				IntField("f2", 8),
			)
		})
		Context("when the offset is valid", func() {
			It("should return the proper fields", func() {
				Expect(t.FieldAt(0).Name).To(Equal("f1"))
				Expect(t.FieldAt(8).Name).To(Equal("f2"))
			})
		})
		Context("when the offset is invalid", func() {
			It("should panic", func() {
				Expect(func() {
					t.FieldAt(1)
				}).To(Panic())
				Expect(func() {
					t.FieldAt(10)
				}).To(Panic())
			})
		})
	})
})

func ExampleNewTemplate() {
	NewTemplate(40,
		IntField("ID", 0),
		IntField("Age", 8),
	)
}

func ExampleTemplate_Field() {
	//grades has 2 fields, the size is calculated
	grades := NewTemplate(-1,
		Uint8Field("math", 0),
		Uint8Field("science", 1),
	)

	//this new template has a field which is the template defined above
	NewTemplate(40,
		IntField("id", 0),
		IntField("age", 8),
		grades.Field("grades", 16),
	)
}
