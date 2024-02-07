package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"todo/delivery/deliveryparam"
	"todo/service"
)

const(
	Network = "tcp"
	Address = "127.0.0.1:9999"
)
var taskService       = service.NewTaskService()

func main() {
	listener,err := net.Listen(Network,Address)
	if err != nil{
		log.Fatalf("error when listen to network becuse: %v\n",err)
	}
	fmt.Printf("server at %v listening \n",listener.Addr())
	defer listener.Close()

	for {
		connection,err:=listener.Accept()
		if err!=nil{
			fmt.Println("error happend when connecting becuse: ",err)

			continue
		}

		data := make([]byte,1024)
		numberOfData,err:=connection.Read(data)
		if err!=nil{
			fmt.Println("Error happend when read data becuse: ",err)

			continue 
		}

		req := &deliveryparam.Request{}
		err = json.Unmarshal(data[:numberOfData],&req)
		if err != nil{
			fmt.Println("Error happend when deserializing becuse:",err)

			continue
		}

		fmt.Println("command",req.Command)
		switch req.Command{
		case "create-task":
			task :=req.Task
			taskResponse,err:=service.NewTaskService().CreateTask(service.CreateTaskRequest{Title: task.Title,Category: task.Category,AuthenticatedUserID: 0})
			if err!=nil{
				fmt.Println("Error happend when create task becuse: ",err)
				
				continue
			}
			newTask:=deliveryparam.CreateTask{
				Title: taskResponse.Task.Title,
				Category: taskResponse.Task.Category,
			}
			taskSerialize,err := json.Marshal(deliveryparam.Request{Command: req.Command,Task: newTask})
			if err!=nil{
				fmt.Println("Error happend when serializing becuse:",err)
			
				continue
			}

			_,err=connection.Write(taskSerialize)
			if err!=nil{
				fmt.Println("Error happend when send data to client becuse:",err)
			
				continue
			}
		case "task-list":
			tasks,err:=taskService.GetListTask(service.ListRequest{UserID: 0})
			fmt.Println("tasks:",tasks)
			if err!=nil{
				fmt.Println("Error happend when geting list becuse:",err)

				continue
			}

			listTasks,err:=json.Marshal(tasks)
			if err!=nil{
				fmt.Println("Error happend when serilizing list of tasks becuse:",err)

				continue
			}

			sendedData,err:=connection.Write(listTasks)
			if err!=nil{
				fmt.Println("Error happend when send list of tasks to clinet becuse:",err)
			}
			fmt.Println("number of data sended",sendedData)
		}
	}
}