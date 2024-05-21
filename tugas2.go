package main

import "fmt"

func main() {
    tinggi := 10
    
    for i := 1; i <= tinggi; i++ {
        // Membuat spasi
        for j := 1; j <= tinggi-i; j++ {
            fmt.Print(" ")
        }

        // Membuat bintang
        for k := 1; k <= 2*i-1; k++ {
            fmt.Print("*")
        }

        // Pindah ke baris baru
        fmt.Println()
    }
}
