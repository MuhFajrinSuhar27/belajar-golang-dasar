package main

import "fmt"

func main() {

	name := "Fajrin"

	if name == "Fajri1n" {
		fmt.Print("Berhasil")
	} else if name == "Denis" {
		fmt.Print("Bukan Nama yang sebenarnyaa")
	} else if name == "Joko" {
		fmt.Print("Bukan nama sebeneranya juga")
	}else {
		fmt.Println("Tidak ada yang benar")
	}



	if length := len(name); length > 5 {
		fmt.Println("minimal panjangnyaa sudah benar")
	} else {
		fmt.Print("Panjang Tidak Mencukupi")
	}

}