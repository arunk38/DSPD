// Go program to illustrate type switch
package main

import (
	"fmt"
)

func myfun(a interface{}) {

	// Using type switch
	switch a.(type) {

	case int:
		fmt.Println("Type: int, Value:", a.(int))
	case string:
		fmt.Println("\nType: string, Value: ", a)
	case float64:
		fmt.Println("\nType: float64, Value: ", a.(float64))
	case []int:
		fmt.Println("\nType: []int, Value: ", a.([]int))
	default:
		fmt.Println("\nType not found")
	}
}

// Main method
func main() {

	myfun("1")
	myfun(67.9)
	myfun(true)
	myfun([]int{1, 2, 3})
	var val interface {
	} = []int{1, 2, 3}

	fmt.Println(val.([]int))

	favSubject := "1"

	fmt.Printf("It's %s\n", favSubject)
}
