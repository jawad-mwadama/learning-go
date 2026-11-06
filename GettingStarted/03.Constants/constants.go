package constants

import ("fmt"
"math")

// const declares a value
const s string="constant"

func GetConstants(){
	fmt.Println(s)

	// const statements can also appear on function body
	const n = 5000000000

	// constant expressions perform arithmetic with arbitrary precision
	const d= 3e20 / n
	fmt.Println(d)

	//  A numeric constant has no type until it is given one such as by an explicit conversion
	fmt.Println(int64(d))


	// A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64.
	fmt.Println(math.Sin(n))
}