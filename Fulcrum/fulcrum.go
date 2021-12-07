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
	"os"
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
		if(!(reloj.X >= in.FReloj.X && reloj.Y >= in.FReloj.Y && reloj.Z >= in.FReloj.Z)){
			log.Printf("---------INCONSISTENCIA DETECTADA---------")
			
			//HACER MERGE

			//PROPAGAR CAMBIOS
			
			log.Printf("---------INCONSISTENCIA RESUELTA---------")
			//ClearLogs()
		}
	}
	log.Printf("COMANDO RECIBIDO:")
	log.Printf(in.GetFCommand())

	//HACER COMANDO
	splitted := strings.Fields(in.GetFCommand())
	var resultado string

	if(splitted[0] == "AddCity") {
		//AddCity nombre_planeta nombre_ciudad [nuevo_valor]
		
		splitted = append(splitted, "0")
		habitantes, _ := strconv.Atoi(splitted[3])
		
		if( !ComandoAddCity(splitted[1], splitted[2], habitantes) ) {
			log.Printf("CIUDAD YA EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			resultado = "ERROR"
		} else {
			resultado = splitted[1] + " " + splitted[2] + " " + strconv.Itoa(DATA[splitted[1]][splitted[2]])
			AumentarReloj()
		}
	}else if(splitted[0] == "DeleteCity"){
		//DeleteCity nombre_planeta nombre_ciudad
		if( !ComandoDeleteCity(splitted[1], splitted[2]) ) {
			log.Printf("CIUDAD NO EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			resultado = "ERROR"
		} else {
			resultado = "Deleted " + splitted[1] + " " + splitted[2]
			AumentarReloj()
		}
	}else if(splitted[0] == "UpdateName"){
		//UpdateName nombre_planeta nombre_ciudad nuevo_valor
		if( !ComandoUpdateName(splitted[1], splitted[2], splitted[3]) ) {
			log.Printf("CIUDAD NO EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			resultado = "ERROR"
		} else {
			AumentarReloj()
			resultado = splitted[1] + " " + splitted[3] + " " + strconv.Itoa(DATA[splitted[1]][splitted[3]])
		}
	}else if(splitted[0] == "UpdateNumber"){
		//UpdateNumber nombre_planeta nombre_ciudad nuevo_valor
		habitantes, _ := strconv.Atoi(splitted[3])
		if( !ComandoUpdateNumber(splitted[1], splitted[2], habitantes) ) {
			log.Printf("CIUDAD NO EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			resultado = "ERROR"
		} else {
			AumentarReloj()
			resultado = splitted[1] + " " + splitted[2] + " " + strconv.Itoa(DATA[splitted[1]][splitted[2]])
		}
	}else{
		//Se murió
		resultado = "ERROR"
	}

	if(resultado != "ERROR") {
		WriteLog(in.GetFCommand())
	}

	log.Printf("----------------")
	PrintDATA()

	return &grpc_fulcrum.F_To_Informante{ FReloj: &reloj, FLog: resultado }, nil
}
 
/*func CheckPlanet(planeta string) bool{
	//Retorna verdadero si existe el planeta en el diccionario Data
	if _, ok := DATA[planeta]; ok {
		// DATA[planeta] existe
		return true
	}

	return false
}*/

func CheckCity(planeta string, ciudad string) bool{
	//Retorna verdadero si existe la ciudad en el planeta en el diccionario Data
	if _, ok := DATA[planeta][ciudad]; ok {
		// DATA[planeta][ciudad] existe
		return true
	}

	return false
}

func ComandoAddCity(planeta string, ciudad string, habitantes int) bool{
	//Agrega una ciudad a un planeta si no existía y retorna true, si ya existía, retorna false

	if _, ok := DATA[planeta]; !ok {
		// DATA[planeta] no existe, se crea
		DATA[planeta] = make(map[string]int)
	}
	if CheckCity(planeta, ciudad) {
		// DATA[planeta][ciudad] no existe, se crea
		return false
	}

	DATA[planeta][ciudad] = habitantes

	return true
}

func ComandoDeleteCity(planeta string, ciudad string) bool{
	//Elimina una ciudad a un planeta y retorna true, si no existía, retorna false

	if !CheckCity(planeta, ciudad) {
		// DATA[planeta][ciudad] no existe
		return false
	} else {
		// DATA[planeta][ciudad] existe, se borra
		delete(DATA[planeta], ciudad)
	}
	
	return true
}

func ComandoUpdateName(planeta string, ciudad string, ciudad_new string) bool{
	//Actualiza el nombre de una ciudad y retorna true, si no existía, retorna false
	//También, si una ciudad con el nuevo nombre ya existía, no hace nada y retorna false
	
	if !CheckCity(planeta, ciudad) {
		// DATA[planeta][ciudad] no existe
		return false
	} else {
		// DATA[planeta][ciudad] existe	
		if !CheckCity(planeta, ciudad_new) {
			// DATA[planeta][ciudad_new] no existe
			aux := DATA[planeta][ciudad]
			ComandoDeleteCity(planeta, ciudad)
			ComandoAddCity(planeta, ciudad_new, aux)
		}else{
			return false
		}
	}

	return true
}

func ComandoUpdateNumber(planeta string, ciudad string, habitantes int) bool{
	//Actualiza la cantidad de rebeldes de una ciudad y retorna true, si no existía retorna false

	if !CheckCity(planeta, ciudad) {
		// DATA[planeta][ciudad] no existe
		return false
	} else {
		// DATA[planeta][ciudad] existe, se actualiza

		DATA[planeta][ciudad] = habitantes
	}

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

func PrintDATA(){
	log.Printf("_______DATA________")
	for planeta, dic2 := range DATA {
		for ciudad, habitantes := range dic2 {
			h := strconv.Itoa(habitantes)
			toPrint := planeta + " " + ciudad + " " + h
        	log.Printf(toPrint)
		}
    }
}

func WriteLog(comando string, planeta string) {
	//Escribir comando recibido en los logs

	index := strconv.Itoa(int(indice_servidor))
	//file, err := os.Create("jugador_"+playerId+"__ronda_1.txt")
	file, err := os.OpenFile("./" + index + "/Log_"+planeta+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
        log.Fatal(err)
    }
    
    defer file.Close()

	if _, err := file.WriteString(comando); err != nil {
		log.Fatalf("Error escribiendo en logs.txt")
		
		return
	}
}

func ClearLogs(){
	//Borrar logs.txt luego de hacer un merge
	index := strconv.Itoa(int(indice_servidor))
	
	for planeta, dic2 := range DATA {
		if err := os.Truncate("./" + index + "/Log_"+planeta+".txt", 0); err != nil {
			log.Printf("Error al eliminar contenidos de un archivo: %v", err)
		}
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
	fmt.Scanf("%d", &indice_servidor)
	
	servidor := servers[indice_servidor]

	log.Printf("Escuchando en " + servidor)

	lis, err := net.Listen("tcp", servidor)
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