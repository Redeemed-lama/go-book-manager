package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	readFile, err := os.Open("note.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer readFile.Close()
	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "done" {
			break
		}
		fmt.Println("all notes saved")
		fmt.Println("You wrote:", line)
	}

}
