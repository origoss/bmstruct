package bmstruct

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Template", func() {
	Describe("Creating template", func() {
		Context("with 0 size", func() {
			It("should panic", func() {
				Expect(func() {
					NewTemplate(0)
				}).To(Panic())
			})
		})
		Context("with negative size", func() {
			It("should panic", func() {
				Expect(func() {
					NewTemplate(-42)
				}).To(Panic())
			})
		})
		Context("without fields", func() {
			It("shall have minLen() = 0", func() {
				Expect(NewTemplate(42).minLen()).To(BeZero())
			})
		})
		Describe("with a field", func() {
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
	Describe("after created", func() {
		var t *Template

		BeforeEach(func() {
			t = NewTemplate(20)
		})
		It("should be report its size correctly", func() {
			Expect(t.Size).To(Equal(20))
		})
	})
})
