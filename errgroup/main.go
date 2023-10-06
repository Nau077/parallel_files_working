package main

import (
	"log"
	"os"
	"test_buffer/helper"

	"golang.org/x/sync/errgroup"
)

var folder = helper.Folder

func main() {
	defer helper.Timer("errgroup")()

	workers, _ := os.ReadDir(folder)
	length := len(workers)
	files := make([][]byte, 0, length)
	eg := &errgroup.Group{}

	for i := 0; i < length; i++ {
		id := i
		eg.Go(func() error {
			return func() error {
				v, err := os.ReadFile(folder + "/" + workers[id].Name())
				if err != nil {
					return err
				}
				files = append(files, v)
				return nil
			}()
		})
	}

	if err := eg.Wait(); err != nil {
		log.Fatal("Error", err)
	}
}
