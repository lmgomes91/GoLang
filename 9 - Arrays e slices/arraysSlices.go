package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("arrays e slices")

	var array1 [5]int8
	fmt.Println(array1)

	array2 := [...]int8{1, 2, 3, 4}
	fmt.Println(array2)

	slice := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 7)
	fmt.Println(slice)

	slice2 := array2[0:2]
	fmt.Println(slice2)

	array2[0] = -1
	fmt.Println(slice2)

	//arrays internos
	fmt.Println("**************")
	slice3 := make([]int, 3, 4)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2)

	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	fmt.Println("**************")

	slice4 := make([]int, 3)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
	slice4 = append(slice4, 2)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	fmt.Println("**************")
	// fazer um cópia entre dois slices
	slice5 := make([]int, len(slice4))
	copy(slice5, slice4)
	slice5[0] = -1
	fmt.Println(slice4)
	fmt.Println(slice5)

}
