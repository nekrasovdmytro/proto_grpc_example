package main

import (
	"fmt"
)

func main() {
	t := `
To run api:
$ go run api/main.go

To run service:
$ go run service/main.go

`

	fmt.Printf(t)
}
