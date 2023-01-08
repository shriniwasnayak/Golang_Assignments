package main

import (

	"fmt"
	"bufio"
	"os"
	"strings"
)

func main(){

	var data string

	fmt.Printf("\nEnter string data : ")
	
	var reader = bufio.NewScanner(os.Stdin)

	//data,_ = reader.ReadString('')
	
	reader.Scan()
	
	data = reader.Text()
	
	// case generalization
	data = strings.ToLower(data)
	
	//if-else ladder      
	
	if(strings.HasPrefix(data, "i") && strings.Contains(data, "a") && strings.HasSuffix(data, "n") ) {
	
		fmt.Printf("\nFound!\n")
	} else{
	
		fmt.Printf("\nNot Found!\n")
	}

}

