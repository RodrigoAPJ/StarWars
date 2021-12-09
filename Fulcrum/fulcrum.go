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
	"math"
	"bufio"
	"time"
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
			
			//Conectarse a los otros servidores
						
			var c1 grpc_fulcrum.FulcrumClient
			var c2 grpc_fulcrum.FulcrumClient

			if (indice_servidor == 0){
				c1 = ConectarFulcrum(servers[1])
				c2 = ConectarFulcrum(servers[2])
			}else if (indice_servidor == 1){
				c1 = ConectarFulcrum(servers[0])
				c2 = ConectarFulcrum(servers[2])
			}else{
				c1 = ConectarFulcrum(servers[0])
				c2 = ConectarFulcrum(servers[1])
			}

			ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
    		defer cancel()	
			res1, err1 := c1.F_Request(ctx, &grpc_fulcrum.Fantasma{Planeta: planeta})
			if err1 != nil {
				log.Fatal(err1)
			}
			res2, err2 := c2.F_Request(ctx, &grpc_fulcrum.Fantasma{Planeta: planeta})
			if err2 != nil {
				log.Fatal(err2)
			}

			log.Printf("LOG1: " + res1.FLog)
			log.Printf("LOG2: " + res2.FLog)

			fmt.Println("RELOJ1: ", strconv.Itoa(int(res1.FReloj.X)), strconv.Itoa(int(res1.FReloj.Y)), strconv.Itoa(int(res1.FReloj.Z)))
			fmt.Println("RELOJ2: ", strconv.Itoa(int(res2.FReloj.X)), strconv.Itoa(int(res2.FReloj.Y)), strconv.Itoa(int(res2.FReloj.Z)))

			//HACER MERGE

			//mezclar los relojes(X=max(s0.X, s1.X, s2.X), Y=....)
			x := int64(math.Max(float64(DATA_Reloj[planeta].X), math.Max(float64(res1.FReloj.X), float64(res2.FReloj.X))))
			y := int64(math.Max(float64(DATA_Reloj[planeta].Y), math.Max(float64(res1.FReloj.Y), float64(res2.FReloj.Y))))
			z := int64(math.Max(float64(DATA_Reloj[planeta].Z), math.Max(float64(res1.FReloj.Z), float64(res2.FReloj.Z))))

			fmt.Println("NUEVO X,Y,Z: ",strconv.Itoa(int(x))," ",strconv.Itoa(int(y))," ",strconv.Itoa(int(z)))

			nuevo_reloj := grpc_fulcrum.F_Reloj{X: x, Y: y, Z: z}
			DATA_Reloj[planeta] = nuevo_reloj

			//aplicar los comandos de los logs de servidor1 y servidor2 a DATA
			split_RES1 := strings.Split(res1.FLog, "*")
			for _, linea := range split_RES1 {
				ApplyCommand(linea, false)
			}

			split_RES2 := strings.Split(res2.FLog, "*")
			for _, linea := range split_RES2 {
				ApplyCommand(linea, false)
			}

			UpdatePlanetFile(planeta)

			//PROPAGAR CAMBIOS

			index := strconv.Itoa(int(indice_servidor))

			file, err := os.Open("./" + index + "/"+planeta+".txt")
			if err != nil {
				log.Fatalf("failed to open")
			
			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var registros []string
			
			for scanner.Scan() {
				registros = append(registros, scanner.Text())
			}
			file.Close()

			var registro string = ""

			for _, each_line := range registros {
				registro = registro + each_line + "*"
			}

			_, err3 := c1.F_Merge(ctx, &grpc_fulcrum.F_Merge_Data{FReloj: &nuevo_reloj, FLog: registro}) //mandar DATA_Reloj[planeta] y DATA[planeta]
			_, err4 := c2.F_Merge(ctx, &grpc_fulcrum.F_Merge_Data{FReloj: &nuevo_reloj, FLog: registro})
			if err3 != nil {
				log.Fatal(err3)
			}
			if err4 != nil {
				log.Fatal(err4)
			}
			
			log.Printf("---------INCONSISTENCIA RESUELTA---------")
			ClearLog(planeta)
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
	planeta := splitted[1]
	reloj := grpc_fulcrum.F_Reloj{X:0, Y:0, Z:0}
	if _, ok := DATA_Reloj[splitted[1]]; ok {
		reloj = DATA_Reloj[splitted[1]]
	}

	if(indice_servidor != in.FServidor && in.FServidor != -1){

		if(!(reloj.X >= in.FReloj.X && reloj.Y >= in.FReloj.Y && reloj.Z >= in.FReloj.Z)){
			log.Printf("---------INCONSISTENCIA DETECTADA---------")
			
			//Conectarse a los otros servidores
						
			var c1 grpc_fulcrum.FulcrumClient
			var c2 grpc_fulcrum.FulcrumClient

			if (indice_servidor == 0){
				c1 = ConectarFulcrum(servers[1])
				c2 = ConectarFulcrum(servers[2])
			}else if (indice_servidor == 1){
				c1 = ConectarFulcrum(servers[0])
				c2 = ConectarFulcrum(servers[2])
			}else{
				c1 = ConectarFulcrum(servers[0])
				c2 = ConectarFulcrum(servers[1])
			}

			res1, err1 := c1.F_Request(ctx, &grpc_fulcrum.Fantasma{Planeta: planeta})
			if err1 != nil {
				log.Fatal(err1)
			}
			res2, err2 := c2.F_Request(ctx, &grpc_fulcrum.Fantasma{Planeta: planeta})
			if err2 != nil {
				log.Fatal(err2)
			}

			log.Printf("LOG1: " + res1.FLog)
			log.Printf("LOG2: " + res2.FLog)

			fmt.Println("RELOJ1: ", strconv.Itoa(int(res1.FReloj.X)), strconv.Itoa(int(res1.FReloj.Y)), strconv.Itoa(int(res1.FReloj.Z)))
			fmt.Println("RELOJ2: ", strconv.Itoa(int(res2.FReloj.X)), strconv.Itoa(int(res2.FReloj.Y)), strconv.Itoa(int(res2.FReloj.Z)))

			//HACER MERGE

			//mezclar los relojes(X=max(s0.X, s1.X, s2.X), Y=....)
			x := int64(math.Max(float64(DATA_Reloj[planeta].X), math.Max(float64(res1.FReloj.X), float64(res2.FReloj.X))))
			y := int64(math.Max(float64(DATA_Reloj[planeta].Y), math.Max(float64(res1.FReloj.Y), float64(res2.FReloj.Y))))
			z := int64(math.Max(float64(DATA_Reloj[planeta].Z), math.Max(float64(res1.FReloj.Z), float64(res2.FReloj.Z))))

			fmt.Println("NUEVO X,Y,Z: ",strconv.Itoa(int(x))," ",strconv.Itoa(int(y))," ",strconv.Itoa(int(z)))

			nuevo_reloj := grpc_fulcrum.F_Reloj{X: x, Y: y, Z: z}
			DATA_Reloj[planeta] = nuevo_reloj

			//aplicar los comandos de los logs de servidor1 y servidor2 a DATA
			split_RES1 := strings.Split(res1.FLog, "*")
			for _, linea := range split_RES1 {
				ApplyCommand(linea, false)
			}

			split_RES2 := strings.Split(res2.FLog, "*")
			for _, linea := range split_RES2 {
				ApplyCommand(linea, false)
			}

			UpdatePlanetFile(planeta)

			//PROPAGAR CAMBIOS

			index := strconv.Itoa(int(indice_servidor))

			file, err := os.Open("./" + index + "/"+planeta+".txt")
			if err != nil {
				log.Fatalf("failed to open")
		  
			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var registros []string
		  
			for scanner.Scan() {
				registros = append(registros, scanner.Text())
			}
			file.Close()

			var registro string = ""

			for _, each_line := range registros {
				registro = registro + each_line + "*"
			}

			_, err3 := c1.F_Merge(ctx, &grpc_fulcrum.F_Merge_Data{FReloj: &nuevo_reloj, FLog: registro}) //mandar DATA_Reloj[planeta] y DATA[planeta]
			_, err4 := c2.F_Merge(ctx, &grpc_fulcrum.F_Merge_Data{FReloj: &nuevo_reloj, FLog: registro})
			if err3 != nil {
				log.Fatal(err3)
			}
			if err4 != nil {
				log.Fatal(err4)
			}
			
			log.Printf("---------INCONSISTENCIA RESUELTA---------")
			ClearLog(planeta)
		}
	}
	log.Printf("COMANDO RECIBIDO: " + in.GetFCommand())

	resultado := ApplyCommand(in.GetFCommand(), true)

	if(resultado != "ERROR") {
		WriteLog(in.GetFCommand(), splitted[1])
		UpdatePlanetFile(splitted[1])
	}

	PrintDATA()

	aux := DATA_Reloj[splitted[1]]


	return &grpc_fulcrum.F_To_Informante{ FReloj: &aux, FLog: resultado }, nil
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
 
func ApplyCommand(com string, cambiarReloj bool) string{
	if(com == "") {
		return "ERROR"
	}
	split := strings.Fields(com)

	//HACER COMANDO
	var res string

	reloj := grpc_fulcrum.F_Reloj{X:0, Y:0, Z:0}
	if _, ok := DATA_Reloj[split[1]]; ok {
		reloj = DATA_Reloj[split[1]]
	}

	if(split[0] == "AddCity") {
		//AddCity nombre_planeta nombre_ciudad [nuevo_valor]
		
		split = append(split, "0")
		habitantes, _ := strconv.Atoi(split[3])
		
		if( !ComandoAddCity(split[1], split[2], habitantes) ) {
			log.Printf("CIUDAD YA EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			res = "ERROR"
		} else {
			res = split[1] + " " + split[2] + " " + strconv.Itoa(DATA[split[1]][split[2]])
			if(cambiarReloj) {
				DATA_Reloj[split[1]] = AumentarReloj(reloj)
			}
		}
	}else if(split[0] == "DeleteCity"){
		//DeleteCity nombre_planeta nombre_ciudad
		if( !ComandoDeleteCity(split[1], split[2]) ) {
			log.Printf("CIUDAD NO EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			res = "ERROR"
		} else {
			res = "Deleted " + split[1] + " " + split[2]
			if(cambiarReloj) {
				DATA_Reloj[split[1]] = AumentarReloj(reloj)
			}
		}
	}else if(split[0] == "UpdateName"){
		//UpdateName nombre_planeta nombre_ciudad nuevo_valor
		if( !ComandoUpdateName(split[1], split[2], split[3]) ) {
			log.Printf("CIUDAD NO EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			res = "ERROR"
		} else {
			res = split[1] + " " + split[3] + " " + strconv.Itoa(DATA[split[1]][split[3]])
			if(cambiarReloj) {
				DATA_Reloj[split[1]] = AumentarReloj(reloj)
			}
		}
	}else if(split[0] == "UpdateNumber"){
		//UpdateNumber nombre_planeta nombre_ciudad nuevo_valor
		habitantes, _ := strconv.Atoi(split[3])
		if( !ComandoUpdateNumber(split[1], split[2], habitantes) ) {
			log.Printf("CIUDAD NO EXISTIA EN ESE PLANETA POR LO QUE NO SE HIZO COMANDO:\n")
			res = "ERROR"
		} else {
			res = split[1] + " " + split[2] + " " + strconv.Itoa(DATA[split[1]][split[2]])
			if(cambiarReloj) {
				DATA_Reloj[split[1]] = AumentarReloj(reloj)
			}
		}
	}else{
		//Se murió
		res = "ERROR"
	}

	return res
}

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
	if _, ok := DATA_Reloj[in.Planeta]; ok {
		reloj = DATA_Reloj[in.Planeta]
	}

	//log := contenido de planeta_Log.txt
	index := strconv.Itoa(int(indice_servidor))
	if _, err := os.Stat("./"+index+"/"+in.Planeta + "_Log.txt"); err != nil {
		// path/to/whatever not exists
		log.Printf("El archivo ./"+in.Planeta + "_Log.txt no existe")
		return &grpc_fulcrum.F_Merge_Data{FReloj: &reloj, FLog: ""}, nil
	} 
	
    file, err := os.Open("./"+index+"/"+in.Planeta + "_Log.txt")
  
    if err != nil {
        log.Fatalf("Error abriendo archivo")
    }
  
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)
    
	var registros []string
		  
	for scanner.Scan() {
		registros = append(registros, scanner.Text())
	}
	file.Close()

	var registro string = ""

	for _, each_line := range registros {
		registro = registro + each_line + "*"
	}
  
	return &grpc_fulcrum.F_Merge_Data{FReloj: &reloj, FLog: registro}, nil
}

func (s *server) F_Merge(ctx context.Context, in *grpc_fulcrum.F_Merge_Data) (*grpc_fulcrum.Fantasma, error) {
	//Funcion que se debe llamar hacia los otros dos servidores luego de haber llamado F_Request
	//Input: el nuevo reloj y el diccionario DATA[planeta] y nombre planeta
	//Output: nada en particular

	//Cuando esta funcion sea llamada se debe reemplazar DATA[planeta] por el recibido por input
	//borrar DATA[planeta]
	log.Printf("Cambios recibidos en F_Merge: " + in.FLog)
	split := strings.Split(in.FLog, "*")
	planeta := strings.Fields(split[0])[0]

	if _, ok := DATA[planeta]; ok {
		// DATA[planeta] existe
		// DATA[planeta][ciudad] existe, se borra

		for ciudad, _ := range DATA[planeta] {
			delete(DATA[planeta], ciudad)
		}
	}
	
	for i, linea := range split {
		if(linea == ""){
			continue
		}
		fmt.Println("INDICE: ", i)
		fmt.Println("LINEA: ", linea)
		elementos := strings.Fields(linea)

		fmt.Println("---ELEMENTOS---")
		fmt.Println(elementos)
		fmt.Println("--------")

		aux := elementos[2]
		num, _ := strconv.Atoi(aux)
		
		if _, ok := DATA[elementos[0]]; !ok {
			DATA[elementos[0]] = make(map[string]int)
		}

		DATA[elementos[0]][elementos[1]] = num
	}
	DATA_Reloj[planeta] = *in.FReloj

	//Actualizar los registros de ese planeta(con UpdatePlanetFile(planeta)) y llamar a ClearLog(planeta)
	UpdatePlanetFile(planeta)
	ClearLog(planeta)
	PrintDATA()

	return &grpc_fulcrum.Fantasma{Planeta: "1"}, nil
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
	fmt.Println("_______DATA________")
	for planeta, dic2 := range DATA {
		fmt.Println("Planeta " + planeta)
		fmt.Println("RELOJ: ", strconv.Itoa(int(DATA_Reloj[planeta].X)), strconv.Itoa(int(DATA_Reloj[planeta].Y)), strconv.Itoa(int(DATA_Reloj[planeta].Z)))

		for ciudad, habitantes := range dic2 {
        	fmt.Println(planeta + " " + ciudad + " " + strconv.Itoa(habitantes))
		}
		fmt.Println("__________________")
    }
	fmt.Println("\n")
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
	
	//Crear el archivo o limpiarlo si ya existia
	file, err := os.OpenFile("./" + index + "/"+planeta+"_Log.txt", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)

	if err != nil {
        log.Fatal(err)
    }
    
    defer file.Close()
/* 
	if err := os.Truncate("./" + index + "/"+planeta+"_Log.txt", 0); err != nil {
		log.Printf("Error al eliminar contenidos de un archivo: %v", err)
	} */
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
		toWrite := planeta + " " + ciudad + " " + h + "\n"
		if _, err := file.WriteString(toWrite); err != nil {
			log.Fatalf("Error escribiendo en registro planetario de " + planeta)
			
			return
		}
	}
}


func Eventual() {
	
	for {
		time.Sleep(120 * time.Second)
		
		for planeta, _ := range DATA {
			//Conectarse a los otros servidores			
			var c1 grpc_fulcrum.FulcrumClient
			var c2 grpc_fulcrum.FulcrumClient

			if (indice_servidor == 0){
				c1 = ConectarFulcrum(servers[1])
				c2 = ConectarFulcrum(servers[2])
			}else if (indice_servidor == 1){
				c1 = ConectarFulcrum(servers[0])
				c2 = ConectarFulcrum(servers[2])
			}else{
				c1 = ConectarFulcrum(servers[0])
				c2 = ConectarFulcrum(servers[1])
			}

			ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
    		defer cancel()
			res1, err1 := c1.F_Request(ctx, &grpc_fulcrum.Fantasma{Planeta: planeta})
			if err1 != nil {
				log.Fatal(err1)
			}
			res2, err2 := c2.F_Request(ctx, &grpc_fulcrum.Fantasma{Planeta: planeta})
			if err2 != nil {
				log.Fatal(err2)
			}

			//log.Printf("LOG1: " + res1.FLog)
			//log.Printf("LOG2: " + res2.FLog)

			//fmt.Println("RELOJ1: ", strconv.Itoa(int(res1.FReloj.X)), strconv.Itoa(int(res1.FReloj.Y)), strconv.Itoa(int(res1.FReloj.Z)))
			//fmt.Println("RELOJ2: ", strconv.Itoa(int(res2.FReloj.X)), strconv.Itoa(int(res2.FReloj.Y)), strconv.Itoa(int(res2.FReloj.Z)))

			//HACER MERGE

			//mezclar los relojes(X=max(s0.X, s1.X, s2.X), Y=....)
			x := int64(math.Max(float64(DATA_Reloj[planeta].X), math.Max(float64(res1.FReloj.X), float64(res2.FReloj.X))))
			y := int64(math.Max(float64(DATA_Reloj[planeta].Y), math.Max(float64(res1.FReloj.Y), float64(res2.FReloj.Y))))
			z := int64(math.Max(float64(DATA_Reloj[planeta].Z), math.Max(float64(res1.FReloj.Z), float64(res2.FReloj.Z))))

			//fmt.Println("NUEVO X,Y,Z: ",strconv.Itoa(int(x))," ",strconv.Itoa(int(y))," ",strconv.Itoa(int(z)))

			nuevo_reloj := grpc_fulcrum.F_Reloj{X: x, Y: y, Z: z}
			DATA_Reloj[planeta] = nuevo_reloj

			//aplicar los comandos de los logs de servidor1 y servidor2 a DATA
			split_RES1 := strings.Split(res1.FLog, "*")
			for _, linea := range split_RES1 {
				ApplyCommand(linea, false)
			}

			split_RES2 := strings.Split(res2.FLog, "*")
			for _, linea := range split_RES2 {
				ApplyCommand(linea, false)
			}

			UpdatePlanetFile(planeta)

			//PROPAGAR CAMBIOS

			index := strconv.Itoa(int(indice_servidor))

			file, err := os.Open("./" + index + "/"+planeta+".txt")
			if err != nil {
				log.Fatalf("failed to open")
			}
			scanner := bufio.NewScanner(file)
			scanner.Split(bufio.ScanLines)
			var registros []string
			
			for scanner.Scan() {
				registros = append(registros, scanner.Text())
			}
			file.Close()

			var registro string = ""

			for _, each_line := range registros {
				registro = registro + each_line + "*"
			}

			_, err3 := c1.F_Merge(ctx, &grpc_fulcrum.F_Merge_Data{FReloj: &nuevo_reloj, FLog: registro}) //mandar DATA_Reloj[planeta] y DATA[planeta]
			_, err4 := c2.F_Merge(ctx, &grpc_fulcrum.F_Merge_Data{FReloj: &nuevo_reloj, FLog: registro})
			if err3 != nil {
				log.Fatal(err3)
			}
			if err4 != nil {
				log.Fatal(err4)
			}
			
			ClearLog(planeta)
		}
	}
}

//Direcciones de los servidores fulcrum
var servers = [3]string{"10.6.43.105:50053", "10.6.43.106:50052", "10.6.43.107:50051"}
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

	go Eventual()

	//ABAJO DE ESTE IF NADA
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}