package main

import (
	"context"
	"example/ginapi/controllers"
	"example/ginapi/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	mongoconnection := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconnection)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connection to mongo db established successfully")

	usercollection = mongoclient.Database("userdb").Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()

}

func main() {
	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basepath)
	log.Fatal(server.Run(":9090"))
}
