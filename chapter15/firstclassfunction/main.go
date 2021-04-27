package main

import "fmt"

func sayHi() {
	fmt.Println("Hi")
}

func sayBye() {
	fmt.Println("Bye")
}

func divide(a int, b int) float64 {
	return float64(a) / float64(b)
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func main() {
	var greeterFunction func()
	var mathFunction func(int, int) float64
	greeterFunction = sayHi
	mathFunction = divide
	greeterFunction()
	fmt.Println(mathFunction(5, 2))
}
