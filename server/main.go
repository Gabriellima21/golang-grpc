package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "example/web-service-gin/publicacao"

	"google.golang.org/grpc"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	pb.UnimplementedPublicacaoServiceServer
}

func (s *server) UpvotePublicacao(ctx context.Context, in *pb.PublicacaoRequest) (*pb.PublicacaoResponse, error) {
	log.Printf("Received: %v", in.GetId())
	return &pb.PublicacaoResponse{
		Id:      10,
		Title:   "A",
		Content: "B",
		Likes:   10,
		Dislike: 7}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPublicacaoServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}

var publicacaoCollection *mongo.Collection
var ctx = context.TODO()

type Publicacao struct {
	Id      int
	Title   string
	Content string
	Likes   int
	Dislike int
}

func getPublicacoesById(id int) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	publicacaoCollection = client.Database("klever").Collection("publicacao")

	filter := bson.D{{"id", id}}
	var result Publicacao
	err = publicacaoCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	result.Likes = result.Likes + 1
	res, err := publicacaoCollection.UpdateOne(ctx,
		bson.M{"id": id},
		bson.D{
			{"$set", bson.D{{"likes", result.Likes}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", res.MatchedCount, res.ModifiedCount)
}
