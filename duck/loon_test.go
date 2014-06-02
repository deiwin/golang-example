package duck_test

import (
	. "github.com/deiwin/golang-example/duck"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Loon", func() {

	Context("with a loon pointer", func() {
		var loon = new(Loon)

		It("should quack", func() {
			Expect(loon.Quack()).To(Equal("Quuuaack!"))
		})

		It("should swim", func() {
			loon.Swim()

			Expect(loon.IsSwimming).To(BeTrue())
		})

		It("should walk", func() {
			loon.Walk()

			Expect(loon.IsSwimming).To(BeFalse())
		})
	})

	Context("with a loon object", func() {
		var loon = Loon{false}

		It("should NOT swim", func() {
			loon.SwimWithoutPointer()

			Expect(loon.IsSwimming).NotTo(BeTrue())
		})
	})
})
