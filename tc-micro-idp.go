package main

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/grpc"
	"log"
	"net"
	"tc-micro-idp/authorizer"
	"tc-micro-idp/models"
	"tc-micro-idp/router"
	"tc-micro-idp/utils/publicFunctions"
)


var app *fiber.App
var grpcServer *grpc.Server

func main() {

	log.Println("ID : ", publicFunctions.IdGenerator.Generate().Int64())
	app = fiber.New(fiber.Config{
		Prefork: false,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,Head,Put,DELETE,PATH",
		AllowHeaders: "",
	}))
	router.SetupRoutes(app)

	lis, err := net.Listen("tcp", ":8000")
	log.Println("done1")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("done2")
	s := authorizer.Server{}
	log.Println("done3")
	grpcServer = grpc.NewServer(grpc.EmptyServerOption{})
	log.Println("done4")

	models.RegisterAuthorizerServer(grpcServer, &s)
	log.Println("done5")
	go grpcStart(lis)

	log.Println("done8")
	err = app.Listen(":3000")
	log.Println("done9")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("done0")

}
func grpcStart(lis net.Listener) {
	log.Println("done6")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s\n", err)
	}
	log.Println("done7")
}


func grpcClient() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("test.idp2.baman.bmn:8080", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := models.NewAccountGrpcClient(conn)

	response, err := c.CreateUserAccount(context.Background(), &models.CreateUserAccountRequest{
		UserId: 0,
	})
	if err != nil {
		log.Fatalf("Error when calling Authorize: %s", err)
	}
	log.Printf("Response from server: \ncode :%v,\nclaims: %s", response.GetOk(), response.ErrorMessage)

}
