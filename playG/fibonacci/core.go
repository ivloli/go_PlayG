package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", fibo1(i))
	}
	fmt.Println("")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", fibo2(i))
	}
	fmt.Println("")
	out := make(chan int)
	quit := make(chan struct{})
	go fibo3(out, quit)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", <-out)
	}
	fmt.Println("")
	quit <- struct{}{}
	f := fibo4()
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", f())
	}
}

func fibo1(n int) int {
	a1 := 1
	a2 := 1
	if n == 1 {
		return a1
	}
	if n == 2 {
		return a2
	}

	var res int
	for i := 3; i <= n; i++ {
		res = a1 + a2
		a1 = a2
		a2 = res
	}
	return res
}
func fibo2(n int) int {
	x, y := 1, 1
	var res int
	for i := 1; i <= n; i++ {
		res = x
		x, y = y, x+y
	}
	return res
}
func fibo3(out chan<- int, quit <-chan struct{}) {
	x, y := 1, 1
	for {
		select {
		case out <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("f3 quit")
			return
		}
	}
}

func fibo4() func() int {
	x, y := 1, 1
	return func() int {
		res := x
		x, y = y, x+y
		return res
	}
}
