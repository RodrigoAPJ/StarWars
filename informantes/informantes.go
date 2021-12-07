package main

import (
	grpc_broker "my_packages/grpc_broker"
	grpc_fulcrum "my_packages/grpc_fulcrum"
	"log"
	"context"
	"time"
	"google.golang.org/grpc"
	"fmt"
	"strconv"
	//"strconv"
)

const (
	address = "localhost:50054"
)

var c grpc_broker.BrokerClient
var ultimoServidor int64 = 0
var ultimoReloj = grpc_fulcrum.F_Reloj {X:0, Y:0, Z:0}

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

func ConectarFulcrum(servidor string) grpc_fulcrum.FulcrumClient {
	// Set up a connection to the server.
	conn, err := grpc.Dial(servidor, grpc.WithInsecure(), grpc.WithBlock())
        
	if err != nil {
		log.Printf("did not connect: %v", err)
		defer conn.Close()
	}

	return grpc_fulcrum.NewFulcrumClient(conn)
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

	servidor := servers[r.Id]

	f := ConectarFulcrum(servidor)
	//Mandar comando ahora a servers{r.id}
	res, errr := f.F_SendCommand(ctx, &grpc_fulcrum.F_From_Informante{FCommand : comando, FReloj: &ultimoReloj, FServidor: ultimoServidor})

	if errr != nil {
        log.Fatal(errr)
    }

	ultimoServidor = r.Id
	ultimoReloj = *(res.FReloj)

	s := strconv.Itoa(int(GetReloj(r.Id)))
	log.Printf(s)
}

func GetReloj(indice int64) int64 {
	if(indice == 1){
		return ultimoReloj.X 
	} else if(indice == 2) {
		return ultimoReloj.Y
	} else {
		return ultimoReloj.Z 
	}
}

var servers = [3]string{"localhost:50051", "localhost:50052", "localhost:50053"}

func main() {
	Conectar()
	
	for {
		SendComm()
	}
}
