package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	opts := false
	sourcefolder := "/sourcefiles"
	destinationfolder := "/destinationfiles"
	mapfolder := "/mapfiles"
	schemafolder := "/schemafiles"
	fmt.Printf("C1D0U484 Inline Parser Utility - (c) Copyright Com1Software 1992-2022\n")
	fmt.Printf("Repository : github.com/Com1Software/Inline-Parser-Utility\n")
	fmt.Printf("Operating System : %s\n", runtime.GOOS)
	exefile := "/C1D0U484/C1D0U484.EXE"

	if CheckForFile(exefile) {
		fmt.Printf("- C1D0U484 Inline Parser Detected\n")
	}

	switch {
	//----------------------------
	case len(os.Args) == 2:
		cmd := os.Args[1]
		switch {
		case cmd == "overview":
			fmt.Printf("Overview\n")
			_, err := os.Stat(sourcefolder)
			x := os.IsNotExist(err)
			if x {
				fmt.Printf("Source Folder %s Not Found.\n", sourcefolder)
			} else {
				fmt.Printf("Source Folder %s Found.\n", sourcefolder)
			}
			_, err = os.Stat(destinationfolder)
			x = os.IsNotExist(err)
			if x {
				fmt.Printf("Destination Folder %s Not Found.\n", destinationfolder)
			} else {
				fmt.Printf("Destination Folder %s Found.\n", destinationfolder)
			}
			_, err = os.Stat(mapfolder)
			x = os.IsNotExist(err)
			if x {
				fmt.Printf("Map Folder %s Not Found.\n", mapfolder)
			} else {
				fmt.Printf("Map Folder %s Found.\n", mapfolder)
			}
			_, err = os.Stat(schemafolder)
			x = os.IsNotExist(err)
			if x {
				fmt.Printf("Schema Folder %s Not Found.\n", schemafolder)
			} else {
				fmt.Printf("Schema Folder %s Found.\n", schemafolder)
			}

		default:
			opts = true
		}
	case len(os.Args) == 1:
		opts = true
	}
	if opts {
		fmt.Println("Command Line Options")
		fmt.Println("")
		fmt.Println("Syntax: " + filepath.Base(os.Args[0]) + " <command> ")
		fmt.Println("")
		fmt.Println("- Commands:")
		fmt.Println("      overview  - Displays Inline Parser environment information .")
		fmt.Println("")

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
