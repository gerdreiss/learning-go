package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	j := `[
		     {
		        "first_name": "Clark",
		        "last_name": "Kent",
		        "hair_color": "black",
		        "has_dog": true
		     },
		     {
		        "first_name": "Bruce",
		        "last_name": "Wayne",
		        "hair_color": "black",
		        "has_dog": false
		     }
		  ]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(j), &unmarshalled)
	if err != nil {
		fmt.Println("Error unmarshalling json", err)
	}

	fmt.Printf("unmarshalled: %v\n\n", unmarshalled)

	marshalled, err := json.MarshalIndent(unmarshalled, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling slice of persons", err)
	}

	fmt.Printf("marshalled: %v\n\n", string(marshalled))
}

func divide(x, y float32) (float32, error) {
	if y == 0.0 {
		return 0.0, errors.New("division by zero")
	}

	return x / y, nil
}
