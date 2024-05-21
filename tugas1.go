package main

import "fmt"

func main() {
    // Deklarasi variabel
	str, age, isNight := "Hello, World!", 23, true
    // str := "Hello, Golang!"
    // age := 23
    // isNight := true

    // Print nilai variabel
    fmt.Printf("String: %s\n", str)
    fmt.Printf("Integer: %d\n", age)
    fmt.Printf("Boolean: %t\n", isNight)

    // Print tipe data
    fmt.Printf("Tipe data dari variabel str: %T\n", str)
    fmt.Printf("Tipe data dari variabel age: %T\n", age)
    fmt.Printf("Tipe data dari variabel isNight: %T\n", isNight)
}
