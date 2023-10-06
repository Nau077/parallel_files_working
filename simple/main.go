package main

import (
	"fmt"
	"os"
	"test_buffer/helper"
)

var folder = helper.Folder

func main() {
	defer helper.Timer("simple")()

	workers, _ := os.ReadDir(folder)
	length := len(workers)
	files := make([][]byte, 0, length)

	for i := 0; i < length; i++ {
		v, err := os.ReadFile(folder + "/" + workers[i].Name())
		if err != nil {
			fmt.Println(err)
		}
		files = append(files, v)
	}
}
