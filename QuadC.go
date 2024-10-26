package main

import (
	"fmt"
)

func QuadC(x, y int) {

	karakter := ' '

	for i := 1; i <= y; i++ {

		for j := 1; j <= x; j++ {

			if (i == 1 && j == 1) || (i == 1 && j == x) {

				karakter = 'A'

			} else if (i == y && j == 1) || (i == y && j == x) {

				karakter = 'C'

			} else if (i == 1 || i == y) || (j == 1 || j == x) {

				karakter = 'B'

			} else {

				karakter = ' '
			}
			fmt.Printf("%c", karakter)

		}

		fmt.Println()
	}
}

func main() {

	QuadC(5, 3)
	
}
