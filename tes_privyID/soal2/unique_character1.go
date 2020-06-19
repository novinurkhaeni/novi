//mendeklarasikan package
package main

//import library
import (
	"fmt"
	"strings"
)

func uniqueCharacter(x string) (string) {
	//deklarasi daftar karakter unik
	unique := "[]{}"
	var res string
	//apabila parameter yang dicari(x) merupakan bagian dari variabel unique
	//maka menghasilkan nilai true
	isExists := strings.Contains(unique, x)
	if isExists { //jika benar
		if x == "[" || x == "]"{ //jika karakter berupa "[" atau "]"
			res = "False, [ berpasangan dengan karakter ]"
		}else if x == "{" || x == "}" { //jika karakter berupa "{" atau "}"
			res = "False, { berpasangan dengan karakter }"
		}else if x == "[]" || x =="{}"{ //jika karakter berupa "[]" atau "{}"
			res = "True"
		}
	} else { //jika salah
		res = "nil"
	}

return res
}

func main() {
	var character string
	fmt.Println("Masukan karakter unik [, ], { & }")
	fmt.Scan(&character)
	fmt.Println(uniqueCharacter(character))
}