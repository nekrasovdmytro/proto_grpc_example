package main

import (
	"fmt"
	"improve/model"
)

func main() {
	fmt.Print()

	car := model.Car{
		Type: "ZBRW",
		Year: 111,
	}

	fmt.Println(car)
}
