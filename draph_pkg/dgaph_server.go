package dgraph_server


import (
	"google.golang.org/grpc"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
   	"context"
	   "fmt"
)


type Events struct {
	Uid string `json:"uid", omitempty`
	EventName string `json:"eventname", omitempty`
	Data string `json:"data", omitempty`
}


type OpenDgraphDockerProtocol interface {
	Subscribe(eventTrigger Events, port string)(*Events, *grpc.ClientConn, error)
	CloseConnection(grpc_Client *grpc.ClientConn)
}

func Fire() OpenDgraphDockerProtocol{
	return &Events{}
}

func (*Events)Subscribe(event Events, port string) (*Events, *grpc.ClientConn, error){
	dgraph_Connect, err := grpc.Dial(":"+port, grpc.WithInsecure()); if err != nil {
		return &Events{},&grpc.ClientConn{}, err 
	}

	// defer dgraph_Connect.Close()
	client := dgo.NewDgraphClient(api.NewDgraphClient(dgraph_Connect))

	fmt.Println("Client:", client)

	err = client.Alter(context.Background(), &api.Operation{
		Schema: ` 
			   eventname : string @index(fulltext) .
			   Id : string .
			   data : string @index(fulltext) .
		`,
	})
	return &event, dgraph_Connect, err
}

func (*Events) CloseConnection(grpc_Client *grpc.ClientConn){ defer grpc_Client.Close() }