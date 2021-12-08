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

func (s *server) F_GetNumberRebels(ctx context.Context, in *grpc_fulcrum.F_FromLeia) (*grpc_fulcrum.F_ToLeia, error) {
	//Retornar numero de rebeldes en ciudad x
	//Quizas desencadene un merge

	splitted := strings.Fields(in.F_LeiaMSG)
	planeta := splitted[1]
	ciudad := splitted[2]
	inServidor := splitted[3]
	
	reloj := grpc_fulcrum.F_Reloj{X:0, Y:0, Z:0}
	if _, ok := DATA_Reloj[planeta]; ok {
		reloj = DATA_Reloj[planeta]
	}

	if(strconv.Itoa(int(indice_servidor)) != inServidor && inServidor != "-1"){

		if(!(reloj.X >= in.FReloj.X && reloj.Y >= in.FReloj.Y && reloj.Z >= in.FReloj.Z)){
			log.Printf("---------INCONSISTENCIA DETECTADA---------")
			
			//HACER MERGE

			//PROPAGAR CAMBIOS
			
			log.Printf("---------INCONSISTENCIA RESUELTA---------")
			//ClearLog(splitted[1])
		}
		
	}

	var rebeldes int64 = 0

	reloj = grpc_fulcrum.F_Reloj{X:0, Y:0, Z:0}
	if _, ok := DATA_Reloj[planeta]; ok {
		reloj = DATA_Reloj[planeta]
	}

	if _, ok := DATA[planeta][ciudad]; ok {
		rebeldes =  int64(DATA[planeta][ciudad])
	}

	return &grpc_fulcrum.F_ToLeia{FRebeldes : rebeldes, FReloj : &reloj }, nil
}

func (s *server) F_SendCommand(ctx context.Context, in *grpc_fulcrum.F_From_Informante) (*grpc_fulcrum.F_To_Informante, error) {
	//Función que llamarán los informantes para mandar comandos
	//Si el servidor previo es igual al servidor actual, es imposible que se lean datos desactualizados
	//Si los servidores son distintos pero reloj_actual[servidor previo] < reloj_previo[servidor previo], los datos están desactualizado

	splitted := strings.Fields(in.GetFCommand())

	reloj := grpc_fulcrum.F_Reloj{X:0, Y:0, Z:0}
	if _, ok := DATA_Reloj[splitted[1]]; ok {
		reloj = DATA_Reloj[splitted[1]]
	}

	if(indice_servidor != in.FServidor && in.FServidor != -1){

		if(!(reloj.X >= in.FReloj.X && reloj.Y >= in.FReloj.Y && reloj.Z >= in.FReloj.Z)){
			log.Printf("---------INCONSISTENCIA DETECTADA---------")
			
			//Conectarse a los otros servidores
			//servidor0 = este servidor
			//servidor1 = c1.F_Request(planeta)
			//servidor2 = c2.F_Request(planeta)

			//HACER MERGE
			//mezclar los relojes(X=max(s0.X, s1.X, s2.X), Y=....)
			//DATA_Reloj[planeta]
			//aplicar los comandos de los logs de servidor1 y servidor2 a DATA

			//PROPAGAR CAMBIOS
			//c1.F_Merge(DATA_Reloj[planeta], DATA[planeta])
			//c2.F_Merge(DATA_Reloj[planeta], DATA[planeta])
			
			log.Printf("---------INCONSISTENCIA RESUELTA---------")
			//ClearLog(planeta)
		}
	}
	log.Printf("COMANDO RECIBIDO")
	log.Printf(in.GetFCommand())

	//HACER COMANDO
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
			DATA_Reloj[splitted[1]] = AumentarReloj(reloj)
		}
	}else if(splitted[0] == "DeleteCity"){
		//DeleteCity nombre_planeta nombre_ciudad
		if( !ComandoDeleteCity(splitted[1], splitted[2]) ) {
			log.Printf("CIUDAD NO EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			resultado = "ERROR"
		} else {
			resultado = "Deleted " + splitted[1] + " " + splitted[2]
			DATA_Reloj[splitted[1]] = AumentarReloj(reloj)
		}
	}else if(splitted[0] == "UpdateName"){
		//UpdateName nombre_planeta nombre_ciudad nuevo_valor
		if( !ComandoUpdateName(splitted[1], splitted[2], splitted[3]) ) {
			log.Printf("CIUDAD NO EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			resultado = "ERROR"
		} else {
			resultado = splitted[1] + " " + splitted[3] + " " + strconv.Itoa(DATA[splitted[1]][splitted[3]])
			DATA_Reloj[splitted[1]] = AumentarReloj(reloj)
		}
	}else if(splitted[0] == "UpdateNumber"){
		//UpdateNumber nombre_planeta nombre_ciudad nuevo_valor
		habitantes, _ := strconv.Atoi(splitted[3])
		if( !ComandoUpdateNumber(splitted[1], splitted[2], habitantes) ) {
			log.Printf("CIUDAD NO EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			resultado = "ERROR"
		} else {
			resultado = splitted[1] + " " + splitted[2] + " " + strconv.Itoa(DATA[splitted[1]][splitted[2]])
			DATA_Reloj[splitted[1]] = AumentarReloj(reloj)
		}
	}else{
		//Se murió
		resultado = "ERROR"
	}

	if(resultado != "ERROR") {
		WriteLog(in.GetFCommand(), splitted[1])
		UpdatePlanetFile(splitted[1])
	}

	log.Printf("----------------")
	PrintDATA()

	aux := DATA_Reloj[splitted[1]]

	return &grpc_fulcrum.F_To_Informante{ FReloj: &aux, FLog: resultado }, nil
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

func (s *server) F_Request(ctx context.Context, in *grpc_fulcrum.Fantasma) (*grpc_fulcrum.F_Merge_Data, error) {
	//Cuando se encuentre una inconsistencia, se debe llamar F_Request hacia los otros dos sevidores
	//Input: nombre del planeta cuyos registros se quieren actualizar
	//Output: los logs y el reloj correspondiente a ese planeta

	reloj := grpc_fulcrum.F_Reloj{X:0, Y:0, Z:0}
	if _, ok := DATA_Reloj[planeta]; ok {
		reloj = DATA_Reloj[planeta]
	}

	log := contenido de planeta_Log.txt

	return grpc_fulcrum.F_Merge_Data{FReloj: reloj, FLog: grpc_broker.F_Log{Log: log}}, nil
}

func (s *server) F_Merge(ctx context.Context, in *grpc_fulcrum.F_Merge_Data) (*grpc_fulcrum.Fantasma, error) {
	//Funcion que se debe llamar hacia los otros dos servidores luego de haber llamado F_Request
	//Input el nuevo reloj y el diccionario DATA[planeta] y nombre planeta
	//Output: nada en particular

	//Cuando esta funcion sea llamada se debe reemplazar DATA[planeta] por el recibido por input
	DATA[planeta] = nuevo_DATA[planeta]
	DATA_Reloj[planeta] = nuevo_reloj

	//Actualizar los registros de ese planeta(con UpdatePlanetFile(planeta)) y llamar a ClearLog(planeta)
	UpdatePlanetFile(planeta)
	ClearLog(planeta)

	return grpc_fulcrum.Fantasma{Fantasma: 1}, nil
}

func AumentarReloj(reloj grpc_fulcrum.F_Reloj) grpc_fulcrum.F_Reloj {
	if(indice_servidor == 1){
		reloj.X += 1
	} else if(indice_servidor == 2) {
		reloj.Y += 1
	} else {
		reloj.Z += 1
	}
	return reloj
}

func PrintDATA(){
	log.Printf("_______DATA________")
	for planeta, dic2 := range DATA {
		log.Printf("RELOJ: ")
		log.Printf(strconv.Itoa(int(DATA_Reloj[planeta].X)))
		log.Printf(strconv.Itoa(int(DATA_Reloj[planeta].Y)))
		log.Printf(strconv.Itoa(int(DATA_Reloj[planeta].Z)))

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
	file, err := os.OpenFile("./" + index + "/"+planeta+"_Log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
        log.Fatal(err)
    }
    
    defer file.Close()

	if _, err := file.WriteString(comando); err != nil {
		log.Fatalf("Error escribiendo en logs.txt")
		
		return
	}
}

func ClearLog(planeta string){
	//Borrar logs.txt luego de hacer un merge
	index := strconv.Itoa(int(indice_servidor))
	
	if err := os.Truncate("./" + index + "/"+planeta+"_Log.txt", 0); err != nil {
		log.Printf("Error al eliminar contenidos de un archivo: %v", err)
	}
}

func UpdatePlanetFile(planeta string) {
	//Reescribir contenidos del registro del planeta

	index := strconv.Itoa(int(indice_servidor))

	//Crear el archivo o limpiarlo si ya existia
	file, err := os.OpenFile("./" + index + "/"+planeta+".txt", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)

	if err != nil {
        log.Fatal(err)
    }
    
    defer file.Close()

	for ciudad, habitantes := range DATA[planeta] {
		h := strconv.Itoa(habitantes)
		toWrite := planeta + " " + ciudad + " " + h
		if _, err := file.WriteString(toWrite); err != nil {
			log.Fatalf("Error escribiendo en registro planetario de " + planeta)
			
			return
		}
	}
}

//Direcciones de los servidores fulcrum
var servers = [3]string{"localhost:50051", "localhost:50052", "localhost:50053"}
var indice_servidor int64

var DATA map[string]map[string]int
var DATA_Reloj map[string]grpc_fulcrum.F_Reloj

func main() {
	DATA = make(map[string]map[string]int)
	DATA_Reloj = make(map[string]grpc_fulcrum.F_Reloj)

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