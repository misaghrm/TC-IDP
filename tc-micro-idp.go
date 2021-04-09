package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
	"tc-micro-idp/authorizer"
	"tc-micro-idp/models"
	"tc-micro-idp/router"
)

//var app *fiber.App
var grpcServer *grpc.Server
var err error
var lis net.Listener

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,Head,Put,DELETE,PATH",
		AllowHeaders: "",
	}))
	router.SetupRoutes(app)

	log.Println("done8")
	err = app.Listen(":3000")
	log.Println("done9")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("done10")

}
func grpcStart() {
	lis, err = net.Listen("tcp", getPort())
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

type port struct {
	mu   *sync.Mutex
	Port int64
}

var a = port{Port: 8000}

func getPort() (b string) {
	a.mu.Lock()
	a.Port = a.Port + 1
	b = fmt.Sprintf(":%d", a.Port)
	log.Println(b)
	a.mu.Unlock()
	return
}
