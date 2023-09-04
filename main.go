package main

import (
	"fmt"

	"github.com/vdattada/toolkit"
)

func main() {
	var tools toolkit.Tools
	randomString := tools.RandomString(10)

	fmt.Println(randomString)
}
