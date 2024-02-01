package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"todo/structures"
)

var (
	err  error
	file *os.File
)

type FileStorage struct {
	Path              string
	SerializationMode string
}

func (f FileStorage) Save(u structures.User) {
	f.SerializeData(u, f.Path)
}
func (f FileStorage) Create() {
	switch f.SerializationMode {
	case "normal":
		_, err = os.Stat(f.Path)
		// if file exist, err == nil
		if err != nil {
			file, err = os.Create(f.Path)
			if err != nil {
				fmt.Printf("error happed when we create file: %v\n", err)
			}
		}
	case "json":
		_, err = os.Stat(f.Path)
		// if file exist, err == nil
		if err != nil {
			file, err = os.Create(f.Path)
			if err != nil {
				fmt.Printf("error happed when we create file: %v\n", err)
			}
		}
	default:
		fmt.Println("some error happend in serialization mode")

		return
	}
}
func (f FileStorage) Load() {
	id := 1
	var data = make([]byte, 512)
	switch f.SerializationMode {
	case "normal":
		fmt.Println("serialze mode is normal")
		file, err = os.Open(f.Path)
		defer file.Close()
		_, err = file.Read(data)

		dataString := string(data)
		dataString = strings.ReplaceAll(dataString, "\x00", "")

		dataSlice := strings.Split(dataString, "\n")

		user := structures.User{}

		for _, u := range dataSlice {

			if u == "" {
				fmt.Println("continue")
				continue
			}
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
		fmt.Println("serialze mode is json")
		file, err = os.Open(f.Path)
		defer file.Close()
		_, err = file.Read(data)

		dataString := string(data)
		dataSlice := strings.Split(dataString, "\n")

		user := structures.User{}
		for _, u := range dataSlice {
			if u[0] != '{' && u[len(u)-1] != '}' {
				continue
			}
			err = json.Unmarshal([]byte(u), &user)
			if err != nil {
				fmt.Printf("error %v happend when deserializing json format", err)
			}
			structures.UserStorage = append(structures.UserStorage, user)
		}

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

func (f FileStorage) SerializeData(user structures.User, Path string) {
	switch f.SerializationMode {
	case "normal":
		file, err = os.OpenFile(f.Path, os.O_APPEND, 0644)
		defer file.Close()
		if err == nil {
			userStr := fmt.Sprintf("name:%s, email:%s, password:%s\n", user.Name, user.Email, user.Password)
			file.Write([]byte(userStr))
		} else {
			fmt.Printf("error happend when open file : %v", err)
		}

	case "json":
		file, err = os.OpenFile(f.Path, os.O_APPEND, 0644)
		defer file.Close()
		if err == nil {
			userJson, err := json.Marshal(user)
			if err != nil {
				fmt.Println("some error happen when serializing :", err)
			}
			userJson = append(userJson, "\n"...)
			file.Write(userJson)
		} else {
			fmt.Printf("error %v happend when open file\n", err)
		}

	default:
		fmt.Println("some error happend in serializing mode")
	}
}
