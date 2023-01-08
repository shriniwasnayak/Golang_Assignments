package main

import (

	"fmt"
	"os"
	"bufio"
	"strings"

)

var reader = bufio.NewReader(os.Stdin)
	
type Person struct{

	name string
	age int

}

func ReadData(arr []Person, id int){

	fmt.Printf("\nEnter your name : ")
	
	data,_ := reader.ReadString('\n')
	
	data = strings.TrimSpace(data)

	arr[id].name = data
	
	fmt.Printf("\nEnter you age : ")
	
	fmt.Scanf("%d",&arr[id].age)

	fmt.Printf("\nName : %s\tAge : %d\n",arr[id].name, arr[id].age)

}

func(p Person) printData(){

	fmt.Printf("\nName : %s\tAge : %d\n",p.name, p.age)
}

func main(){

	var arr [3]Person
	
	for i:=0 ; i<3; i++{
	
		ReadData(arr[:], i)
	
	}

	fmt.Printf("\n*******************\n")

	for i:=0 ; i<3; i++{
	
		arr[i].printData()
	
	}

}
