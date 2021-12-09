package main

import (
	grpc_broker "my_packages/grpc_broker"
	grpc_fulcrum "my_packages/grpc_fulcrum"
	"log"
	"context"
	"time"
	"google.golang.org/grpc"
	"bufio"
  	"os"
	"strconv"
	"fmt"
	"strings"
)

const (
	address = "10.6.43.108:50053"
)

var c grpc_broker.BrokerClient
var ultimoServidor int64 = -1

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
	
	comando := bufio.NewReader(os.Stdin)
	linea, err := comando.ReadString('\n')

	ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
    defer cancel()

    r, err := c.SendCommand(ctx, &grpc_broker.Command{Command:linea})
    
    if err != nil {
        log.Fatal(err)
    }

	servidor := servers[r.Id]

	log.Printf("Conectando a: " + servidor)

	f := ConectarFulcrum(servidor)
	//Mandar comando ahora a servers{r.id}

	aux_splitted := strings.Fields(linea)
	relojAMandar := grpc_fulcrum.F_Reloj{X:0, Y:0, Z:0}
	if _, ok := DATA_Reloj[aux_splitted[1]]; ok {
		relojAMandar = DATA_Reloj[aux_splitted[1]]
	}

	res, errr := f.F_SendCommand(ctx, &grpc_fulcrum.F_From_Informante{FCommand : linea, FReloj: &relojAMandar, FServidor: ultimoServidor})

	if errr != nil {
        log.Fatal(errr)
    }

	ultimoServidor = r.Id

	splitted := strings.Fields(res.FLog)
	planeta    := splitted[0]
	DATA_Reloj[planeta] = grpc_fulcrum.F_Reloj{ X:(*(res.FReloj)).X, Y:(*(res.FReloj)).Y, Z:(*(res.FReloj)).Z}

	log.Printf("-------------------------------")
	log.Printf("Log recibido:")
	log.Printf(res.FLog)

	ModificarDATA(linea, res.FLog)
	PrintDATA()
}

func PrintDATA(){
	log.Printf("_______DATA________")
	for planeta, dic2 := range DATA {
		fmt.Println("RELOJ: ", strconv.Itoa(int(DATA_Reloj[planeta].X)), strconv.Itoa(int(DATA_Reloj[planeta].Y)), strconv.Itoa(int(DATA_Reloj[planeta].Z)))

		for ciudad, habitantes := range dic2 {
			h := strconv.Itoa(habitantes)
			toPrint := planeta + " " + ciudad + " " + h
        	log.Printf(toPrint)
		}
    }
}

func ModificarDATA(comando string, res string) {
	
	if(res == "ERROR") {
		log.Printf("No se pudo llevar a cabo el comando... :(")
		return
	}

	splitted := strings.Fields(comando)
	accion     := splitted[0]
	planeta    := splitted[1]
	ciudad     := splitted[2]
	habitantes, _ := strconv.Atoi(splitted[3])

	if (accion == "AddCity") {

		if _, ok := DATA[planeta]; !ok {
			// DATA[planeta][ciudad] does not exist -- create it!
			DATA[planeta] = make(map[string]int)
		}

		DATA[planeta][ciudad] = habitantes

	} else if (accion == "DeleteCity") {
		
		if _, ok := DATA[planeta][ciudad]; ok {
			delete(DATA[planeta], ciudad)	
		}

	} else if (accion == "UpdateName") {
		if _, ok := DATA[planeta][ciudad]; ok {
			aux_habitantes := DATA[planeta][ciudad]
			aux_splitted := strings.Fields(res)
			new_ciudad := aux_splitted[1]

			delete(DATA[planeta], ciudad)

			if _, ok := DATA[planeta]; !ok {
				// DATA[planeta][ciudad] does not exist -- create it!
				DATA[planeta] = make(map[string]int)
			}
	
			DATA[planeta][new_ciudad] = aux_habitantes
		}

	} else if (accion == "UpdateNumber") {
		if _, ok := DATA[planeta][ciudad]; ok {
			DATA[planeta][ciudad] = habitantes
		}
	} else {
		log.Printf("Comando invalido")
	}
}

var servers = [3]string{"10.6.43.105:50053", "10.6.43.106:50052", "10.6.43.107:50051"}
var DATA map[string]map[string]int
var DATA_Reloj map[string]grpc_fulcrum.F_Reloj

func main() {
	DATA = make(map[string]map[string]int)
	DATA_Reloj = make(map[string]grpc_fulcrum.F_Reloj)

	Conectar()
	
	for {
		SendComm()
	}
}
