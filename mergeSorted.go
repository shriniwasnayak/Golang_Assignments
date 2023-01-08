package main

import "fmt"

func merge(a []int, b []int) []int{

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

	return arr

}

func main(){


	fmt.Printf("%v\n", merge([]int{2,4,6}, []int{1,3,5,7,8}))
}
