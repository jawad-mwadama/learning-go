package main

import (
	values "example/hello/01.values"
	variables "example/hello/02.variables"
	constants "example/hello/03.constants"
	loops "example/hello/04.For"
	ifs "example/hello/05.ifs"
	"fmt"
)

func main() {

	fmt.Println("Hello, World!")

	values.PrintValues()
	variables.GoVariables()
	constants.GetConstants()
	loops.ForLoops()
	ifs.Ifs()
}
