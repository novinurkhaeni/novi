//deklarasi package
package main

//import library
import (
	"fmt"
)

//fungsi dengan parameter dan return slice dengan tipedata string
func fungsiA(slice []string) []string {
	//variabel fungsiMap dideklarasikan sebagai map
	//dengan tipe data key adalah string
	//dan value-nya sesuai dengan tipedata parameter
	//make() untuk menampung key dan value dari parameter
	fungsiMap := make(map[string]struct{})

	//perulangan for.. range 
	//_ = key
	//v = value
	//mengisi v dengan value dari parameter
	for _, v := range slice{ 
		fungsiMap[v] = struct{}{} 
	} 

	//variabel fungsiSlice dideklarasikan sebagai slice
	//dengan lebar adalah 0 element

	//fungsiSlice := make([]string, 0, len(uniqMap)) 
	fungsiSlice := make([]string, 0, len(fungsiMap))
	//perulangan for.. range 
	//v = value
	//mengisi v dengan value dari fungsiMap
	for v := range fungsiMap { 
		//menambah element slice
		fungsiSlice = append(fungsiSlice, v) 
	} 

	//mengembalikan nilai fungsiSlice
	return fungsiSlice 
}


//fungsi utama/method main
func main(){
	str := []string{"novi", "oti", "budi", "lina"}
	fmt.Println(fungsiA(str))
}