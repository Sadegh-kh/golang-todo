package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"todo/delivery/deliveryparam"
	"todo/entity"
	"todo/service"
	"todo/structures"
)

const(
	Network = "tcp"
	Address = "127.0.0.1:9999"
)
var (
	scanner = bufio.NewScanner(os.Stdin)
	
)

func main() {
	
	listener, err := net.Dial(Network,Address)
	if err!=nil{
		log.Fatalln("error happend when connecting to network becuse: ",err)
	}
	defer listener.Close()
	command := flag.String("command-task", "exit", "command for create , edit and ...")
	flag.Parse()
	for {
		fmt.Println("local address:",listener.LocalAddr())
		fmt.Println("local address:",listener.RemoteAddr())
		runCommand(*command,listener)
		println("please enter another command or exit")
		scanner.Scan()
		*command = scanner.Text()
	
	 }
	// listener.Write([]byte(*command))

	// response:=make([]byte,1024)
	// numberOfResponseByte,err:=listener.Read(response)
	// if err!=nil{
	// 	log.Fatalln("Error happend when read data from server becuse: ",err)
	// }
	// fmt.Println("response from server:",string(response[:numberOfResponseByte]))


}


func runCommand(command string,listener net.Conn) {
	switch command {
	case "create-task":
		newTask:=createTask()
		req := deliveryparam.Request{Command: command,Task: newTask}
		DataSerialized,err := json.Marshal(req)
		if err != nil{
			log.Fatalln("Error happend when serialize data becuse:",err)
		}

		_,err =listener.Write(DataSerialized)
		if err != nil{
			log.Fatalln("Error happend when send data from client to server becuse:",err)
		}

		var response = make([]byte,1024)
		numberOfData,err:=listener.Read(response)
		if err!=nil{
			log.Fatalln("error happend when Reading from server becuse: ",err)
		}

		task:=&deliveryparam.Request{}
		err=json.Unmarshal(response[:numberOfData],task)
		if err!=nil{
			log.Fatalln("error happend when deserializing data response becuse: ",err)
		}
		fmt.Println("New task created:",task.Task)
		
	case "create-category":
		// TODO
	case "task-list":
		req := deliveryparam.Request{
			Command: command,
		}
		ser,err:=json.Marshal(req)
		if err!=nil{
			log.Fatalln("Error happend when serialing data in task list becuse:",err)
		}

		numberOfBytesData,err:=listener.Write(ser)
		if err!=nil{
			log.Fatalln("Error happend when send data to server in task list becuse:",err)
		}
		fmt.Println("byte sended",numberOfBytesData)

		
		var response = make([]byte,1024)
		numberOfBytesData,err=listener.Read(response)
		if err!=nil{
			log.Fatalln("Error happend when resive data from server in task list becuse:",err)
		}
		listTasks:=service.ListResponse{}
		err=json.Unmarshal(response[:numberOfBytesData],&listTasks)
		fmt.Println("listResponse:",listTasks)
		if err!=nil{
			log.Fatalln("Error happend when desrialize list tasks becuse:",err)
		}

		fmt.Println("List of your task:",listTasks.Tasks)


	case "category-list":
		fmt.Println(structures.GetCategoryList())
	case "exit":
		os.Exit(0)
	default:
		fmt.Println("command is not valid :", command)
	}
}

func createTask() deliveryparam.CreateTask {
	fmt.Println("enter a title for Task:")
	scanner.Scan()
	var title = scanner.Text()

	newCategory := createCategory()

	newTask := deliveryparam.CreateTask{
		Title: title,
		Category: newCategory,
	}

	return newTask
	// fmt.Printf("Task %v Created\n",task)
}

func createCategory() entity.Category {
	fmt.Println("enter a title for category:")
	scanner.Scan()
	var titleCategory = scanner.Text()
	newCategory := entity.Category{Title: titleCategory}
	return newCategory
}