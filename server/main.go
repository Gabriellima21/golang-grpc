package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"

	pb "grpc/publicacao/publicacao"

	"google.golang.org/grpc"
)

type Publicacao struct {
	id              int
	titulo          string
	conteudo        string
	quantidadeLikes int
}

var publicacoes []Publicacao

func (s *server) UpvotePublicacao(ctx context.Context, in *pb.UpvoteRequest) (*pb.PublicacaoResponse, error) {
	publicacao, err := upvote(int(in.GetId()))
	if err != nil {
		return nil, err
	}
	publicacaoResponse := converterPublicacaoParaPublicacaoResponse(*publicacao)
	return &publicacaoResponse, nil
}

func (s *server) CreatePublicacao(ctx context.Context, in *pb.PublicacaoRequest) (*pb.PublicacaoResponse, error) {
	fmt.Println("Recebendo uma nova Publicacao!")
	fmt.Println("Dados: ", in.Title, in.Content)

	publicacao := criarNovaPublicacao(in)
	publicacaoResponse := converterPublicacaoParaPublicacaoResponse(publicacao)
	return &publicacaoResponse, nil
}

func upvote(id int) (*Publicacao, error) {
	publicacao, err := encontrarPeloId(id)
	if err != nil {
		return nil, err
	}
	publicacao.quantidadeLikes += 1
	return publicacao, nil
}

func converterPublicacaoParaPublicacaoResponse(publicacao Publicacao) pb.PublicacaoResponse {
	return pb.PublicacaoResponse{
		Id:      int64(publicacao.id),
		Title:   publicacao.titulo,
		Content: publicacao.conteudo,
		Likes:   int64(publicacao.quantidadeLikes)}
}

func encontrarPeloId(id int) (*Publicacao, error) {
	for i := 0; i < len(publicacoes); i++ {
		if publicacoes[i].id == id {
			return &publicacoes[i], nil
		}
	}
	return nil, errors.New("publicacao nao existe")
}

func criarNovaPublicacao(pub *pb.PublicacaoRequest) Publicacao {
	publicacao := Publicacao{definirNovoId(), pub.GetTitle(), pub.GetContent(), 0}
	publicacoes = append(publicacoes, publicacao)
	return publicacao
}

func definirNovoId() int {
	if len(publicacoes) == 0 {
		return 1
	}
	return publicacoes[len(publicacoes)-1].id + 1
}

type server struct {
	pb.UnimplementedPublicacaoServiceServer
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Falha ao : %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterPublicacaoServiceServer(s, &server{})
	log.Printf("Servidor escutando na porta %v", lis.Addr())
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Falha ao ini: %s", err)
	}

}
