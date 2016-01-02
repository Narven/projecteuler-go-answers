package main

import "fmt"

func main() {

	var first int
	var second = 1
	var next int
	var total int

	for {
		next = first + second
		first = second
		second = next

		if next < 4000000 {

			if next%2 == 0 {
				total += next
			}
		} else {
			break
		}

	}

	fmt.Println(total)
}
