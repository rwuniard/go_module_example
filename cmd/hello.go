package main

import (
	"fmt"

	"example.com/pkg/dog"
)

func main() {
	dog_years := dog.Years(5)

	fmt.Println(dog_years)
}
