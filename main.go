package main

import (
	"WebP-Crusher/libs"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	const numWorkers = 20
	jobs := make(chan Job, 100)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		go worker(i, jobs, &wg)
	}

	filePathClean := flag.String("p", "", `Path to input images directory`)
	flag.Parse()

	if *filePathClean == "" {
		fmt.Printf("\033[31m" + "Please provide a valid path to the input images directory using the -p flag.\033[0m\n")
		os.Exit(1)
	}

	if _, err := os.Stat(*filePathClean); os.IsNotExist(err) {
		fmt.Printf("\033[31m"+"File not exists: %s\033[0m\n", *filePathClean)
		os.Exit(1)
	}

	webpPath := filepath.Join(*filePathClean, "webp")
	if _, err := os.Stat(webpPath); os.IsNotExist(err) {
		libs.CreateWebpDir(webpPath)
	}

	if err := filepath.Walk(*filePathClean, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		ext := filepath.Ext(path)
		if ext == ".jpg" || ext == ".png" {
			wg.Add(1)
			jobs <- Job{Path: path, Ext: ext, DestPath: webpPath}
		}
		return nil
	}); err != nil {
		fmt.Printf("\033[31m"+"Error walking the input path: %v\033[0m\n", err)
		os.Exit(1)
	}
	close(jobs)
	wg.Wait()
}
