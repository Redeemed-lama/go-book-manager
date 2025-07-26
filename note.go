package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open or create file for appending
	file, err := os.OpenFile("note.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter note title (or 'exit' to quit): ")
		if !scanner.Scan() {
			break
		}
		title := scanner.Text()
		if title == "exit" {
			fmt.Println("Exiting...")
			break
		}

		fmt.Print("Enter note content: ")
		if !scanner.Scan() {
			break
		}
		content := scanner.Text()

		note := fmt.Sprintf("Title: %s | Content: %s\n", title, content)
		_, err := file.WriteString(note)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			continue
		}

		fmt.Println("Note saved successfully! ✅")
	}
}
