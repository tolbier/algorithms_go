package multiplier_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMultiplier(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Internal Tests")
}
