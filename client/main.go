package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "github.com/aristorinjuang/grpc/article"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func init() {
	godotenv.Load()
}

func main() {
	conn, err := grpc.Dial(os.Getenv("SERVER"), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Panic(err)
	}
	defer conn.Close()

	c := pb.NewArticleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ReadArticles(ctx, &pb.Params{
		Query:   "Test...",
		Page:    2,
		PerPage: 10,
	})
	if err != nil {
		log.Panic(err)
	}
	log.Println("Aritcles:", r.GetArticles())
}
