package main

import (

	"fmt"
	"sync"
	"time"
)

var philCount int = 5
var eatCount int = 3

var mut sync.Mutex

var printInfo bool = false

var time_to_sleep int = 1

type Phil struct{

	id int
	state string
	left *Chopstick
	right *Chopstick
}

type Chopstick struct{

	id int
	inUse bool 
}

func host(phil *Phil) bool{

	var response bool = false

	mut.Lock()
	
	if( (!phil.left.inUse) && (!phil.right.inUse) ){
	
		phil.left.inUse = true
		phil.right.inUse = true
		response = true
	}
	
	mut.Unlock()

	return response
}

func(p Phil) PrintPhil(eatId int){

	fmt.Printf("\nPhil id : %d\tState : %s\tEat Id : %d\tLeft Chopstick Id : %d\tLeft Chopstick in Use : %v\tRight Chopstick Id : %d\tRight Chopstick in Use : %v\n", p.id, p.state, eatId, p.left.id, p.left.inUse, p.right.id, p.right.inUse)
	
}


func(phil *Phil) eat(eatId int, wg *sync.WaitGroup){

	if(printInfo){
		
		(*phil).PrintPhil(eatId)
	}
	
	
	for host(phil) != true{
	
		//wait till host gives permission
	}
	
	fmt.Printf("\nstart to eat %d, eat id : %d\n",phil.id, eatId)
	
	phil.state = "eat"
	
	if(printInfo){
		
		(*phil).PrintPhil(eatId)
	}
	
	time.Sleep(time.Duration(time_to_sleep) * time.Second)
	
	fmt.Printf("\nfinish eating %d, eat id : %d\n",phil.id, eatId)	
	
	mut.Lock()
	
	phil.left.inUse = false
	phil.right.inUse = false
	
	mut.Unlock()
	
	if(printInfo){
		
		(*phil).PrintPhil(eatId)
	}
	
	wg.Done()
}

func main(){

	var wg sync.WaitGroup

	var philArr []*Phil = make([]*Phil,0,0)

	var chopstickArr []*Chopstick = make([]*Chopstick,0,0)
	
	for i := 0; i < philCount; i++{
	
		chopstickArr = append(chopstickArr, &Chopstick{id : i, inUse : false})
	}
	
	for i := 0; i < philCount; i++{
	
		philArr = append(philArr, &Phil{id : i+1, state : "think", left : chopstickArr[i], right : chopstickArr[(i+1)%philCount]})
	}
	
	for i:= 0; i < eatCount; i++{
	
		wg.Add(philCount)
	
		for j:=0; j < philCount; j++{
			
			go (*philArr[j]).eat(i, &wg)
		}
	}

	wg.Wait()
}
