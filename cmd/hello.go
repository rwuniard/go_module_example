package main

import (
	"fmt"

	"github.com/rwuniard/go_module_example/pkg/dog"
)

func main() {
	dog_years := dog.Years(5)

	fmt.Println(dog_years)
}
