package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"todo/entity"
	"todo/service"
	"todo/service/pkg"
	"todo/storage"
	"todo/structures"
)

const (
	UserStorageNormalPath = "users_storage.txt"
	UserStorageJsonPath   = "users_storage_json.txt"
)

var (
	scanner           = bufio.NewScanner(os.Stdin)
	authenticatedUser *structures.User
	serializationMode *string
	myFile            = storage.FileStorage{Path: "./data.txt"}
	taskService       = service.NewTaskService()
)

func main() {
	fmt.Println(("Wellcome to Todo Application"))
	command := flag.String("command-task", "exit", "command for create , edit and ...")
	serializationMode = flag.String("serialize-mode", "normal", "serializtion mode for save status")
	flag.Parse()
	myFile.SerializationMode = *serializationMode

	CreateStorage(myFile)
	LoadStorage(myFile)
	for {
		runCommand(*command)
		println("please enter another command or exit")
		scanner.Scan()
		*command = scanner.Text()

	}

}
func CreateStorage(s storage.Storage) {
	s.Create()
}

func LoadStorage(s storage.Storage) {
	s.Load()

}

func runCommand(command string) {
	if authedUser() {
		authRquiredCommands(command)
	} else {
		fmt.Println("you must login or register first!")
		switch command {
		case "register":
			register(myFile)
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
		userID := pkg.ListRequest{UserID: authenticatedUser.ID}
		tasks, err := taskService.GetListTask(userID)
		if err != nil {
			fmt.Println("Error happend when geting list of tasks: ", err)
		}
		fmt.Println("List of Tasks: ", tasks)
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

	password = hashPassword(password)

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

func register(storage storage.Storage) {
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

	password = hashPassword(password)

	if !(structures.UserExist(email)) {
		newUser.CreateUser(name, email, password)
		newUser.AppendToStorage()

		storage.Save(newUser)

	} else {
		fmt.Printf("this email  %s exist!\n", email)
	}

}
func hashPassword(password string) string {
	hash := md5.Sum([]byte(password))
	return hex.EncodeToString(hash[:])
}

func createCategory() entity.Category {
	fmt.Println("enter a title for category:")
	scanner.Scan()
	var titleCategory = scanner.Text()
	newCategory := entity.Category{Title: titleCategory}
	return newCategory
}

func createTask() {
	fmt.Println("enter a title for Task:")
	scanner.Scan()
	var title = scanner.Text()

	newCategory := createCategory()

	newTask := pkg.CreateTaskRequest{
		Title:               title,
		Category:            newCategory,
		AuthenticatedUserID: authenticatedUser.ID,
	}
	_, err := taskService.CreateTask(newTask)
	if err != nil {
		fmt.Println("error happend when createing task becuse: ", err)

		return
	}
	println("task created")
}
