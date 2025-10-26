package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {
	fmt.Println("function sqrt ")
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("function power ")
	_, err = power(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {

	if f < 0 {
		return 0, errors.New(" square root of negative number")
	}
	return 42, nil
}
func power(f float64) (float64, error) {
	if f < 0 {
		ErrNegativeMath := fmt.Errorf("power of negative number %v", f)
		return 0, ErrNegativeMath
	}
	return 42, nil
}
