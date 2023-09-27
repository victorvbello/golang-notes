package main

import (
	"context"
	"io"
	"log"

	pb "github.com/victorvbello/gonotes/course/grpc/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

const addr = "127.0.0.1:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Error on grpc.Dial %v on addr %s\n", err, addr)
	}

	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)                         // valid
	readBlog(c, "aInvalidMongoDBId")        // invalid
	readBlog(c, "65023ef53223f075ed0ec9a6") // not found
	updateBlog(c, id)
	listBlog(c)
	deleteBlog(c, "64ffa56383b3e3551b5d14a9") // not found
	deleteBlog(c, id)                         // valid
}

func createBlog(c pb.BlogServiceClient) string {
	log.Println("-------createBlog request-------")

	newBlog := &pb.Blog{
		Author:  "Victor Bello",
		Title:   "How to learning gRPC",
		Content: "Simple, you just need to read and a lot of practice",
	}

	res, err := c.CreateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("blog was created successfully:  %s", res.Id)
	return res.Id
}

func readBlog(c pb.BlogServiceClient, ID string) *pb.Blog {
	log.Println("-------readBlog request-------")

	res, err := c.ReadBlog(context.Background(), &pb.BlogId{
		Id: ID,
	})

	if err != nil {
		log.Printf("Unexpected error: %v\n", err)
		return nil
	}
	log.Printf("blog was read successfully:  %v", res)
	return res
}

func updateBlog(c pb.BlogServiceClient, ID string) {
	log.Println("-------updateBlog request-------")

	blog := &pb.Blog{
		Id:      ID,
		Author:  "Victor Bello",
		Title:   "How to learning go",
		Content: "Easy, you just need to studied more",
	}

	_, err := c.UpdateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Println("blog was updated successfully")
}

func listBlog(c pb.BlogServiceClient) {
	log.Println("-------listBlog request-------")

	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Unexpected error in stream.Recv: %v\n", err)
		}

		log.Println(res)
	}

	log.Println("list blog was ending successfully")
}

func deleteBlog(c pb.BlogServiceClient, ID string) {
	log.Println("-------deleteBlog request-------")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{
		Id: ID,
	})

	if err != nil {
		log.Printf("Unexpected error: %v\n", err)
		return
	}

	log.Println("blog was deleted successfully")
}
