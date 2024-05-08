package main

import (
	"fmt"
)

func nod(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func f2(nums []int) []int {
	cmn := nums[0]
	for _, num := range nums[1:] {
		cmn = nod(cmn, num)
	}

	var res []int
	for i := 2; cmn >= i; i++ {
		if cmn%i == 0 {
			res = append(res, i)
		}
	}
	return res

}

func main() {
	var input []int
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}
	res := f2(input)
	fmt.Println(res)
}
