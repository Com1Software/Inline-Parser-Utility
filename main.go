package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("C1D0U484 Inline Parser Utility - (c) Copyright Com1Software 1992-2022\n")
	fmt.Printf("Repository : github.com/Com1Software/Inline-Parser-Utility\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	exefile := "/C1D0U484/C1D0U484.EXE"

	if CheckForFile(exefile) {
		fmt.Printf("- Parser Detected")
	}

}

func CheckForFile(path string) bool {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	file.Close()
	return true
}
