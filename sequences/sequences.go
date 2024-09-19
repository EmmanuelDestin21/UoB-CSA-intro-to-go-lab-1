package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func square(a int) int {
	return a * a
}

func double(slice []int) {
	slice = append(slice, slice...)
	fmt.Println(slice)
}

func mapSlice(f func(a int) int, slice []int) []int {
	result := make([]int, len(slice))
	for i, v := range slice {
		result[i] = f(v)
	}
	return result
}

func mapArray(f func(a int) int, array [3]int) [3]int {
	result := [3]int{}
	for i, v := range array {
		result[i] = f(v)
	}
	return result
}

func main() {
	intsSlice := []int{1, 2, 3, 4, 5}
	double(intsSlice)
	//intsArray := [3]int{1, 2, 3, 4, 5}
	//intsArray = mapArray(addOne, intsArray)
	//fmt.Println(intsArray)
}
