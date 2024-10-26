import (
	"fmt"
)

func QuadA(x, y int) {

	karakter := ' '

	for i := 1; i <= y; i++ {

		for j := 1; j <= x; j++ {

			if (i == 1 && j == 1) || (i == y && j == 1) || (i == 1 && j == x) || (i == y && j == x) {

				karakter = 'o'

			} else if i == 1 || i == y {
				karakter = '-'

			} else if j == 1 || j == x {
				karakter = '|'

			} else {
				karakter = ' '
			}
			fmt.Printf("%c", karakter)

		}

		fmt.Println()
	}

}
func main() {

	QuadA(5, 3)

}
