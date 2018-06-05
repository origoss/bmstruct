package bmstruct

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBmstruct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bmstruct Suite")
}
