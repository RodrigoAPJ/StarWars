FULCRUM: clean
	go run ./Fulcrum/fulcrum.go

BROKER: clean
	go run ./broker/broker.go

INFO: clean
	go run ./informantes/informantes.go

LEIA: clean
	go run ./leia/leia.go

clean: 
	rm -f ./Fulcrum/0/*.txt
	rm -f ./Fulcrum/1/*.txt
	rm -f ./Fulcrum/2/*.txt