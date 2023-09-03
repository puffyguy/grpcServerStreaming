package main

import (
	"context"
	"log"
	"time"

	pb "github.com/puffyguy/grpcServerStreaming/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:5001"

func main() {
	con, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error while connecting: %v", err)
	}

	defer con.Close()

	client := pb.NewAirthmeticServiceClient(con)

	doAirthmetic(client)
}

func doAirthmetic(c pb.AirthmeticServiceClient) {
	log.Println("doAirthmetic function invoked")
	stream, err := c.AirthmeticOperation(context.Background(), &pb.ServerRequest{
		Number1: 2,
		Number2: 10,
	})
	if err != nil {
		log.Fatalf("Error while calling Airthmetic Operation function: %v\n", err)
	}

	res, err := stream.Recv()
	if err != nil {
		log.Fatalf("Error while receiving stream: %v\n", err)
	}

	log.Printf("Sum: %v\n", res.Result.Sum)
	time.Sleep(1 * time.Second)
	log.Printf("Sub: %v\n", res.Result.Sub)
	time.Sleep(1 * time.Second)
	log.Printf("Mul: %v\n", res.Result.Mul)
	time.Sleep(1 * time.Second)
	log.Printf("Div: %v\n", res.Result.Div)
}
