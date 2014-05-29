package goroutine_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestGoroutine(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Goroutine Suite")
}
