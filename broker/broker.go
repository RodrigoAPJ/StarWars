package main

import (
	grpc_broker "my_packages/grpc_broker"
	"context"
	"math/rand"
	"log"
	"net"
	"time"
	"google.golang.org/grpc"
)

const (
	port = "localhost:50054"
)

type server struct {
	grpc_broker.UnimplementedBrokerServer
}

/*func (s *server) GetNumberRebels(ctx context.Context, in *grpc_broker.FromLeia) (*grpc_broker.ToLeia, error) {
	//Llamar servidor fulcrum aleatorio
	choosenServer := int64(rand.Intn(2) + 1)

	//Preguntar numero de rebeldes en ciudad(in)
	

	//Retornar numero de rebeldes
	return grpc_broker.ToLeia{rebeldes   : grpc_broker.Rebeldes{cantidad: },
								reloj    : grpc_broker.Reloj{X: , Y: , Z: }, 
								servidor : grpc_broker.Servidor{id: choosenServer}
							 }, nil
}*/

func (s *server) SendCommand(ctx context.Context, in *grpc_broker.Command) (*grpc_broker.Servidor, error) {
	//Devolver direccion de un servidor Fulcrum aleatorio
	choosenServer := int64(rand.Intn(3))
	
	log.Printf(in.Command)

	return &grpc_broker.Servidor{Id: choosenServer}, nil
}

var servers = [3]string{"localhost:50051", "localhost:50052", "localhost:50053"}

func main() {
	rand.Seed(time.Now().UnixNano())

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	defer lis.Close()

	s := grpc.NewServer()
	grpc_broker.RegisterBrokerServer(s, &server{})

	//ABAJO DE ESTE IF NADA
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}