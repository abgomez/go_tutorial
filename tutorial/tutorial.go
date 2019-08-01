package main

import (
	"fmt"
	"math"
	"time"
)

const s string = "constant"

//gobyexample.com
func main() {
	// Values, just an example of the values of GO
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 = ", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// variables
	var a = "initial" //no need to declare type
	fmt.Println(a)
	var b, c int = 1, 2 //multiple variables
	fmt.Println(b, c)
	var d = true
	fmt.Println(d)
	var e int //intialize to 0
	fmt.Println(e)
	f := "apple" //declare and initialize a variable, without the var keyword
	fmt.Println(f)

	//constants
	fmt.Println(s)
	const n = 50000000 //no type until nedeed
	const z = 3e20 / n
	fmt.Println(z)
	fmt.Println(int64(z))
	fmt.Println(math.Sin(n))

	//for
	for n := 0; n <= 5; n++ {
		if n%2 == 0 { // any variable declared in this statement are available in all branches.
			continue
		}
		fmt.Println(n)
	}

	//switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: //two clauses one case
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	whatAmI := func(i interface{}) { //swith as a type
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	//slices similar to arrays
	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	s = append(s, "e", "f")
	fmt.Println(s)
	l := s[2:5]
	fmt.Println(l)
}
