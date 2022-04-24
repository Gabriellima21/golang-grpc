package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "example/web-service-gin/publicacao"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultId = 5
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	Id   = flag.Int64("Id", defaultId, "Id to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPublicacaoServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.UpvotePublicacao(ctx, &pb.PublicacaoRequest{Id: *Id})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %d", r.GetLikes())
}
