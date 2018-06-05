package bmstruct

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Value", func() {
	Describe("Uint64", func() {
		It("should convert in both direction properly", func() {
			Expect(Uint64(uint64(0)).Uint64()).To(Equal(uint64(0)))
			Expect(Uint64(uint64(1324)).Uint64()).To(Equal(uint64(1324)))
		})
	})
})
