package main
import "fmt"


func sayHello(name string) string { 
	hello := "Hello " + name
	return hello

}

func main() {

	result := sayHello("Fajrin")
	fmt.Println(result)

}
