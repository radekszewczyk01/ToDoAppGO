package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"toDoApp/models"
	"toDoApp/utils"
)

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	var tasks []models.Task
	err = json.Unmarshal(bytes, &tasks)
	if err != nil {
		fmt.Println("Error parsing json:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter commands (type 'help' for a list of commands):")
	for {
		fmt.Print("> ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		args := strings.Split(cmd, " ")
		switch args[0] {

		case "help":
			utils.OnHelp()
			continue

		case "show":
			utils.OnShow(tasks)
			continue
		}
	}
	utils.OnExit(tasks)
}
