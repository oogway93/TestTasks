package main

import (
	"fmt"
	"math"
)

func is_prime(x int) bool {
	if x <= 1 {
		return false
	}
	for i:=2;i<int(math.Sqrt(float64(x))+1);i++{
		if x % i == 0 {
			return false
		}
	}
	return true
}

func f3(x, y int) []int {
	var res []int
	for i:=x;i<=y;i++{
		if prime:=is_prime(i); prime {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	var input, input2 int
	_, err := fmt.Scan(&input, &input2)
	if err != nil {
		panic(err)
	}
	res := f3(input, input2)
	fmt.Println(res)
}