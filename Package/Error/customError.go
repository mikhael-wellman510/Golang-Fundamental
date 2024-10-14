package main

import "fmt"

type validationError struct {
	Message string
}

// Method validator
func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

// Method notFound
func (n *notFoundError) Error() string {
	return n.Message
}

func saveData(id string, data any) error {

	if id == "" {
		return &validationError{"Validation Error"}
	}

	if id != "eko" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {

	err := saveData("", nil)
	fmt.Println(err)

	cek := err.(*validationError)

}
