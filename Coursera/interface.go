package main

import (

	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal interface{

	Eat()
	Move()
	Speak()
}

type Cow struct{

	name string
	food string
	locomotion string
	noise string
}

type Bird struct{

	name string
	food string
	locomotion string
	noise string
}

type Snake struct{

	name string
	food string
	locomotion string
	noise string
}


func(obj Cow) Eat(){

	fmt.Printf("\nFood for Cow: %s\n", obj.food)
}

func(obj Cow) Move(){

	fmt.Printf("\nLocomotion for Cow: %s\n", obj.locomotion)
}

func(obj Cow) Speak(){

	fmt.Printf("\nNoise for Cow: %s\n", obj.noise)
}

func(obj Bird) Eat(){

	fmt.Printf("\nFood for bird: %s\n", obj.food)
}

func(obj Bird) Move(){

	fmt.Printf("\nLocomotion for bird: %s\n", obj.locomotion)
}

func(obj Bird) Speak(){

	fmt.Printf("\nNoise for bird: %s\n", obj.noise)
}

func(obj Snake) Eat(){

	fmt.Printf("\nFood for snake: %s\n", obj.food)
}

func(obj Snake) Move(){

	fmt.Printf("\nLocomotion for snake: %s\n", obj.locomotion)
}

func(obj Snake) Speak(){

	fmt.Printf("\nNoise for snake: %s\n", obj.noise)
}

func main(){
	
	var obj_map map[string] Animal = make(map[string] Animal)
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Printf("\nEnter command as :\n(newanimal animal_name animal_type)\nOR\n(query animal_name animal_function),\nPress 'x' to exit : \n>")
	
	query,_ := reader.ReadString('\n')
	
	query = strings.ToLower(strings.TrimSpace(query))
	
	for ;; {
	
		if(query == "x"){
			break
		}
		
		query_list := strings.Split(query, " ")
		
		if(query_list[0] == "newanimal"){
		
			var animal_type string = query_list[2]
			
			switch animal_type{
			
				case "cow":
				
					new_obj := Cow{name : query_list[1], food : "grass", locomotion : "walk", noise : "moo"}
					fmt.Printf("\nCreated it!\nCow Name : %s\n", new_obj.name)
					obj_map[new_obj.name] = new_obj
					
				case "bird":
			
					new_obj := Bird{name : query_list[1], food : "worms", locomotion : "fly", noise : "peep"}
					fmt.Printf("\nCreated it!\nBird Name : %s\n", new_obj.name)			
					obj_map[new_obj.name] = new_obj
			
				case "snake":
				
					new_obj := Snake{name : query_list[1], food : "mice", locomotion : "slither", noise : "hss"}
					fmt.Printf("\nCreated it!\nSnake Name : %s\n", new_obj.name)
					obj_map[new_obj.name] = new_obj
			
				default:
				
					fmt.Printf("\n:: Invalid Animal Type ::\n")
			
			}
		
		}
		
		if(query_list[0] == "query"){
		
			generic_obj, check := obj_map[query_list[1]]
			
			if(!check){
			
				fmt.Printf("\n:: Invalid Animal Name ::\n")
			}else{
			
					var func_name string = query_list[2] 
			
					switch func_name{
				
					case "eat":
					
						generic_obj.Eat()
						
					case "move":
					
						generic_obj.Move()
						
					case "speak":
					
						generic_obj.Speak()
						
					default:
					
						fmt.Printf("\n:: Invalid Function ::\n")
					}		
			
				}
		
		}
		
		fmt.Printf("\nEnter command as\n(newanimal animal_name animal_type)\nOR\n(),\nPress 'x' to exit : \n>")
		query,_ = reader.ReadString('\n')
		query = strings.ToLower(strings.TrimSpace(query))
	}
	
}




