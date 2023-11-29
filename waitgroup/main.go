package main

import (
	"fmt"
	"os"
	"sync"
	"test_buffer/helper"
)

var folder = "../test_files"

func main() {
	defer helper.Timer("wg")()

	workers, _ := os.ReadDir(folder)
	length := len(workers)
	files := make([][]byte, 0, length)
	ch := make(chan int, length/2)
	var wg sync.WaitGroup
	wg.Add(length)

	for i := 0; i < length; i++ {
		ch <- 1

		go func(i int) {
			v, err := os.ReadFile(folder + "/" + workers[i].Name())
			if err != nil {
				fmt.Println(err)
			}
			files = append(files, v)

			<-ch
			wg.Done()
		}(i)
	}
	wg.Wait()
}
