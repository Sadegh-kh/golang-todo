package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"todo/structures"
)

const (
	UserStorageNormalPath = "users_storage.txt"
	UserStorageJsonPath   = "users_storage_json.txt"
)

var (
	scanner           = bufio.NewScanner(os.Stdin)
	authenticatedUser *structures.User
	file              *os.File
	err               error
	serializationMode *string
)

func main() {
	fmt.Println(("Wellcome to Todo Application"))
	command := flag.String("command-task", "exit", "command for create , edit and ...")
	serializationMode = flag.String("serialize-mode", "normal", "serializtion mode for save status")
	flag.Parse()

	CreateFile()

	loadFile()

	fmt.Println(command)
	for {
		runCommand(*command)
		println("please enter another command or exit")
		scanner.Scan()
		*command = scanner.Text()

	}

}

func CreateFile() {
	switch *serializationMode {
	case "normal":
		_, err = os.Stat(UserStorageNormalPath)
		// if file exist, err == nil
		if err != nil {
			file, err = os.Create(UserStorageNormalPath)
			if err != nil {
				fmt.Printf("error happed when we create file: %v\n", err)
			}
		}
	case "json":
		_, err = os.Stat(UserStorageJsonPath)
		// if file exist, err == nil
		if err != nil {
			file, err = os.Create(UserStorageJsonPath)
			if err != nil {
				fmt.Printf("error happed when we create file: %v\n", err)
			}
		}
	default:
		fmt.Println("some error happend in serialization mode")

		return
	}
}

func loadFile() {
	var data = make([]byte, 512)
	id := 1
	switch *serializationMode {
	case "normal":
		fmt.Println("serialze mode is normal")
		file, err = os.Open(UserStorageNormalPath)
		defer file.Close()
		_, err = file.Read(data)

		data_string := string(data)
		data_slice := strings.Split(data_string, "\n")

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
	case "json":
		// decode json format (deserialize)
	default:
		fmt.Println("can't serialize")

		return
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

		SerializeData(name, email, password)
	} else {
		fmt.Printf("this email  %s exist!\n", email)
	}

}

func SerializeData(name, email, password string) {
	user := structures.User{
		Name:     name,
		Email:    email,
		Password: password,
		ID:       len(structures.UserStorage) + 1,
	}
	switch *serializationMode {
	case "normal":
		file, err = os.OpenFile(UserStorageNormalPath, os.O_APPEND, 0644)
		defer file.Close()
		if err == nil {
			userStr := fmt.Sprintf("\nname:%s, email:%s, password:%s", name, email, password)
			file.Write([]byte(userStr))
		} else {
			fmt.Printf("error happend when open file : %v", err)
		}

	case "json":
		file, err = os.OpenFile(UserStorageJsonPath, os.O_APPEND, 0644)
		defer file.Close()
		if err == nil {
			userJson, err := json.Marshal(user)
			if err != nil {
				fmt.Println("some error happen when serializing :", err)
			}
			file.Write(userJson)
		} else {
			fmt.Printf("error %v happend when open file\n", err)
		}

	default:
		fmt.Println("some error happend in serializing mode")
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
