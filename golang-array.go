package main

import "fmt"

func main() {

	var fruits = make([]string, 2)
	fruits[0] = "apple"
	fruits[1] = "manggo"

	fmt.Println(fruits)

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for _, fruit := range fruits {
	// 	fmt.Printf("nama buah : %s\n", fruit)
	// }

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
	// }

	// var numbers1 = [2][3]int{[3]int{3, 2, 3}, [3]int{3, 4, 5}}
	// var numbers2 = [2][3]int{{3, 2, 3}, {3, 4, 5}}

	// fmt.Println("numbers1", numbers1)
	// fmt.Println("numbers2", numbers2)

	// var numbers = [...]int{2, 3, 2, 4, 3}
	// fmt.Println("data array :", numbers)
	// fmt.Println("Jumlah elemen \t:", len(numbers))

	// var drunk = [4]string{"mansion", "iceland", "jackdaniel", "vodka"}

	// fmt.Println("Jumlah element \t\t", len(drunk))
	// fmt.Println("Jumlah element \t\t", drunk)

	// var names [4]string
	// names[0] = "Raihan "
	// names[1] = "Anubhawa "
	// names[2] = "Eka "
	// names[3] = "Putra "

	// fmt.Print(names[0], names[1], names[2], names[3])
}
