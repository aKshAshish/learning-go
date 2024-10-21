package main

import "fmt"

func alterValue(arr []int) {
	arr[0] = -1
}

func alterSlice(s []int) {
	s = append(s, 100)
}

func alterSliceByRef(s *[]int) {
	*s = append(*s, 200)
}

func loopOverSliceRef(s *[]int) {
	fmt.Print("[")
	for idx, val := range *s {
		fmt.Printf("(%d, %d)", idx, val)
		if idx != len(*s)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")
}

func Slice() {
	// initializing slice in various ways
	// using array slicing
	arr := [10]int{2, 4, 1, 6, 7, 9, 11, 15, 10, 0}
	s := arr[3:8] // [3:15] will cause panic as the slicing is beyond the length of underlying array
	fmt.Println(arr, s)

	// using make function
	// will create an underlying array with capacity 20 and initialize the first 10 elements to the
	// zero value of the type and return it as the alice
	smake := make([]int, 10, 20)
	fmt.Println(len(smake), cap(smake), smake) // 10, 20, [0...10times]
	// fmt.Print(smake[10]) will cause panic since we are indexing out of range
	// re-slicing the array will allow us to access all the elements of underlying array
	smake = smake[:cap(smake)]
	fmt.Println(len(smake), cap(smake), smake) // 20, 20, [0...20times]

	// altering the value of slice will change the value in underlying array as well
	// since slice is already a pointer to array we can pass it as value we do not need to pass it as a reference
	alterValue(s)
	fmt.Println(arr, s) // [2 4 1 -1 7 9 11 15 10 0] [-1 7 9 11 15]

	// Note: Although the changes made to element of slice are reflected here but if
	// we were to make changes to slice itself say appending new values those won't get reflected if
	// slice is passed as a value.
	alterValue(smake)
	fmt.Println(len(smake), cap(smake), smake) // 20, 20, [0...20times]; no changes
	// if we pass by reference the capacity of the underlying array increases and the appended values are reflected in the slice
	// increasing the length of the slice by lenght of appended values
	alterSliceByRef(&smake)
	fmt.Println(len(smake), cap(smake), smake) // 20, 400, [0...20times, 100]

	// Note:
	// Something interesting happens if we try passing the slice that was created from slicing an array,
	// The appended values are reflected in the array iff the length of appended values do not cross the length of the array,
	// on which the slice was created. This behaviour makes sense if we understand how append is written.
	// Append first checks if appending the value will surpass the capacity of the array or not. If the capacity it not surpassed it
	// simply appends the value which causes the underlying array to reflect the changes. But if the capacity is surpassed
	// the append function will create a new array with twice the capacity and copy the exisiting values of slice and then appending
	// the values. Since overflow creates a new underlying array for the slice passed those appends don't get reflected in the underlying
	// array and slice here.
	alterSlice(s)
	fmt.Println(arr, s) // [2 4 1 -1 7 9 11 15 100 0] [-1 7 9 11 15]; notice the value has changed after 15 to 100, earlier it was 10.
	// In this case if the overflow does happen since slice is passed by reference the slice will get updated,
	// but the underlying array will not get affected
	alterSliceByRef(&s)
	fmt.Println(arr, s)

	fmt.Print("[")
	// Looping over an arr/slice
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d", arr[i])
		if i != len(arr)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")

	fmt.Print("[")
	for idx, val := range arr {
		fmt.Printf("(%d, %d)", idx, val)
		if idx != len(arr)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println("]")

	// Looping over slice ref
	loopOverSliceRef(&s)
}
