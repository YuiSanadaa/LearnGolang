package main

import "fmt"

func main() {
	var positiveNumber uint8 = 89                                                         //1
	var negativeNumber = -123456789                                                       //2
	var decimalNumber = 2.52                                                              //3
	var exist bool = true                                                                 //4
	var message = `Nama Saya "Raihan Anubhawa Ekaputra", Salam Kenal And Koe "Jancok !".` //5

	fmt.Printf("Bilangan Positif: %d\n", positiveNumber)         //1
	fmt.Printf("Bilangan Negative: %d\n", negativeNumber)        //2
	fmt.Printf("Bilangan Decimal adalah: %f\n", decimalNumber)   //3
	fmt.Printf("Bilangan Decimal adalah: %.3f\n", decimalNumber) //3
	fmt.Printf("exist ? %t \n", exist)                           //4
	fmt.Println(message)                                         //5

}
