package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		isMult3 := i%3 == 0
		isMult5 := i%5 == 0

		if isMult3 && isMult5 {
			fmt.Println("Pin Pan")
		} else if isMult3 {
			fmt.Println("Pin")
		} else if isMult5 {
			fmt.Println("Pan")
		} else {
			fmt.Println(i)
		}
	}
}
