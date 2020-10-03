package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	newFile, err := os.Create("a.txt")
	checkErr("Can't create", err)

	err = os.Truncate("a.txt", 0)
	checkErr("Can't truncate", err)

	newFile.Close()

	file, err := os.Open("a.txt")
	checkErr("Can't open", err)
	file.Close()

	file, err = os.OpenFile("a.txt", os.O_APPEND, 0644)
	checkErr("Can't open file", err)
	file.Close()

	var fileinfo os.FileInfo
	fileinfo, err = os.Stat("a.txt")

	p := fmt.Println
	p("File name:", fileinfo.Name())
	p("Size in bytes:", fileinfo.Size())
	p("Modification time:", fileinfo.ModTime())
	p("Is it directory?", fileinfo.IsDir())
	p("Permissions:", fileinfo.Mode())

	oldPath := "a.txt"
	newPath := "aaa.txt"

	err = os.Rename(oldPath, newPath)
	checkErr("Can't rename", err)

	fileinfo, err = os.Stat("aaa.txt")
	checkErr("The file doesn't exist! ", err)

	err = os.Remove("aaa.txt")
	checkErr("Can't remove", err)

}

func checkErr(msg string, err error) {
	if err != nil {
		log.Fatal(msg, err)
	}
}
