package main

import (
	
	"fmt"

)

type Data struct{

	name string 
	age int
	is []int
}

func(obj Data) printIt(){

	fmt.Printf("\n==========\nName : %s\nAge : %d\nIS : %v\n=========\n", obj.name, obj.age, obj.is)
}

func(obj Data) alter(){

	obj.name = "alter1"
	obj.age = 100
	obj.is[1] = 50

}

func(obj *Data) pointerAlter(){

	(*obj).name = "alter1"
	obj.age = 100
	obj.is[1] = 50

}

func main(){

	var obj Data = Data{name : "Shri", age : 12, is : []int{1,2,3}}
	
	obj.printIt()

	obj.alter()
	fmt.Printf("\nAfter Alter\n")
	obj.printIt()
	
	obj.pointerAlter()
	fmt.Printf("\nAfter Pointer Alter\n")
	obj.printIt()
}
