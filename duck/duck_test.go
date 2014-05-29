package duck_test

import (
	. "github.com/deiwin/golang-example/duck"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Duck", func() {
	Context("with a loon", func() {
		var loon = new(Loon)

		It("should act like a duck", func() {
			Expect(FeedBreadcrumbsTo(loon)).To(Equal("Quuuaack!"))
		})
	})
})
