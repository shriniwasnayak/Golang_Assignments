package main

import (

	"fmt"
	"encoding/json"
	"bufio"
	"strings"
	"os"
)

func main(){


	// creating map
	var data_map map[string] string
	
	data_map = make(map[string] string)
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("\nEnter Name : ")
	data,_ := reader.ReadString('\n')
	data_map["name"] = strings.TrimSpace(data)
	
	fmt.Print("\nEnter Address : ")
	data,_ = reader.ReadString('\n')
	data_map["address"] = strings.TrimSpace(data)
	
	json_obj,_ := json.Marshal(data_map)
	
	fmt.Println("\nJSON Object : ", json_obj,)
	
	fmt.Println("\nJSON Object String format : ", string(json_obj))

}
