package main

import (
	"context"
	"fmt"
	"time"

	pb "grpc/publicacao/publicacao"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	c, conn := conectarComServidor()
	for {
		var comando int
		fmt.Println("-----------------------------------")
		fmt.Println("Bem vindo ao sistema de Publicacoes")
		fmt.Println("-----------------------------------")
		fmt.Println("Digite a opcao desejada:")
		fmt.Println("1- Incluir uma nova Publicacao")
		fmt.Println("2- Dar like em uma Publicacao existente")
		fmt.Println("Qualquer outra tecla para sair")
		fmt.Println("-----------------------------------")
		fmt.Print(":")
		fmt.Scan(&comando)
		if comando == 1 {
			receberDadosECriarNovaPublicacao(c)
		} else if comando == 2 {
			receberIdParaUpvoteEChamarUpvote(c)
		} else {
			break
		}
	}
	conn.Close()
}

func receberDadosECriarNovaPublicacao(c pb.PublicacaoServiceClient) {
	fmt.Println("Digite o titulo: ")
	var titulo string
	fmt.Scan(&titulo)
	fmt.Println("Digite o conteudo: ")
	var conteudo string
	fmt.Scan(&conteudo)
	criarNovaPublicacao(c, titulo, conteudo)
}

func criarNovaPublicacao(c pb.PublicacaoServiceClient, titulo string, conteudo string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	publicacaoResponse, err := c.CreatePublicacao(ctx, &pb.PublicacaoRequest{Title: titulo, Content: conteudo})
	exibirErroSeExistir("Erro ao criar a Publicacao:", err)
	exibirDadosDeRespostaSeExistir(publicacaoResponse, "Publicacao Criada com Sucesso!")
}

func receberIdParaUpvoteEChamarUpvote(c pb.PublicacaoServiceClient) {
	fmt.Println("Digite o id da Publicacao: ")
	var id int
	fmt.Scan(&id)
	upvote(c, id)
}

func upvote(c pb.PublicacaoServiceClient, id int) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	publicacaoResponse, err := c.UpvotePublicacao(ctx, &pb.UpvoteRequest{Id: int64(id)})
	exibirErroSeExistir("Erro ao dar o like:", err)
	exibirDadosDeRespostaSeExistir(publicacaoResponse, "Like adicionado com sucesso!")
}

func exibirErroSeExistir(mensagem string, err error) {
	if err != nil {
		fmt.Println(mensagem, err)
	}
}

func exibirDadosDeRespostaSeExistir(publicacaoResponse *pb.PublicacaoResponse, cabecalho string) {
	if publicacaoResponse != nil {
		fmt.Println("-----------------------------------")
		fmt.Println(cabecalho)
		fmt.Println("Dados da Publicacao:")
		fmt.Println("Id:", publicacaoResponse.GetId())
		fmt.Println("Titulo:", publicacaoResponse.GetTitle())
		fmt.Println("Content:", publicacaoResponse.GetContent())
		fmt.Println("Likes:", publicacaoResponse.GetLikes())
		fmt.Println("-----------------------------------")
	}
}

func conectarComServidor() (pb.PublicacaoServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	exibirErroSeExistir("Nao foi possivel conectar ao servidor", err)
	return pb.NewPublicacaoServiceClient(conn), conn
}
