package main

import (
	"fmt"
	"log"
	"net"
	"web-solutions-rpc-golang/server"
)

func main(){
	server := server.NewServer()
	port := 5021

	netListener, listenerErr := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if listenerErr != nil {
		log.Fatalf("Could not listen on port %d.\n%s", port, listenerErr)
	}

	fmt.Printf("Running server on port: %d\n", port)

	if err := server.Serve(netListener); err != nil {
		log.Fatal(err)
	}
}