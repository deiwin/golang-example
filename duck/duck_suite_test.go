package duck_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDuck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Duck Suite")
}
