package main

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"text-to-speech-translation-service/api"
	server2 "text-to-speech-translation-service/pkg/infrastructure/transport"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
	}
}

func run() error {
	envConf, err := server2.ParseEnv()
	if err != nil {
		return err
	}
	go func() {
		if err := runGRPCService(envConf); err != nil {
			fmt.Println(err.Error())
		}
	}()
	return runHTTPProxy(envConf.GRPCAddress, envConf.HTTPProxyAddress)
}

func runGRPCService(envConf *server2.Config) error {
	db := pg.Connect(&pg.Options{
		User:     envConf.DBUser,
		Password: envConf.DBPass,
		Addr:     envConf.DBHost + ":" + envConf.DBPort,
		Database: envConf.DBName,
	})
	rabbitMqChannel, err := getRabbitMqChannel(envConf)
	if err != nil {
		return err
	}
	dependencyContainer := server2.NewDependencyContainer(db, *envConf, rabbitMqChannel)
	lis, err := net.Listen("tcp", envConf.GRPCAddress)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	api.RegisterTranslationServiceServer(server, &server2.TranslationServer{DependencyContainer: dependencyContainer})
	fmt.Println("starting grpc transport at " + envConf.GRPCAddress)
	return server.Serve(lis)
}

func getRabbitMqChannel(envConf *server2.Config) (*amqp.Channel, error) {
	rabbitMqInfo := fmt.Sprintf("amqp://%s:%s@%s//", envConf.RabbitMqUser, envConf.RabbitMqPass, envConf.RabbitMqHost)
	conn, err := amqp.Dial(rabbitMqInfo)
	if err != nil {
		return nil, err
	}
	return conn.Channel()
}

func runHTTPProxy(serviceAddr string, httpProxyPort string) error {
	grpcConn, err := grpc.Dial(serviceAddr, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer grpcConn.Close()
	grpcGWMux := runtime.NewServeMux()
	err = api.RegisterTranslationServiceHandler(context.Background(), grpcGWMux, grpcConn)
	if err != nil {
		return err
	}
	mux := http.NewServeMux()
	mux.Handle("/", grpcGWMux)
	fmt.Println("starting http transport at " + httpProxyPort)
	return http.ListenAndServe(httpProxyPort, mux)
}
