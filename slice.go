package main

import "fmt"

func main() {

	name := [...]string{"Muh", "Fajrin", "Suhar", "Rudy", "Peter", "Qaffal"}

	slice1 := name[4:6]
	fmt.Println(slice1)


	days := [...]string {"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}


	slice2 := days[5:]
	hitungpanjang := len(slice2)


	fmt.Println(slice2)
	fmt.Println(hitungpanjang)


	slice2[0] = "Sabtu Baru" 
	slice2[1] = "Minggu Baru"

	fmt.Println(slice2)
	fmt.Println(days)
	

	// days[] 

}