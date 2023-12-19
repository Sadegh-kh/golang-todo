package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"todo/structures"
)

var scanner = bufio.NewScanner(os.Stdin)
var userStorage []structures.User

func main() {
	fmt.Println(("Wellcome to Todo Application"))
	command := flag.String("commend-task", "exit", "commend for create , edit and ...")
	flag.Parse()
	for {
		runCommand(*command)
		println("please enter another commend or exit")
		scanner.Scan()
		*command = scanner.Text()

	}

}

func runCommand(command string) {
	switch command {
	case "create-task":
		createTask()
	case "create-category":
		createCategory()
		println("category created")

	case "register-user":
		register()
	case "login":
		login()
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid :", command)

	}
}

func login() {
	fmt.Println("Enter email :")
	scanner.Scan()
	email := scanner.Text()
	fmt.Println("Enter password :")
	scanner.Scan()
	password := scanner.Text()

	if userExist(email) {
		if currectPass(email, password) {
			fmt.Println("login successfuly")
		} else {
			fmt.Println("your password or email is wrong!")
		}

	}
}

func currectPass(email, password string) bool {
	for _, value := range userStorage {
		if value.Email == email && value.Password == password {
			return true
		}
	}
	return false
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
	newUser.CreateUser(name, email, password)
	if !userExist(newUser.Email) {
		userStorage = append(userStorage, newUser)
	} else {
		fmt.Printf("user %s exist!\n", newUser)
	}
	fmt.Println(userStorage)

}

func userExist(email string) bool {
	for _, value := range userStorage {
		// email is the primary key
		if email == value.Email {
			return true
		}
	}
	return false
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

	newTask.CreateTask(title, date, newCategory, false)
	println("task created")
}
