package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"todo/structures"
)

var scanner = bufio.NewScanner(os.Stdin)
var authenticatedUser *structures.User

func main() {
	fmt.Println(("Wellcome to Todo Application"))
	command := flag.String("command-task", "exit", "command for create , edit and ...")
	flag.Parse()

	for {
		runCommand(*command)
		println("please enter another command or exit")
		scanner.Scan()
		*command = scanner.Text()

	}

}

func runCommand(command string) {
	if authedUser() {
		authRquiredCommands(command)
	} else {
		fmt.Println("you must login or register first!")
		switch command {
		case "register":
			register()
		case "login":
			login()
		case "exit":
			os.Exit(0)
		}
	}

}

func authRquiredCommands(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
	case "task-list":
		fmt.Println(structures.GetTaskList(authenticatedUser.ID))
	case "category-list":
		fmt.Println(structures.GetCategoryList())
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid :", command)
	}

}
func authedUser() bool {
	if authenticatedUser != nil {
		return true
	}
	return false
}

func login() {
	fmt.Println("Enter email :")
	scanner.Scan()
	email := scanner.Text()
	fmt.Println("Enter password :")
	scanner.Scan()
	password := scanner.Text()

	if structures.UserExist(email) {
		if structures.CheckPass(email, password) {
			fmt.Println("login successfuly")
			user := structures.GetUser(email)
			authenticatedUser = &user
		} else {
			fmt.Println("your password or email is wrong!")
		}

	} else {
		fmt.Println("email not exist!")
	}
}

func register() {
	var newUser = structures.User{}
	fmt.Println("Enter name :")
	scanner.Scan()
	name := scanner.Text()
	fmt.Println("Enter email :")
	scanner.Scan()
	email := scanner.Text()
	fmt.Println("Enter password :")
	scanner.Scan()
	password := scanner.Text()
	if !(structures.UserExist(email)) {
		newUser.CreateUser(name, email, password)
		newUser.AppendToStorage()
	} else {
		fmt.Printf("this email  %s exist!\n", email)
	}

}

func createCategory() structures.Category {
	fmt.Println("enter a title for category:")
	scanner.Scan()
	var titleCategory = scanner.Text()
	newCategory := structures.Category{}
	newCategory.CreateCategory(titleCategory)
	return newCategory
}

func createTask() {
	var newTask = structures.Task{}
	fmt.Println("enter a title for Task:")
	scanner.Scan()
	var title = scanner.Text()

	newCategory := createCategory()

	fmt.Println("enter a date for Task:")
	scanner.Scan()
	var date = scanner.Text()

	newTask.CreateTask(title, date, newCategory, false, authenticatedUser.ID)
	println("task created")
}
