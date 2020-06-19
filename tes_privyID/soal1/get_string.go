//mendeklarasikan package
package main

//import library
import (
	"fmt"
	"strings"
)

//fungsi get string dengan 1 parameter dan 2 return int dan string
func getString(x string) (int, string) {
	//mendeklarasikan variabel
	word1 := "19374048"
	word2 := "aiueobcd"
	var indexOf int
	var change string

	fmt.Println("Kata pertama adalah ", word1)
	fmt.Println("Kata kedua adalah ", word2)

	//apabila parameter yang dicari(x) merupakan bagian dari variabel word1
	//maka menghasilkan nilai true
	isExists := strings.Contains(word1, x)
	if isExists { //jika benar
		indexOf = strings.Index(word1, x) //mencari posisi dari index x pada variabel word1
		change = word2[indexOf:] //mengisi change dengan karakter mulai dari karakter ke-x sampai karakter terakhir	
	} else { //jika salah
		fmt.Println("Index tidak ditemukan!")
	}

	//mengembalikan nilai dari indexOf dan change
	return indexOf, change
}

func main() {
	var x string
	fmt.Println("Angka apa yang anda cari? ")
	fmt.Scan(&x)
	indexInfo, changeInfo := getString(x) //mendeklarasikan 2 nilai yang dikembalikan fungsi getString()
	fmt.Println(x, " merupakan index ke", indexInfo, "dari kata pertama")
	fmt.Println("Huruf ke", indexInfo, "sampai terakhir pada kata kedua adalah ", changeInfo)
}