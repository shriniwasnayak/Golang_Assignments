package main

import (

	"fmt"
	"bufio"
	"os"
	"strings"
	"io/ioutil"

)

type Person struct{

	fname [20]byte
	lname [20]byte

}

func main(){

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\nEnter file name : ")
	
	data,_,_ := reader.ReadLine()
	
	var file_name string = strings.TrimSpace(string(data[:]))

	file_data,_ := ioutil.ReadFile(file_name)	

	struct_arr := make([]Person, 0, 0)
	
	for _, val := range strings.Split(string(file_data[:]), "\n"){
	
		if(strings.Contains(val, " ")){
		
			temp_arr := strings.Split(val, " ")
	
			var temp_struct Person
		
			copy(temp_struct.fname[:], temp_arr[0])
			copy(temp_struct.lname[:], temp_arr[1])
			 
			struct_arr = append(struct_arr, temp_struct)
		}	
	}

	fmt.Print("\n\nNames in files are : \n")
	
	for _,obj := range struct_arr{
	
		fmt.Printf("\n%s %s\n", string(obj.fname[:]), string(obj.lname[:]))
	}

}
