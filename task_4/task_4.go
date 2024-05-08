package main

import "fmt"

func f4(x int) string {
	for i:=1;i<=x;i++{
		for j:=1;j<=x;j++{
			fmt.Printf("%3d ", j*i)
		}
		fmt.Printf("\n")
	}
	return ""
}

func main() {
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		panic(err)
	}
	res:= f4(input)
	fmt.Println(res)
}