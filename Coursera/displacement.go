package main

import (
	
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"

)

func GenDisplaceFn(acceleration float64, initial_velocity float64, initial_disp float64) func(float64) float64{

	fn := func(time float64) float64{
	
			var final_disp float64
	
			final_disp = initial_disp + (time*initial_velocity) + (0.5 * acceleration * time * time)
			
			return final_disp
	
			}
			
	return fn

}

func main(){

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("\nEnter Acceleration, initial velocity and initial displacement separated by space (eg : 2.5 0 2.3)\n")
	
	data,_ := reader.ReadString('\n')
	
	data = strings.TrimSpace(data)
	
	data_arr := strings.Split(data, " ")
	
	acceleration,_ := strconv.ParseFloat(data_arr[0], 64)
	initial_velocity,_ := strconv.ParseFloat(data_arr[1], 64)
	initial_disp,_ := strconv.ParseFloat(data_arr[2], 64)

	
	fn := GenDisplaceFn(acceleration , initial_velocity, initial_disp)
	
	var time float64
	
	fmt.Printf("\nEnter Time : ")
	
	fmt.Scanf("%f", &time)
	
	fmt.Printf("\nFianl Displacement : %f\n", fn(time))
	
}
