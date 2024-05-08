package main

import (
	"fmt"
)

func f(x int) string {
	rod := ""
	if x%10 == 1 && x%100 != 11 {
		rod="компьютер"
	} else if x%10 >= 2 && x%10 <= 4 {
		rod="компьютера"
	} else if (x%10 >= 5 && x%10 <= 9) || (x%10==0){
		rod="компьютеров"
	}
	res := fmt.Sprintf("%d %s", x, rod)
	return res
}

func main() {
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}
	res := f(input)
	fmt.Println(res)
}
