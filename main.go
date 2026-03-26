package main

import (
	"WebP-Crusher/libs"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {

	filePathClean := flag.String("p", "", `Path to input images directory`)
	numWorkers := flag.Int("w", 20, `Number of worker goroutines`)
	flag.Parse()

	jobs := make(chan Job, 100)
	var wg sync.WaitGroup

	for i := 0; i < *numWorkers; i++ {
		go worker(i, jobs, &wg)
	}

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
		// skip the webp output directory itself
		if info.IsDir() && filepath.Clean(path) == filepath.Clean(webpPath) {
			return filepath.SkipDir
		}

		// only handle files
		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))

		if ext == ".webp" {
			destFile := filepath.Join(webpPath, filepath.Base(path))
			srcClean := filepath.Clean(path)
			destClean := filepath.Clean(destFile)
			if srcClean == destClean {
				return nil
			}
			if _, statErr := os.Stat(destFile); os.IsNotExist(statErr) {
				srcF, oerr := os.Open(path)
				if oerr != nil {
					fmt.Printf("\033[31mfailed to open source webp %s: %v\033[0m\n", path, oerr)
					return nil
				}
				defer srcF.Close()
				dstF, cerr := os.Create(destFile)
				if cerr != nil {
					fmt.Printf("\033[31mfailed to create dest webp %s: %v\033[0m\n", destFile, cerr)
					return nil
				}
				defer dstF.Close()
				if _, copyErr := io.Copy(dstF, srcF); copyErr != nil {
					fmt.Printf("\033[31mfailed copying webp %s -> %s: %v\033[0m\n", path, destFile, copyErr)
				}
			}
			return nil
		}

		if ext == ".jpg" || ext == ".jpeg" || ext == ".png" || ext == ".gif" || ext == ".jfif" {
			base := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))
			destFile := filepath.Join(webpPath, base+".webp")
			if _, statErr := os.Stat(destFile); os.IsNotExist(statErr) {
				wg.Add(1)
				jobs <- Job{Path: path, Ext: ext, DestPath: webpPath}
			}
		}
		return nil
	}); err != nil {
		fmt.Printf("\033[31m"+"Error walking the input path: %v\033[0m\n", err)
		os.Exit(1)
	}
	close(jobs)
	wg.Wait()
}
