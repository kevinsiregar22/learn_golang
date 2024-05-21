package main

import "fmt"

func main(){
	// deklarasi variabel

	var name string = "Kevin"
    var country string 
	var age int = 23

	country = "Indonesia"

	var data1, data2, data3 string = "satu", "dua", "tiga"
	var nilai1, nilai2 int
	nilai1, nilai2 = 2, 3

	fmt.Println("Nama saya adalah : ", name)
	fmt.Println("Umur saya adalah : ", age)
	fmt.Println("Negara saya adalah : ", country)

	fmt.Println("data 1 : ", data1)
	fmt.Println("data 2 : ", data2)
	fmt.Println("data 3 : ", data3)

	fmt.Println("Nilai 1 : ", nilai1)
	fmt.Println("Nilai 2 : ", nilai2)

	hobi := "Badminton"

	// Multiple Declaration Variable

	var numerik1, string1, string2 = 2, "Data String", 3
	fmt.Println(numerik1,  string1, string2)



	fmt.Printf("%T, %T. %T", name, age, country, hobi)
}