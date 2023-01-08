package main

import (

	"fmt"
	"sync"
)

var core int

var mut sync.Mutex

func incCore(num int, wg *sync.WaitGroup){

	mut.Lock()

	core += num

	mut.Unlock()

	(*wg).Done()

}

func main(){


	core = 0
		
	var wg sync.WaitGroup

	wg.Add(3)

	go incCore(4, &wg)
	go incCore(2, &wg)
	go incCore(5, &wg)

	wg.Wait()
		
	fmt.Print("\n",core, "\n\n")
}
