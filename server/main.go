package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/aristorinjuang/grpc/article"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedArticleServiceServer
}

func (s *server) ReadArticles(ctx context.Context, in *pb.Params) (*pb.Articles, error) {
	log.Println("Params:", in.GetQuery(), in.GetPage(), in.GetPerPage())

	return &pb.Articles{
		Articles: []*pb.Article{
			{
				Id:          1,
				Title:       "Title 1",
				Description: "This is the description 1",
			},
			{
				Id:          2,
				Title:       "Title 2",
				Description: "This is the description 2",
			},
			{
				Id:          3,
				Title:       "Title 3",
				Description: "This is the description 3",
			},
		},
	}, nil
}

func init() {
	godotenv.Load()
}

func main() {
	listen, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterArticleServiceServer(s, &server{})
	if err := s.Serve(listen); err != nil {
		log.Panic(err)
	}
}
