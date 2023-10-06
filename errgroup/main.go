package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"test_buffer/helper"

	"golang.org/x/sync/errgroup"
)

var folder = "./test_files"

func Task(task int) error {
	if rand.Intn(10) == task {
		return fmt.Errorf("Task %v failed", task)
	}
	fmt.Printf("Task %v completed", task)
	return nil
}

func main() {
	defer helper.Timer("errgroup")()

	workers, _ := os.ReadDir(folder)
	length := len(workers)
	files := make([][]byte, 0, length)
	eg := &errgroup.Group{}

	for i := 0; i < length; i++ {
		eg.Go(func() error {
			return func(i int) error {
				v, err := os.ReadFile(folder + "/" + workers[i].Name())
				if err != nil {
					return err
				}

				files = append(files, v)
				return nil
			}(i)
		})
	}

	if err := eg.Wait(); err != nil {
		log.Fatal("Error", err)
	}

	fmt.Println(files)
}
