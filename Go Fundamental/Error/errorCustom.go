package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

func (c *CustomError) Succes() string {

	return fmt.Sprintf("Succes code %d , With Message %s", c.Code, c.Message)

}

func pembagian(val int, pembagi int) (int, error) {

	if pembagi != 0 {
		return val / pembagi, nil
	} else {
		return 0, errors.New("Failed Boss")
	}
}

func main() {
	succes := &CustomError{202, "Succes"}

	if hasil, res := pembagian(10, 2); res == nil {
		fmt.Println(succes.Succes(), " dan hasil nya : ", hasil)
	} else {
		fmt.Println(res)
	}
}
