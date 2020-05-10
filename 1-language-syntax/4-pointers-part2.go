package main

// Pointer Semantics, only need to share data across boundaries.

func main() {
	// Declare variable of type int with a value of 10
	count := 10
	// Display the "value of" and "address of" count.
	println("count: \tValue Of [", count, "]\tAddr Of[", &count, "]")

	// Pass the "address of" the count.
	increment(&count)

	println("count: \tValue Of [", count, "]\tAddr Of[", &count, "]")
}

func increment(inc *int) {
	// Increment the "value at the address" in inc.
	*inc++
	println("count: \tValue Of [", inc, "]\tAddr Of[", &inc, "]")
}
