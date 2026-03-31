package main

import "fmt"

func main() {

	var NilaiAkhir = 90
	var AbsenAkhir = 80

	var luluNilaiAkhir bool =  NilaiAkhir > 80
	var luluAbsenAkhir bool = AbsenAkhir > 80

	var lulusMataKuliah = luluNilaiAkhir &&  luluAbsenAkhir 
	
	fmt.Println(lulusMataKuliah)

	
}