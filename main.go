package main

import(
	."github.com/ali2210/verbose-waffle/draph_pkg"
	"fmt"
	// "google.golang.org/grpc"
	// "github.com/dgraph-io/dgo/v2"
	// "github.com/dgraph-io/dgo/v2/protos/api"
   	// "context"
	//    "fmt"
)

var(
	trigger *Events = &Events{}
)


func main() {
		
		
	event,client,  err := trigger.Subscribe(Events{
		Uid: "_sub0",
		EventName : "subcbe",
		Data : "HelloWorld",
	});if err != nil {
		fmt.Println("Error", err)
		return 
	}
	fmt.Println(event, err)
	trigger.CloseConnection(client)
}


















