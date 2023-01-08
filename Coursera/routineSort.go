/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

*/


package main

import(

	"fmt"
	"strings"
	"strconv"
	"os"
	"bufio"
	"sync"
	"sort"
)

var global_chan chan int

func readArray() []int {

	var data string
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Printf("\nEnter list of space separated intergers (for eg : '1 2 3 4', size at least 4) : \n")
	
	data,error_obj := reader.ReadString('\n')
	
	if(error_obj != nil){
		
		fmt.Printf("\n :: INVAILD INPUT ::\n")
		os.Exit(0)	
		
	}
	
	data = strings.TrimSpace(data)
	
	var arr []int = make([]int, 0, 0)
	
	for _,temp := range strings.Split(data, " "){
	
		val,_ := strconv.Atoi(temp)
		
		arr = append(arr, val)	
	}
	
	return arr

}

func sortSlice(arr []int, routineId int, wg *sync.WaitGroup){

	fmt.Printf("\nGo-routine number : %d\tArray to sort : %v\n",routineId, arr)
	
	sort.Ints(arr)

	fmt.Printf("\nGo-routine number : %d\tArray sorted : %v\n",routineId, arr)

	wg.Done()
}

func merge(a []int, b []int, c chan []int) []int{

	var s1,s2 int = len(a), len(b)
	
	var i,j int
	
	var arr []int
	
	for i < s1 && j <s2{
	
		if(a[i] < b[j]){
			
			arr = append(arr, a[i])
			i+=1
			
		}else{
		
			arr = append(arr, b[j])
			j+=1
		}
	
	}
	
	for i<s1 {
	
		arr = append(arr, a[i])
	
		i+=1
	}
	
	for j<s2 {
	
		arr = append(arr, b[j])
		
		j+=1
	}

	if(c == nil){
		return arr
	}
		

	c <- arr

	return nil
}

func main(){

	var arr []int = readArray()

	var size int = len(arr)

	if(size < 4){
		
		fmt.Printf("\n :: Size of array less than 4 ::\n")
		
		os.Exit(1)
	}

	var list_of_slice [4][]int

	list_of_slice[0] = arr[0 : size/4]
	list_of_slice[1] = arr[size/4 : size/2]
	list_of_slice[2] = arr[size/2 : 3*size/4]
	list_of_slice[3] = arr[3*size/4 : size]

	//creating a wait group for functionality : all go routines should complete individual sort before merge can begin
	var wg sync.WaitGroup
	
	wg.Add(4)	 
	
	go sortSlice(list_of_slice[0], 1, &wg)
	go sortSlice(list_of_slice[1], 2, &wg)
	go sortSlice(list_of_slice[2], 3, &wg)
	go sortSlice(list_of_slice[3], 4, &wg)
	
	wg.Wait()
	
	//creating channel for : 
	var c chan []int = make(chan []int, 1)
	
	go merge(list_of_slice[0], list_of_slice[1], c)
	go merge(list_of_slice[2], list_of_slice[3], c)
	
	var m1 []int = <- c
	var m2 []int = <- c
	 
	var final_arr []int = merge(m1, m2, nil)
	
	fmt.Printf("\n=======================\n")
	fmt.Printf("Final Sorted array : %v", final_arr)
	fmt.Printf("\n=======================\n\n")

}
