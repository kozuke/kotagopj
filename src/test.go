package main

import "fmt"

func main() {
	// testmake()
	// testmap()
	// testbyte()
	r1, r2 := add(10, 20)
	fmt.Println(r1, r2)

	f := func(x int) {
		fmt.Println("inner func", x)
	}
	f(2)

	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}

func testmake() {
	// c := make([]int, 5)
	c := make([]int, 0, 5)
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)

}

func testmap() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)

	v, ok := m["apple"]
	fmt.Println(v, ok)
}

func testbyte() {
	b := []byte{72, 73}
	fmt.Println(string(b))
}

func add(x, y int) (int, int) {
	fmt.Println("add function")
	return x + y, x - y
}

// named return value
func cal(price, item int) (result int) {
	result = price * item
	return // resultを返してくれる　任意の文字を与えられる　　// ネイキッドリターン
}

// Closure
func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
