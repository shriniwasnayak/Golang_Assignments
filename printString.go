package main

import (

	"fmt"
	"strings"
	"bufio"
	"os"

)

func main(){

	var data string
	
	fmt.Printf("\nEnter data : ")
	
	var reader = bufio.NewReader(os.Stdin)
	
	data,_ = reader.ReadString('\n')
	
	data = strings.TrimSpace(data)
	
	for x,v := range data{
	
		fmt.Printf("%d : %x\n", x, v)
	}

}

