package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"todo/structures"
)

const userStoragePath = "users_storage.txt"

var scanner = bufio.NewScanner(os.Stdin)
var authenticatedUser *structures.User
var file *os.File
var err error

func main() {
	fmt.Println(("Wellcome to Todo Application"))
	command := flag.String("command-task", "exit", "command for create , edit and ...")
	flag.Parse()
	_, err = os.Stat(userStoragePath)
	// if file exist, err == nil
	if err != nil {
		file, err = os.Create(userStoragePath)
		if err != nil {
			fmt.Printf("error happed when we create file: %v\n", err)
		}
	}
	loadFile()

	for {
		runCommand(*command)
		println("please enter another command or exit")
		scanner.Scan()
		*command = scanner.Text()

	}

}

func loadFile() {
	var data = make([]byte, 512)
	file, err = os.Open(userStoragePath)
	defer file.Close()
	_, err = file.Read(data)
	data_string := string(data)
	data_slice := strings.Split(data_string, "\n")
	id := 1
	for _, u := range data_slice {
		user := structures.User{}
		userfields := strings.Split(u, ",")
		for _, field := range userfields {
			fields := strings.Split(field, ":")
			fieldName := strings.ReplaceAll(fields[0], " ", "")
			fieldValue := fields[1]
			// TODO wrong load user
			loadToUserStorage(fieldName, fieldValue, &user)

		}
		user.ID = id
		structures.UserStorage = append(structures.UserStorage, user)
		id += 1
	}
}

func loadToUserStorage(fieldName, fieldValue string, user *structures.User) {
	switch fieldName {
	case "name":
		user.Name = fieldValue
	case "email":
		user.Email = fieldValue
	case "password":
		user.Password = fieldValue
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
	return authenticatedUser != nil
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
		file, err = os.OpenFile(userStoragePath, os.O_APPEND, 0644)
		defer file.Close()
		if err == nil {
			user := fmt.Sprintf("name:%s, email:%s, password:%s\n", name, email, password)
			file.Write([]byte(user))
		} else {
			fmt.Printf("error happend when open file : %v", err)
		}
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
