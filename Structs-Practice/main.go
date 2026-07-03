package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/structspractice/note"
	"example.com/structspractice/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	printSomething(1)
	printSomething(1.5)
	printSomething("Welcom to Note & TODO:")

	title, content := getNoteData()

	todoText := getUsrInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)

	if err != nil {
		return
	}

	outputData(userNote)

}

func printSomething(value any) {
	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println(value)
	default:
		fmt.Println(value)
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return err
	}

	fmt.Println("Saving the note succeded!")

	return nil
}

func getNoteData() (string, string) {
	title := getUsrInput("Note title: ")

	content := getUsrInput("Note content: ")

	return title, content

}

func getUsrInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
