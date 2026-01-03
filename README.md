# brwsr

[![Go Reference](https://pkg.go.dev/badge/github.com/speedyhoon/brwsr.svg)](https://pkg.go.dev/github.com/speedyhoon/brwsr)
[![Go Report Card](https://goreportcard.com/badge/github.com/speedyhoon/brwsr)](https://goreportcard.com/report/github.com/speedyhoon/brwsr)

`brwsr` provides a way to open URL links in the user's default browser, compatible with most operating systems.

## Example

```go
package main

import "github.com/speedyhoon/brwsr"

func main() {
	// Open locally hosted Godoc.
	err := brwsr.Open("https://localhost:6060")
	if err != nil {
		println(err)
	}
}
```
