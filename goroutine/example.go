package goroutine

func Fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func FibonacciWithChannel(n int) chan int {
	x, y := 0, 1
	c := make(chan int, n)

	go func() {
		for i := 0; i < n; i++ {
			x, y = y, x+y
			c <- x
		}
		close(c)
	}()

	return c
}
