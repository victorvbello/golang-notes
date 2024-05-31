package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"

	pb "github.com/victorvbello/gonotes/course/grpc/blog/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	addr                = "127.0.0.1:50052"
	uploadFileBatchSize = 100
)

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
	uploadBlogImg(c, id)
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

func uploadBlogImg(c pb.BlogServiceClient, ID string) {
	log.Println("-------uploadBlogImg request-------")

	stream, err := c.UploadBlogImg(context.Background())

	if err != nil {
		log.Fatalf("Unexpected c.UploadBlogImg error: %v\n", err)
	}

	file, fileName, mime, err := getRandomImg()
	if err != nil {
		log.Fatalf("Unexpected getRandomImg error: %v\n", err)
	}

	buf := make([]byte, uploadFileBatchSize)

	fileMetadata := pb.FileMetadata{
		FileName: fileName,
		MimeType: mime,
		BlogID:   ID,
	}

	err = stream.Send(&pb.UploadBlogImgRequest{
		Data: &pb.UploadBlogImgRequest_Metadata{
			Metadata: &fileMetadata,
		},
	})
	if err != nil {
		log.Fatalf("Unexpected stream.Send Metadata error: %v\n", err)
	}

	chunkCount := 1
	for {
		tBytesRead, err := file.Read(buf)

		if err != nil {
			if err == io.EOF {
				log.Println("-------uploadBlogImg EOF-------")
				break
			}
			log.Fatalf("Unexpected file.Read error: %v\n", err)
		}

		dataFileChunk := buf[:tBytesRead]
		err = stream.Send(&pb.UploadBlogImgRequest{
			Data: &pb.UploadBlogImgRequest_FileData{
				FileData: dataFileChunk,
			},
		})
		if err != nil {
			log.Fatalf("Unexpected stream.Send FileData error: %v\n", err)
		}

		log.Printf("-------uploadBlogImg chuck[%d] processed-------\n", chunkCount)
		chunkCount += 1

	}

	_, err = stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Unexpected stream.CloseAndRecv error: %v\n", err)
	}

	log.Println("blog img was successfully uploaded")
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

func getRandomImg() (reader io.Reader, fileName string, mimeType string, err error) {
	url := "https://picsum.photos/200"
	res, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("http.Get error:%v", err)
		return reader, fileName, mimeType, err
	}
	defer res.Body.Close()
	mimeType = res.Header.Get("Content-Type")
	_, params, err := mime.ParseMediaType(res.Header.Get("Content-Disposition"))
	fileName = params["filename"]

	if err != nil {
		err = fmt.Errorf("mime.ParseMediaType error :%v", err)
		return reader, fileName, mimeType, err
	}
	var buf bytes.Buffer
	_, err = io.Copy(&buf, res.Body)
	if err != nil {
		err = fmt.Errorf("io.Copy :%v", err)
		return reader, fileName, mimeType, err
	}

	reader = bytes.NewReader(buf.Bytes())

	return reader, fileName, mimeType, nil
}
