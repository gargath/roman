package roman

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRoman(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Roman Numerals Test Suite")
}
