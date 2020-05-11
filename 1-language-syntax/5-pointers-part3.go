package main

// Escape analysis
// go build -gcflags "-m -m"

type user struct {
	name  string
	email string
}

func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	println("u1", &u1, "u2", &u2)

}

func createUserV1() user {
	u := user{
		name:  "Shyam",
		email: "shyam.vegeta@gmail.com",
	}

	println("V1", &u)

	return u
}

func createUserV2() *user {
	u := user{
		name:  "Shyam",
		email: "shyam.vegeta@gmail.com",
	}

	println("V2", &u)

	return &u
}
