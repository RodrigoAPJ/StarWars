package main

import (
	grpc_broker "my_packages/grpc_broker"
	"log"
	"context"
	"time"
	"google.golang.org/grpc"
	"fmt"
	//"strconv"
)

const (
	address = "localhost:50054"
)

var c grpc_broker.BrokerClient

func Conectar(){
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
        
	if err != nil {
		log.Printf("did not connect: %v", err)
		defer conn.Close()
		return
	}

	c = grpc_broker.NewBrokerClient(conn)
}

func SendComm(){  //Mandar comando al broker para recibir una direccion de un servidor Fulcrum
	ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
    defer cancel()
	
	var comando string
	fmt.Scanf("%s", &comando)

    r, err := c.SendCommand(ctx, &grpc_broker.Command{Command:comando})
        
    if err != nil {
        log.Fatal(err)
    }

	log.Printf(servers[r.Id])
	//Mandar comando ahora a servers{r.id}
}

var servers = [3]string{"localhost:50051", "localhost:50052", "localhost:50053"}

func main() {
	Conectar()
	
	for {
		SendComm()
	}
}
