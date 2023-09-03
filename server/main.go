package main

import (
	"log"
	"net"

	pb "github.com/puffyguy/grpcServerStreaming/proto"
	"google.golang.org/grpc"
)

var addr string = "localhost:5001"

type Server struct {
	pb.AirthmeticServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Error while listening: %v\n", err)
	}

	log.Printf("Server listening at: %v\n", addr)

	server := grpc.NewServer()
	pb.RegisterAirthmeticServiceServer(server, &Server{})

	if err = server.Serve(lis); err != nil {
		log.Fatalf("Error while listening: %v\n", err)
	}
}

func (s *Server) AirthmeticOperation(req *pb.ServerRequest, stream pb.AirthmeticService_AirthmeticOperationServer) error {
	log.Printf("Arithmetic Operation function invoked with %v\n", req)

	num1 := req.Number1
	num2 := req.Number2

	if num1 == 0 || num2 == 0 {
		stream.Send(&pb.ServerResponse{
			Result: &pb.Result{
				Sum: 0,
				Sub: 0,
				Mul: 0,
				Div: 0,
			},
		})
	} else {
		for {
			stream.Send(&pb.ServerResponse{
				Result: &pb.Result{
					Sum: float64(num1) + float64(num2),
					Sub: float64(num1) - float64(num2),
					Mul: float64(num1) * float64(num2),
					Div: float64(num1) / float64(num2),
				},
			})

		}
	}
	return nil
}
