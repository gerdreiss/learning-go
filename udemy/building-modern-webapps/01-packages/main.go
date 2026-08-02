package main

import (
	"fmt"

	"github.com/gerdreiss/packages/helpers"
)

func main() {
	var dt helpers.DataType
	dt.TypeName = "Int"
	dt.TypeNumber = 1
	fmt.Println(dt)
}
