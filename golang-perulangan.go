package main

import "fmt"

func main() {

outerLoop:

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

	// for i := 0; i < 10; i++ {
	// 	for j := i; j < 10; j++ {
	// 		fmt.Print(j, "*")
	// 	}

	// 	fmt.Println()
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 != 1 {
	// 		continue
	// 	}

	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("Angka", i)
	// }

	// for {
	// 	fmt.Println("angka", i)

	// 	i++

	// 	if i == 10 {
	// 		break
	// 	}
	// }

	// for i < 5 {
	// 	fmt.Println("angka", i)
	// 	i++
	// }

	// for i := 0; i < 4; i++ {
	// 	fmt.Println("angka", i)
	// }
}
