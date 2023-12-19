package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"todo/structures"
)

func main() {
	fmt.Println(("Wellcome to Todo Application"))
	commend := flag.String("commend-task", "exit", "commend for create , edit and ...")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)

	switch *commend {
	case "create-task":
		var newTask = structures.Task{}
		fmt.Println("enter a title for Task:")
		scanner.Scan()
		var title = scanner.Text()

		fmt.Println("enter a title for category's Task:")
		scanner.Scan()
		var titleCategory = scanner.Text()

		fmt.Println("enter a date for Task:")
		scanner.Scan()
		var date = scanner.Text()

		newCategory := createCategory(titleCategory)
		newTask.CreateTask(title, date, newCategory, false)
		print("task created")

	case "create-category":
		fmt.Println("enter a title for category:")
		scanner.Scan()
		var titleCategory = scanner.Text()
		createCategory(titleCategory)
		println("category created")
	}

}
func createCategory(title string) structures.Category {
	newCategory := structures.Category{}
	newCategory.CreateCategory(title)
	return newCategory
}
