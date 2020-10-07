package main

import (
	"fmt"
	"strings"
)

func f1(a ...int) {
	fmt.Printf("%T\n", a)
	fmt.Printf("%#v\n", a)
}

func f2(a ...int) {
	a[0] = 50
}

func sumAndProduct(a ...float64) (float64, float64) {
	s := 0.
	p := 1.
	for _, v := range a {
		s += v
		p *= v
	}
	return s, p
}

func personInfo(age int, names ...string) string {
	fullName := strings.Join(names, `\/`)
	returnString := fmt.Sprintf("Age: %d Full name: %s", age, fullName)
	return returnString
}

func main() {
	f1(1, 2, 3, 4)
	f1()
	nums := []int{1, 2}
	nums = append(nums, 3, 4, 5)
	s, p := sumAndProduct(1.08, 6.66, 10.8, 66.6)
	fmt.Println(s, p)
	info := personInfo(78, "Lama", "Ole", "Nydahl")
	fmt.Println(info)

}
