package main

import (
    "fmt"
    "os"
    "strings"
    //"log"
)

func createFile(fname string) bool {
	myfile, e := os.Create(fname)
	if e != nil {
		//log.Fatal(e)
		return false
	}
	//log.Println(myfile)
	myfile.Close()
	return true
}

func main() {
	if len(os.Args) > 1 {
		onfile := string(os.Args[1])
		extension := strings.Split(onfile,".")
		parts := extension[len(extension)-1]
		new_file := ""
		if parts != "" {
			if parts != "txt" {
				if (len(parts) <= 3) && (len(parts) >= 1) {
					isfile := createFile(onfile)
					if isfile { fmt.Println("Text file created:",onfile) } else {
							fmt.Println("Error! The file has not been created.")
					}
				} else {
					new_file = onfile + ".txt"
					isfile := createFile(new_file)
					if isfile { fmt.Println("Text file created:",new_file) } else {
							fmt.Println("Error! The file has not been created.")
					}
				}
			} else {
				isfile := createFile(onfile)
				if isfile { fmt.Println("Text file created:",onfile) } else {
					fmt.Println("Error! The file has not been created.")
				}
			}
		}
	}
}
