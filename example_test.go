package brwsr_test

import "github.com/speedyhoon/brwsr"

func ExampleOpen() {
	// Open locally hosted Godoc.
	err := brwsr.Open("https://localhost:6060")
	if err != nil {
		println(err)
	}
}
