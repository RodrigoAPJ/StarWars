package main

import (
	grpc_broker "my_packages/grpc_broker"
	"context"
	"time"
	"bufio"
	"os"
	"strings"
	"strconv"
	"log"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50054"
)

var c grpc_broker.BrokerClient
var ultimoServidor int64 = -1
var DATA_Reloj map[string]grpc_broker.Reloj
var DATA map[string]map[string]int


func SendComm(){  //Mandar comando al broker para recibir una direccion de un servidor Fulcrum	
	//Comando: "GetNumberRebelds nombre_planeta nombre_ciudad"
	comando := bufio.NewReader(os.Stdin)
	linea, err := comando.ReadString('\n')

	splitted := strings.Fields(linea)
	planeta := splitted[1]
	ciudad := splitted[2]

	var relojAMandar grpc_broker.Reloj
	relojAMandar = grpc_broker.Reloj{X:0, Y:0, Z:0}
	if _, ok := DATA_Reloj[planeta]; ok {
		relojAMandar = DATA_Reloj[planeta]
	}

	//Falta mandar el servidor previo

	aux_uS := strconv.Itoa(int(ultimoServidor))
	linea += " " + aux_uS

	ctx, cancel := context.WithTimeout(context.Background(), 60 * time.Second)
    defer cancel()
    r, err := c.GetNumberRebels(ctx, &grpc_broker.FromLeia{ LeiaMSG: linea, Reloj: &relojAMandar})
    
    if err != nil {
        log.Fatal(err)
    }

	ultimoServidor = r.Servidor
	
	DATA_Reloj[planeta] = grpc_broker.Reloj{ X:(*(r.Reloj)).X, Y:(*(r.Reloj)).Y, Z:(*(r.Reloj)).Z}

	if _, ok := DATA[planeta]; !ok {
		// DATA[planeta] does not exist -- create it!
		DATA[planeta] = make(map[string]int)
	}

	DATA[planeta][ciudad] = int(r.Rebeldes)
	log.Printf(strconv.Itoa(DATA[planeta][ciudad]))
}

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

func main() {
	DATA = make(map[string]map[string]int)
	DATA_Reloj = make(map[string]grpc_broker.Reloj)

	Conectar()

	for {
		SendComm()
	}
}