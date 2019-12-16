package main

import "fmt"

func main() {
	// var fruits = []string{"Apple", "Grape", "Banana", "Melon"}
	// fmt.Println(fruits[3])

	// var fruitsA = []string{"apple", "grape"}
	var fruitsB = []string{"apple", "grape", "teh sisri", "teh jus"}
	// var fruitsC = []string{"apple", "grape"}
	fmt.Println(fruitsB[2])

	var fruits = []string{"apple", "grape", "banana", "melon"}

	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]

	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(fruits)   // [apple grape banana melon]
	fmt.Println(aFruits)  // [apple grape banana]
	fmt.Println(bFruits)  // [grape banana melon]
	fmt.Println(aaFruits) // [grape]
	fmt.Println(baFruits) // [grape]
	fmt.Println(len(baFruits))
	fmt.Println(cap(baFruits))

	// // Buah "grape" diubah menjadi "pinnaple"
	// baFruits[0] = "pinnaple"

	// fmt.Println(fruits)   // [apple pinnaple banana melon]
	// fmt.Println(aFruits)  // [apple pinnaple banana]
	// fmt.Println(bFruits)  // [pinnaple banana melon]
	// fmt.Println(aaFruits) // [pinnaple]
	// fmt.Println(baFruits) // [pinnaple]
}
