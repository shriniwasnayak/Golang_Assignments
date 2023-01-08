package main

import(
	"fmt"
)

func main(){

	var input_data float32
	var output_data int
	
	fmt.Printf("\nEnter float number : ")
	
	fmt.Scanf("%f", &input_data)
	
	//convert data using type conversion
	
	output_data = int(input_data)
	
	fmt.Printf("\nConverted Integer : %d\n", output_data)

}
