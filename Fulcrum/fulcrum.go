package main

import (
	grpc_fulcrum "my_packages/grpc_fulcrum"
	"context"
	"log"
	"google.golang.org/grpc"
	"strings"
	"strconv"
	"fmt"
	"net"
)

const (
	port = ""
)

type server struct {
	grpc_fulcrum.UnimplementedFulcrumServer
}

/*func (s *server) F_GetNumberRebels(ctx context.Context, in *grpc_fulcrum.F_FromLeia) (*grpc_fulcrum.F_ToLeia, error) {
	//Retornar numero de rebeldes en ciudad x
	//Quizas desencadene un merge
	return grpc_fulcrum.F_ToLeia{FRebeldes   : grpc_broker.F_Rebeldes{Cantidad: },
								FReloj    : grpc_broker.F_Reloj{X: , Y: , Z: }, 
							 	}, nil
}*/

func (s *server) F_SendCommand(ctx context.Context, in *grpc_fulcrum.F_From_Informante) (*grpc_fulcrum.F_To_Informante, error) {
	//Función que llamarán los informantes para mandar comandos
	//Si el servidor previo es igual al servidor actual, es imposible que se lean datos desactualizados
	//Si los servidores son distintos pero reloj_actual[servidor previo] < reloj_previo[servidor previo], los datos están desactualizado
	if(indice_servidor != in.FServidor){
		if(reloj.X < in.FReloj.X || reloj.Y < in.FReloj.Y || reloj.Z < in.FReloj.Z){
			//HACER MERGE


			//PROPAGAR CAMBIOS
		}
	} 

	//HACER COMANDO
	splitted := strings.Fields(in.GetFCommand())

	if(splitted[0] == "AddCity") {

		splitted = append(splitted, "0")
		habitantes, _ := strconv.Atoi(splitted[2])
		if( !ComandoAddCity(splitted[1], splitted[2], habitantes) ) {
			log.Printf("CIUDAD YA EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			log.Printf(in.GetFCommand())
		}
	}

	AumentarReloj()

	return &grpc_fulcrum.F_To_Informante{ FReloj: &reloj, FLog: in.FCommand }, nil
}

func ComandoAddCity(planeta string, ciudad string, habitantes int) bool{

	if _, ok := DATA[planeta]; !ok {
		// DATA[planeta][ciudad] does not exist -- create it!
		DATA[planeta] = make(map[string]int)
	} else {
		return false
	}

	DATA[planeta][ciudad] = habitantes

	return true
}

/*func (s *server) F_Request(ctx context.Context, in *grpc_fulcrum.Fantasma) (*grpc_fulcrum.F_Merge_Data, error) {
	//Funcion que llamará un servidor fulcrum a los otros dos servidores para comenzar el proceso de merge
	//Se deberán retornar los logs y relojes pertinentes para que el nodo que hizo la request
	//determine como hacer el merge
	return grpc_fulcrum.F_Merge_Data{FReloj: grpc_broker.F_Reloj{X: , Y: , Z: }, FLog: grpc_broker.F_Log{Log: }}, nil
}*/

/*func (s *server) F_Merge(ctx context.Context, in *grpc_fulcrum.F_Merge_Data) (*grpc_fulcrum.Fantasma, error) {
	//Se envian los datos de como debe quedar todo con el merge desde el nodo que lo realizó
	return grpc_fulcrum.Fantasma{Fantasma: 1}, nil
}*/

func AumentarReloj() {
	if(indice_servidor == 1){
		reloj.X += 1
	} else if(indice_servidor == 2) {
		reloj.Y += 1
	} else {
		reloj.Z += 1
	}
}

//Direcciones de los servidores fulcrum
var servers = [3]string{"localhost:50051", "localhost:50052", "localhost:50053"}
var indice_servidor int64
var reloj = grpc_fulcrum.F_Reloj {X:0, Y:0, Z:0}

var DATA map[string]map[string]int

func main() {
	DATA = make(map[string]map[string]int)

	//Recibir 1 2 o 3 y abrir un server distinto
	log.Printf("Ingrese servidor a iniciar (0|1|2)")
	fmt.Scanf("%i", &indice_servidor)
	
	lis, err := net.Listen("tcp", servers[indice_servidor])
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}
	defer lis.Close()

	s := grpc.NewServer()
	grpc_fulcrum.RegisterFulcrumServer(s, &server{})

	//ABAJO DE ESTE IF NADA
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}