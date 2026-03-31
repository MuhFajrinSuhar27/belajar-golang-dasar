package main

import "fmt"

func getComplateName() (firstName, middleName, lastName string) {

	firstName = "Muh" 
	middleName = "Fajrin" 
	lastName = "Suhar" 

	return firstName, middleName, lastName
} 


func main() {

	a,b,c := getComplateName()

	fmt.Println(a,b,c)



}