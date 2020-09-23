package main

import "fmt"

func main() {
	numbers := []int{2, 3}
	numbers = append(numbers, 10)
	fmt.Println(numbers)
	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers)
	n := []int{100, 200, 300}
	numbers = append(numbers, n...)
	fmt.Println(numbers)

	nums := []int{1}
	fmt.Printf("Lenght: %d, capacity: %d\n", len(nums), cap(nums))
	nums = append(nums, 2)
	fmt.Printf("Lenght: %d, capacity: %d\n", len(nums), cap(nums))
	nums = append(nums, 3)
	fmt.Printf("Lenght: %d, capacity: %d\n", len(nums), cap(nums))
	nums = append(nums, 4, 5)
	fmt.Printf("Lenght: %d, capacity: %d\n", len(nums), cap(nums))

	src := []int{10, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Println(src, dst, nn)
}
