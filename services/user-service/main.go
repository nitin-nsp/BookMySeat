package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"ticket-booking-platform/proto/user"
	"ticket-booking-platform/services/user-service/internal/config"
	"ticket-booking-platform/services/user-service/internal/server"
	"ticket-booking-platform/services/user-service/internal/svc"

	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"
)

func main() {
	var c config.Config
	data, err := os.ReadFile("etc/user.yaml")
	if err != nil {
		log.Fatal(err)
	}
	if err := yaml.Unmarshal(data, &c); err != nil {
		log.Fatal(err)
	}

	svcCtx := svc.NewServiceContext(c)
	defer svcCtx.DB.Close()

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", c.Host, c.Port))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	user.RegisterUserServiceServer(grpcServer, server.NewUserServer(svcCtx))

	log.Printf("User service listening on %s:%d", c.Host, c.Port)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
}
