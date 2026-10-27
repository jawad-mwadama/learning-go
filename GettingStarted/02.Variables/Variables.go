package Variables

import "fmt"

func GoVariables() {
	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// one can declare multiple variables at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go will infer types of initialized variables
	var d = true
	fmt.Println(d)

	// variables declared withut a corresponding initialization are zero-valued. eg the value of int is 0
	var e int
	fmt.Println(e)

	// the := syntax is shorthand dor declaring or initializinga variable. eg for var f string = "apple" in this case. This syntax is only available inside functions
	f := "apple"
	fmt.Println(f)
}
