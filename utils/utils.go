package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"toDoApp/models"
)

func OnExit(tasks []models.Task) {
	jsonData, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error encoding json:", err)
		return
	}

	err = os.WriteFile("data.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("Data successfully written to json file")
}

func OnHelp() {
	fmt.Println("Available commands:")
	fmt.Println(" show - Shows tasks table")
	fmt.Println(" help - Shows tasks table")
}

func OnShow(tasks []models.Task) {
	if len(tasks) == 0 {
		fmt.Println("Brak zadań do wyświetlenia.")
		return
	}

	fmt.Println("Lista zadań:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n   Opis: %s\n   Status: %s\n\n", i+1, task.Name, task.Description, task.Status)
	}
}
