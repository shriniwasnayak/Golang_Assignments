package main

import (

	"fmt"
	"math"
	"strconv"
)

type Vehicle interface{

	requiredUnits(int) int
	returnInfo() string
}

type Car struct{

	name string
	model string
}

type Bike struct{

	name string
	ccPower int
}

func(c *Car) requiredUnits(persons int) int{

	return int(math.Ceil(float64(persons)/4.0)) 

}

func(c *Car) returnInfo() string{

	return "Model of Car : " + c.model
}

func(b *Bike) requiredUnits(persons int) int{

	return int(math.Ceil(float64(persons)/2.0))

}

func(b *Bike) returnInfo() string{

	return "Power of Bike : " + strconv.Itoa(b.ccPower)
}

func main(){

	var c1 Car = Car{"Ford", "ABC123"}
	
	var b1 Bike = Bike{"Splender", 180}

	var generic Vehicle
	
	cp := &c1
	bp := &b1
	
	generic = cp
	
	fmt.Println(generic.requiredUnits(23))
	fmt.Println(generic.returnInfo())
	
	generic = bp
	
	fmt.Println(generic.requiredUnits(23))
	fmt.Println(generic.returnInfo())
}
