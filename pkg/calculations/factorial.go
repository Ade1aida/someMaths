package calculations

import (
	"fmt"
	"os"
)

func Factorial(n int64, logir bool) int64 {
	if n > 0 {
		var factorial int64 = 1
		if logir == true {
			fmt.Printf("Start calculations... \n Calculate %v!\n", n)
			for n != 1 {
				factorial = factorial * n
				n -= 1
			}
			fmt.Println("Calculations complete!")
		} else {
			for n != 1 {
				factorial = factorial * n
				n -= 1
			}
		}

		return factorial
	} else {
		if n == 0 {
			return 1
		} else {
			fmt.Println("ERROR! n MUST BE >= 0")
			os.Exit(1)
			return 0
		}
	}

}
