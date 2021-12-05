package main

import (
	grpc_fulcrum "my_packages/grpc_fulcrum"
	"context"
	"log"
	"google.golang.org/grpc"
)

const (
	port = ""
)

type server struct {
	grpc_fulcrum.UnimplementedFulcrumServer
}

func (s *server) F_GetNumberRebels(ctx context.Context, in *F_FromLeia) (*F_ToLeia, error) {
	return nil, nil
}
func (s *server) F_SendCommand(ctx context.Context, in *F_From_Informante) (*F_To_Informante, error) {
	return nil, nil
}
func (s *server) F_Request(ctx context.Context, in *Fantasma) (*F_Merge_Data, error) {
	return nil, nil
}
func (s *server) F_Merge(ctx context.Context, in *F_Merge_Data) (*Fantasma, error) {
	return nil, nil
}

//Direcciones de los servidores fulcrum
var servers = [3]string{"localhost:50051", "localhost:50052", "localhost:50053"}

func main() {

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	defer lis.Close()

	s := grpc.NewServer()
	grpc_fulcrum.RegisterBrokerServer(s, &server{})

	//ABAJO DE ESTE IF NADA
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}