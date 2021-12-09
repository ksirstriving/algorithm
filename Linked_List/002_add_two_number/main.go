package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {

	}
	i:=1;
	for i < 10 {
		i++
	}
	for k, v := range [5]int{1,2,3,4,5} {
		fmt.Println(k,v)
	}
}