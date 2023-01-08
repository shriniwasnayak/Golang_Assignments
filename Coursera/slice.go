package main

import (

	"fmt"
	"bufio"
	"os"
	"strings"
	"sort"
	"strconv"

)

func main(){

	// creating empty slice, with capacity  = 3 and length = 0
	var data_slice = make([]int, 0, 3)
	
	//var iter int = 0
	
	var reader = bufio.NewReader(os.Stdin)
	
	for ;; {
	
		fmt.Printf("\nEnter integer or 'X' to exit : ")
		
		data,_ := reader.ReadString('\n')		
	
		data = strings.TrimSpace(data)
		
		data = strings.ToUpper(data)
		
		if(data == "X"){
			break
		}
		
		int_data, _ := strconv.Atoi(data)
		
		data_slice = append(data_slice, int_data)
		
		sort.Slice(data_slice, func(p, q int) bool {
			return data_slice[p] < data_slice[q]})
		
		fmt.Printf("\nSorted Slice : %v\n", data_slice)
	
	}

}

