package main

import "fmt"

type Person struct{

	name string
}

func(p Person) alterName(){

	fmt.Println("In normal alter name")
	
	p.name = "Alter normal"
}

/*
func(p *Person) alterNamePointer(){

	fmt.Println("In pointer alter name")
	
	p.name = "Alter Pointer"
}*/


func(p *Person) alterName(){

	fmt.Println("In pointer alter name")
	
	p.name = "Alter Pointer"

}

func main(){

	var p1 Person = Person{"temp name"}

	fmt.Println(p1)
	
	p1.alterName()
	
	fmt.Println(p1)
	
	fmt.Println(p1)
}
