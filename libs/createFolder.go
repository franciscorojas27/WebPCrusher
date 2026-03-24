package libs

import (
	"fmt"
	"os"
	"strings"
)

func CreateWebpDir(path string) {
	if err := os.MkdirAll(path, 0755); err != nil {
		fmt.Printf("\033[31m"+"Error creating output directory: %s\033[0m\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("\033[32m"+"Output directory created successfully: %s\033[0m\n", path)
		large := len(path) + 50
		fmt.Println("\033[33m" + "|" + strings.Repeat("-", large-2) + "|" + "\033[0m")
	}
}
