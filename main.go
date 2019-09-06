package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/sunjin110/gRPC/server"
	"github.com/sunjin110/gRPC/service"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("===== proto =====")
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	catServer := &service.MyCatService{}

	//実行したい実処理をserverに登録
	pb.RegisterCatServer(server, catServer)
	server.Serve(listenPort)
}
