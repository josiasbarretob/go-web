package main

import (
	"errors"
	"fmt"
)
const (
	dog = "dog"
	cat = "cat"
	hamster = "hamster"
	tarantula = "tarantula"
	)
	//dogFunc returns in grams
func dogFunc(quantity int) int {
	return quantity * 10000
}
	//catFunc returns in grams
func catFunc(quantity int) int {
	return quantity * 5000
}
	//hamsterFunc returns in grams
func hamsterFunc(quantity int) int {
	return quantity * 250
}
	//tarantulaFunc returns in grams
func tarantulaFunc(quantity int) int {
	return quantity * 150
}
	//Animal returns a func for a specific animal
func Animal(t string) (func(quantity int) int, error) {
	if t == dog {
		return dogFunc, nil
	}
	if t == cat {
		return catFunc, nil
	}
	
	if t == hamster {
		return hamsterFunc, nil
	}
	if t == tarantula {
		return tarantulaFunc, nil
	}
		return nil, errors.New("invalid animal type")
	}
func main() {
	dfunc, _ := Animal(dog)
	fmt.Printf("quandidade de alimento do(s) cachorro(s) (em gramas):%d gramas\n", dfunc(5))
	cfunc, _ := Animal(cat)
	fmt.Printf("quandidade de alimento do(s) gato(s) (em gramas): %d gramas\n", cfunc(8))
	_, err := Animal("invalid animal")
	if err != nil {
	fmt.Println("animal inválido")
	}
}