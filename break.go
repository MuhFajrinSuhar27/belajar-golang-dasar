package main 

import "fmt"

func main() {

	for i := 0; i < 9; i++ {
		if i == 5 {
			break
		}

		fmt.Println("Perulangan ke ", i)
		// fmt.Println("Selesai")
	}
}