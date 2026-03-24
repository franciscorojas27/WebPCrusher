package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
	"sync"

	webp "github.com/chai2010/webp"
)

const (
	workerWidth = 6
	inputWidth  = 60
	outputWidth = 60
)

func PrintHeader() {
	format := fmt.Sprintf("| %%-%ds | %%-%ds | %%-%ds |\n", workerWidth, inputWidth, outputWidth)
	fmt.Printf(format, "Worker", "Input", "Output")
	innerLen := workerWidth + inputWidth + outputWidth + 8
	fmt.Printf("|%s|\n", strings.Repeat("-", innerLen))
}

func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max > 3 {
		return s[:max-3] + "..."
	}
	return s[:max]
}

var headerOnce sync.Once

func worker(id int, channel <-chan Job, wg *sync.WaitGroup) {
	for job := range channel {
		func(j Job) {
			defer wg.Done()

			file, err := os.Open(j.Path)
			if err != nil {
				fmt.Printf("\033[31mWorker %d: failed to open %s: %v\033[0m\n", id, j.Path, err)
				return
			}
			defer file.Close()

			img, _, err := image.Decode(file)
			if err != nil {
				fmt.Printf("\033[31mWorker %d: failed to decode %s: %v\033[0m\n", id, j.Path, err)
				return
			}

			baseName := strings.TrimSuffix(filepath.Base(j.Path), filepath.Ext(j.Path))
			outPath := filepath.Join(j.DestPath, baseName+".webp")

			out, err := os.Create(outPath)
			if err != nil {
				fmt.Printf("\033[31mWorker %d: failed to create output %s: %v\033[0m\n", id, outPath, err)
				return
			}
			defer out.Close()

			if err := webp.Encode(out, img, &webp.Options{Lossless: false, Quality: 75}); err != nil {
				fmt.Printf("\033[31mWorker %d: failed to encode %s -> %s: %v\033[0m\n", id, j.Path, outPath, err)
				return
			}

			inName := truncate(filepath.Base(j.Path), inputWidth)
			outName := truncate(filepath.Base(outPath), outputWidth)
			fmt.Printf("\033[32m| %-6d | %-60s | %-60s |\033[0m\n", id, inName, outName)
		}(job)
	}
}
