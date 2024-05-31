package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	pb "github.com/victorvbello/gonotes/course/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

const addr = "127.0.0.1:50052"

type Server struct {
	pb.BlogServiceServer
}

type BlogItem struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Img     string             `bson:"img"`
	Author  string             `bson:"author"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:      data.ID.Hex(),
		Author:  data.Author,
		Title:   data.Title,
		Content: data.Content,
	}
}

var collection *mongo.Collection

func main() {

	// Setting mongo connection
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI("mongodb://dbrootus:dbrootpass@localhost:27018/"))

	if err != nil {
		log.Fatalf("Error on mongo.NewClient %v\n", err)
	}

	err = mongoClient.Connect(context.Background())

	if err != nil {
		log.Fatalf("Error on mongoClient.Connect %v\n", err)
	}

	collection = mongoClient.Database("grpcdb").Collection("blog")

	// Setting gRPC server
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Error on net.Listen %v\n", err)
	}

	log.Printf("Blog Listening on %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error on s.Server %v\n", err)
	}
}

func (s *Server) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Printf("-------CreateBlog request: %v-------\n", in)

	blog := BlogItem{
		Author:  in.Author,
		Title:   in.Title,
		Content: in.Content,
	}

	res, err := collection.InsertOne(ctx, blog)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("internal error: %v\n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"cannot cast inserted id",
		)
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}

func (s *Server) ReadBlog(ctx context.Context, in *pb.BlogId) (*pb.Blog, error) {
	log.Printf("-------ReadBlog request: %v------\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"cannot cast id",
		)
	}
	data := &BlogItem{}
	filter := bson.M{"_id": oid}

	res := collection.FindOne(ctx, filter)
	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			"cannot find blog whit the current id",
		)
	}

	return documentToBlog(data), nil
}

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("-------UpdateBlog request: %v------\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			"cannot cast id",
		)
	}

	data := &BlogItem{
		Author:  in.Author,
		Title:   in.Title,
		Content: in.Content,
	}

	res, err := collection.UpdateOne(
		ctx,
		bson.M{"_id": oid},
		bson.M{"$set": data},
	)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			"could not update",
		)
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"cannot find blog whit Id",
		)
	}

	return &emptypb.Empty{}, nil
}

func (s *Server) UploadBlogImg(stream pb.BlogService_UploadBlogImgServer) error {
	log.Printf("-------UploadBlogImg request------\n")

	var blogID, fileName, mimeType string
	var file *os.File
	fileSize := uint32(0)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("os.Getwd error: %v", err)
		return status.Errorf(
			codes.Internal,
			"could not upload img",
		)
	}

	for {
		req, err := stream.Recv()

		if err != nil {
			if err == io.EOF {
				log.Printf("io.EOF error: %v", err)
				break
			}
			log.Printf("stream.Recv error: %v", err)
			return status.Errorf(
				codes.Internal,
				"could not upload img",
			)
		}

		switch reqV := req.Data.(type) {
		case *pb.UploadBlogImgRequest_Metadata:
			blogID = reqV.Metadata.BlogID
			fileName = reqV.Metadata.FileName
			mimeType = reqV.Metadata.MimeType

			if file == nil {
				completeFileBasePath := filepath.Join(currentDir, "course/grpc/blog/server/tmp", blogID)
				if _, err := os.Stat(completeFileBasePath); os.IsNotExist(err) {
					os.MkdirAll(completeFileBasePath, 0700)
				}
				file, err = os.Create(filepath.Join(completeFileBasePath, fileName))
				if err != nil {
					log.Printf("os.Create error: %v", err)
					return status.Errorf(
						codes.Internal,
						"could not upload img",
					)
				}

			}
		case *pb.UploadBlogImgRequest_FileData:
			if file == nil {
				log.Println("file is nil")
				return status.Errorf(
					codes.Internal,
					"could not upload img",
				)
			}
			dataChunk := reqV.FileData
			fileSize += uint32(len(dataChunk))

			dataWrite, err := file.Write(dataChunk)

			if err != nil {
				log.Printf("file.Write error: %v", err)
				return status.Errorf(
					codes.Internal,
					"could not upload img",
				)
			}

			if dataWrite != len(dataChunk) {
				log.Println("file.Write error: written data are different from data chunk")
				return status.Errorf(
					codes.Internal,
					"could not upload img",
				)
			}
		}
	}

	file.Close()

	filePath := filepath.Base(file.Name())

	log.Printf("success upload\n-file %s \n-size: %d\n-mime:%s\n-blogID: %s\n", filePath, fileSize, mimeType, blogID)

	return stream.SendAndClose(&emptypb.Empty{})
}

func (s *Server) ListBlog(in *emptypb.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("-------ListBlogs request------")

	cur, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error on find data from DB: %v", err),
		)
	}

	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Error while decode data from DB: %v", err),
			)
		}

		stream.Send(documentToBlog(data))
	}

	if err := cur.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error on read data from DB: %v", err),
		)
	}

	return nil
}

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("-------DeleteBlog request: %v------\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)

	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("cannot cast id: %v", err),
		)
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})

	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("could not delete: %v", err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"blog not found",
		)
	}

	return &emptypb.Empty{}, nil
}
