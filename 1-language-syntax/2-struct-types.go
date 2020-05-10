package main

import (
	"fmt"
)

type example struct {
	pi      float32
	counter int16
	flag    bool
}

type alice struct {
	flag    bool
	counter int16
	pi      float32
}

type queen struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// Declare a variable of type example set to its zero value
	var ex example

	// Display the value
	fmt.Printf("%-v\n", ex)

	// Deaclare a variable of type example and init using
	// a struct literal.
	e1 := example{
		pi:      3.141592,
		counter: 10,
		flag:    true,
	}

	// Declare a variable of a anonymous type set to its zero value.
	var e2 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Declare a variable of an anonymous type and init using a struct literal.
	e3 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the values
	fmt.Printf("%+v\n", e3)
	fmt.Printf("Flag", e3.flag)
	fmt.Printf("Counter", e3.counter)
	fmt.Printf("Pi", e3.pi)

	var a queen
	var b alice

	// b = a This wont work. Go wont do implicit conversion.
	b = alice(a)
	fmt.Print(b, a)

	b = e2 // This will work bcoz e2 is a literal type not a named type.
	fmt.Print(b, e3)
}
