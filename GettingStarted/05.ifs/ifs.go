package ifs

import "fmt"

func Ifs(){

	if 7%2 ==0{
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%2==0{
		fmt.Println("8 is divisible by 4")
	}

	if 8%2==0 || 7%2==0 {
		fmt.Println("either are divisible 8 or 7 are even ")
	}

	if num :=9; num <0 {
		fmt.Println(num, "is negative")
	} else if num<10 {
		fmt.Println(num,"only has one digit")
	} else {
		fmt.Println(num,"has a multiple of 3 digits")
	}
}