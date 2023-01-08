package main

import (

	"fmt"
	"bufio"
	"os"
	"strings"
)

type Animal struct{

	food string
	locomotion string
	noise string
}

func(obj Animal) printObj(){

	fmt.Printf("\n==========\nFood : %sLocomotion : %s\nNoise : %s\n===========\n", obj.food, obj.locomotion, obj.noise)
}

func(obj Animal) Eat(){

	fmt.Printf("\nFood : %s\n", obj.food)
}

func(obj Animal) Move(){

	fmt.Printf("\nLocomotion : %s\n", obj.locomotion)
}

func(obj Animal) Speak(){

	fmt.Printf("\nNoise : %s\n", obj.noise)
}

func main(){

	var cow Animal = Animal{food : "grass", locomotion : "walk", noise : "moo"}
	
	var bird Animal = Animal{food : "worms", locomotion : "fly", noise : "peep"}
	
	var snake Animal = Animal{food : "mice", locomotion : "slither", noise : "hss"}
	
	var obj_map map[string]Animal = make(map[string]Animal)
	obj_map["cow"] = cow
	obj_map["bird"] = bird
	obj_map["snake"] = snake
	
	/*var func_map map[string] func(Animal)() = make(map[string] func(Animal)())
	func_map["eat"] = Animal.Eat
	func_map["move"] = Animal.Move
	func_map["speak"] = Animal.Speak*/
	
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Printf("\nEnter query [animal info_tyoe] (eg : cow eat), Press 'x' to exit : \n>")
	
	query,_ := reader.ReadString('\n')
	
	query = strings.ToLower(strings.TrimSpace(query))
	
	for ;; {
	
		if(query == "x"){
			break
		}
		
		query_list := strings.Split(query, " ")
		
		var obj Animal
		var check bool
		
		obj, check = obj_map[query_list[0]]
		
		if(!check){
		
			fmt.Printf("\n :: Invalid Animal, Defaulting to cow:: \n")
			obj = obj_map["cow"]
			
		}
		
		var func_name string = query_list[1]
		
		switch func_name{
			
			case "eat":
			
				obj.Eat()
				
			case "move":
			
				obj.Move()
				
			case "speak":
			
				obj.Speak()
				
			default:
			
				fmt.Printf("\n:: Invalid Function ::\n")
		}
		
		fmt.Printf("\nEnter query [animal info_tyoe] (eg : cow eat), Press 'x' to exit : \n>")
		query,_ = reader.ReadString('\n')
		query = strings.ToLower(strings.TrimSpace(query))
	}
	
}




