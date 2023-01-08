
// EXPLANATION AT END OF CODE

package main

import (

	"fmt"
	"time"
)

//global variable (shared by all go routines)
var flag bool = true


// function sets flag and prints value accordingly
func printValue(num int){

	if(num % 2 == 0){
	
		fmt.Printf("\n:: Flag is altered to true :: for value %d\n", num)
		
		flag = true
	}else{
	
		fmt.Printf("\n:: Flag is altered to false :: for value %d\n", num)
	
		flag = false
	}
	
	// sleep to create race condition
	time.Sleep(1*time.Second)
	
	fmt.Printf("\nNumber : %d :: Flag : %v\n",num,flag)
}

func main(){

	fmt.Printf("\n ======== EXECUTION WITHOUT GO ROUTINES BEGIN =========== \n")

	printValue(5)
	
	printValue(4)
	
	printValue(7)
	
	printValue(12)

	fmt.Printf("\n ======== EXECUTION WITHOUT GO ROUTINES COMPLETED =========== \n")

	fmt.Printf("\n ======== EXECUTION WITH GO ROUTINES BEGIN =========== \n")


	go printValue(5)
	
	go printValue(4)
	
	go printValue(7)
	
	go printValue(12)
	
	time.Sleep(3*time.Second)
	
	fmt.Printf("\n ======== EXECUTION WITH GO ROUTINES COMPLETED =========== \n\n")

}


/*

The code has three main components :

	1) Global boolean variable 'flag'
	2) Function to print number and flag
	3) Go routines
	
The function, accepts an integer value, sets the flag to true if number is even, false otherwise.

Under normal execution ::

	The function sets the flag as per the number (eg : if num is 2, flag is true, if num is 3, flag is false) and prints the same.
	
Using go routines ::

	The function still alters the flags, however the value of the flag is altered by another go routine before the funtions prints it.
	
Race condition ::

	Sicne the variable flag is shared among all go routines, the value is being altered by a routne before it is printed by the routine in execution there by creating issue of DIRTY READ.
	


TLDR :

what is the race condition ? :: the varaible flag is altered without synchronization

why does it occur ? :: go routines alter the flag before execution of one routine is completed, i.e. due to interleaving of instructions. 

*/
