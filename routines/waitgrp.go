package main

import (
"fmt"
"sync"
)

func printMsg(data string, wg *sync.WaitGroup) {

	for i := 0; i < 5; i++ {
		fmt.Printf("%s", data)
	}

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go printMsg("my msg\n", &wg)

	wg.Add(1)

	go printMsg("my msg 2\n", &wg)

	fmt.Printf("Main Msg\n")

	wg.Wait()

//time.Sleep(time.Second)
}

