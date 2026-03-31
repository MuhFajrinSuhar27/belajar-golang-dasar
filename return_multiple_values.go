package main 

import "fmt"

func getFullName() (string, string) {

	return "Fajrin", "Suhar"
}


func main() {

	firstName, lastName := getFullName() 

	fmt.Println(firstName, lastName)

}