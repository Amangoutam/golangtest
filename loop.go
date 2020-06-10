package main

import "fmt"

func main() {

	for j := 1; j < 100; j++ {
		//fmt.Println(j)
		if j%5 == 0 {
			if j%10 == 0 {
				fmt.Println(j / 5)
			} else if j%15 == 0 {
				fmt.Println(j / 3)
			} else {
				fmt.Println(j)
			}
		}
	}
}
