package helper

import (
	"fmt"
	"time"
)

var Folder = "./test_files"

func Timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}
