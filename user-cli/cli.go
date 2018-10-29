package main

import (
	"context"
	pb "github.com/caquillo07/grpc-demo-shipping-containers/user-service/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	microclient "github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"log"
	"os"
)

func main() {
	cmd.Init()

	client := pb.NewUserServiceClient("go.micro.srv.user", microclient.DefaultClient)

	// define our flags
	service := micro.NewService(
		micro.Flags(
			cli.StringFlag{
				Name:  "name",
				Usage: "Your full name",
			},
			cli.StringFlag{
				Name:  "email",
				Usage: "Your email",
			},
			cli.StringFlag{
				Name:  "password",
				Usage: "Your password",
			},
			cli.StringFlag{
				Name:  "company",
				Usage: "Your company",
			},
		),
	)

	// start as service
	service.Init(
		micro.Action(action(client)),
	)
}

func action(client pb.UserServiceClient) func(c *cli.Context) {
	return func(c *cli.Context) {
		name := c.String("name")
		email := c.String("email")
		password := c.String("password")
		company := c.String("company")

		// Call our user service
		r, err := client.Create(context.TODO(), &pb.User{
			Name:     name,
			Email:    email,
			Password: password,
			Company:  company,
		})
		if err != nil {
			log.Fatalf("could not create: %v\n", err)
		}
		log.Printf("created user: %s", r.User.Id)

		getAll, err := client.GetAll(context.Background(), &pb.Request{})
		if err != nil {
			log.Fatalf("could not list users: %v", err)
		}

		for _, v := range getAll.Users {
			log.Println(v)
		}

		os.Exit(0)
	}
}
