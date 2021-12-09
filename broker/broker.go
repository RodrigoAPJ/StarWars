package main

import (
	grpc_broker "my_packages/grpc_broker"
	grpc_fulcrum "my_packages/grpc_fulcrum"
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

func (s *server) GetNumberRebels(ctx context.Context, in *grpc_broker.FromLeia) (*grpc_broker.ToLeia, error) {
	//Llamar servidor fulcrum aleatorio
	ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
    defer cancel()

	choosenServer := rand.Intn(2) + 1

	//Preguntar numero de rebeldes en ciudad
	f := ConectarFulcrum(servers[choosenServer])

	res, err := f.F_GetNumberRebels(ctx, &grpc_fulcrum.F_FromLeia{F_LeiaMSG: in.LeiaMSG, FReloj: &grpc_fulcrum.F_Reloj{X: in.Reloj.X, Y: in.Reloj.Y, Z: in.Reloj.Z}})

	if err != nil {
        log.Fatal(err)
    }


	//Retornar numero de rebeldes
	return &grpc_broker.ToLeia{Rebeldes : res.FRebeldes, Reloj : &grpc_broker.Reloj{X: res.FReloj.X, Y: res.FReloj.Y, Z: res.FReloj.Z }, Servidor : int64(choosenServer)}, nil
}

func (s *server) SendCommand(ctx context.Context, in *grpc_broker.Command) (*grpc_broker.Servidor, error) {
	//Devolver direccion de un servidor Fulcrum aleatorio
	choosenServer := int64(rand.Intn(3))
	
	log.Printf(in.Command)

	return &grpc_broker.Servidor{Id: choosenServer}, nil
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

var servers = [3]string{"10.6.43.105:50053", "10.6.43.106:50052", "10.6.43.107:50051"}

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