package goroutine_test

import (
	example "github.com/deiwin/golang-example/goroutine"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example", func() {
	Describe("fibonacci func", func() {
		fibonacci_numbers := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

		Context("with one fibonacci function", func() {
			var fib func() int

			BeforeEach(func() {
				fib = example.Fibonacci()
			})

			It("should return first 10 fibonacci numbers", func() {
				for _, number := range fibonacci_numbers {
					Expect(fib()).To(Equal(number))
				}
			})
		})

		Context("with two fibonacci functions", func() {
			var fib1, fib2 func() int

			BeforeEach(func() {
				fib1 = example.Fibonacci()
				fib2 = example.Fibonacci()
			})

			It("should return first 10 fibonacci numbers with separate closures", func() {
				Expect(fib1()).To(Equal(1))
				Expect(fib1()).To(Equal(1))
				Expect(fib1()).To(Equal(2))

				for _, number := range fibonacci_numbers {
					Expect(fib2()).To(Equal(number))
				}

				for _, number := range fibonacci_numbers[3:] {
					Expect(fib1()).To(Equal(number))
				}
			})
		})

		Context("with one fibonacci function with a channel", func() {
			var fib chan int

			BeforeEach(func() {
				fib = example.FibonacciWithChannel(10)
			})

			It("should return first 10 fibonacci numbers", func() {
				for _, number := range fibonacci_numbers {
					Expect(<-fib).To(Equal(number))
				}
			})
		})
	})
})
