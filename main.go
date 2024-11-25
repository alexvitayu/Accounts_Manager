package main

import "fmt"

type myType = [4]int

func main() {
	trans := [4]int{1, 2, 3, 4}
	fmt.Println(trans)
	reverse(&trans)
	fmt.Println(trans)

}

func reverse(array *myType) {
	for index, value := range *array {
		(*array)[len(*array)-1-index] = value
	}
}
