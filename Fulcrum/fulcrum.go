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

func (s *server) F_GetNumberRebels(ctx context.Context, in *grpc_fulcrum.F_FromLeia) (*grpc_fulcrum.F_ToLeia, error) {
	//Retornar numero de rebeldes en ciudad x
	//Quizas desencadene un merge
	return grpc_fulcrum.F_ToLeia{FRebeldes   : grpc_broker.F_Rebeldes{Cantidad: },
								FReloj    : grpc_broker.F_Reloj{X: , Y: , Z: }, 
							 	}, nil
}

func (s *server) F_SendCommand(ctx context.Context, in *grpc_fulcrum.F_From_Informante) (*grpc_fulcrum.F_To_Informante, error) {
	//Función que llamarán los informantes para mandar comandos
	return grpc_fulcrum.F_To_Informante{FReloj    : grpc_broker.F_Reloj{X: , Y: , Z: }, 
										FLog      : grpc_broker.F_Log{Log: }
										}, nil
}

func (s *server) F_Request(ctx context.Context, in *grpc_fulcrum.Fantasma) (*grpc_fulcrum.F_Merge_Data, error) {
	//Funcion que llamará un servidor fulcrum a los otros dos servidores para comenzar el proceso de merge
	//Se deberán retornar los logs y relojes pertinentes para que el nodo que hizo la request
	//determine como hacer el merge
	return grpc_fulcrum.F_Merge_Data{FReloj    : grpc_broker.F_Reloj{X: , Y: , Z: }, 
									FLog      : grpc_broker.F_Log{Log: }
									}, nil
}

func (s *server) F_Merge(ctx context.Context, in *grpc_fulcrum.F_Merge_Data) (*grpc_fulcrum.Fantasma, error) {
	//Se envian los datos de como debe quedar todo con el merge desde el nodo que lo realizó
	return grpc_fulcrum.Fantasma{Fantasma: 1}, nil
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