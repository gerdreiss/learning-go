package main

import (
	"errors"
	"fmt"
	"net/http"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	x := 100.0
	y := 10.0
	res, err := divide(x, y)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	} else {
		fmt.Fprintf(w, "%f divided by %f is %f", x, y, res)
	}
}
