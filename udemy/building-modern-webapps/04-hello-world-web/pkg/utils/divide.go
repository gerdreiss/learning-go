package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func divide(x, y float64) (float64, error) {
	return x / y, nil
}

func Divide(w http.ResponseWriter, r *http.Request) {
	p1 := r.PathValue("x")
	p2 := r.PathValue("y")

	x, err := strconv.ParseFloat(p1, 64)
	if err != nil {
		fmt.Fprintf(w, "Error parsing X param %s: %v", p1, err)
		return
	}

	y, err := strconv.ParseFloat(r.PathValue("y"), 64)
	if err != nil {
		fmt.Fprintf(w, "Error parsing Y param %s: %v", p2, err)
		return
	}

	if y == 0.0 {
		fmt.Fprintf(w, "division by zero")
		return
	}

	res, err := divide(x, y)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	fmt.Fprintf(w, "%f divided by %f is %f", x, y, res)
}
