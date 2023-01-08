package main

import (

	"fmt"
	"bufio"
	"os"
	"strings"
	"encoding/json"

)

type Person struct{

	Name string
	Address string

}

func(p Person) printIt(){

	fmt.Printf("\nName : %s\nAddress : %s\n", p.Name, p.Address)

}

func main(){

	var person_obj Person
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Printf("\nEnter Name : ")
	
	data,_ := reader.ReadString('\n')

	person_obj.Name = strings.TrimSpace(data)

	fmt.Printf("\nEnter Address : ")
	
	data,_ = reader.ReadString('\n')

	person_obj.Address = strings.TrimSpace(data)

	person_obj.printIt()
	
	json_obj,_ := json.Marshal(person_obj)
	
	fmt.Print("Json object : ", json_obj, "\n")

	fmt.Print("Json object in string format : ", string(json_obj), "\n")

}
