package main

import(
	."github.com/ali2210/verbose-waffle/draph_pkg"
	"fmt"
	"os"
)

var(
	trigger *Events = &Events{}
)


func main() {
		
	port := os.Getenv("Port")	
	event,client,  err := trigger.Subscribe(Events{
		Uid: "_sub0",
		EventName : "subcbe",
		Data : "HelloWorld",
	}, port);if err != nil {
		fmt.Println("Error", err)
		return 
	}
	fmt.Println(event, err)
	trigger.CloseConnection(client)
}


















