package main

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"text-to-speech-translation-service/api"
	"text-to-speech-translation-service/pkg/infrastructure"
	server2 "text-to-speech-translation-service/pkg/infrastructure/transport"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err.Error())
	}
}

func run() error {
	envConf, err := infrastructure.ParseEnv()
	log.SetFormatter(&log.JSONFormatter{})
	if err != nil {
		return err
	}
	go func() {
		if err := runGRPCService(envConf); err != nil {
			log.Fatal(err.Error())
		}
	}()
	return runHTTPProxy(envConf.GRPCAddress, envConf.HTTPProxyAddress)
}

func runGRPCService(envConf *infrastructure.Config) error {
	rabbitMqChannel, err := getRabbitMqChannel(envConf)
	if err != nil {
		return err
	}
	dependencyContainer, err := infrastructure.NewDependencyContainer(getDataBaseConnect(envConf), *envConf, rabbitMqChannel)
	if err != nil {
		return err
	}
	lis, err := net.Listen("tcp", envConf.GRPCAddress)
	if err != nil {
		return err
	}
	server := grpc.NewServer(grpc.UnaryInterceptor(makeGRPCUnaryInterceptor()))
	api.RegisterTranslationServiceServer(server, &server2.TranslationServer{DependencyContainer: dependencyContainer})
	log.WithFields(log.Fields{"grpc address": envConf.GRPCAddress}).Info("successfully starting grpc transport")
	return server.Serve(lis)
}

func getDataBaseConnect(envConf *infrastructure.Config) pg.DBI {
	return pg.Connect(&pg.Options{
		User:     envConf.DBUser,
		Password: envConf.DBPass,
		Addr:     envConf.DBHost + ":" + envConf.DBPort,
		Database: envConf.DBName,
	})
}

func getRabbitMqChannel(envConf *infrastructure.Config) (*amqp.Channel, error) {
	rabbitMqInfo := fmt.Sprintf("amqp://%s:%s@%s//", envConf.RabbitMqUser, envConf.RabbitMqPass, envConf.RabbitMqHost)
	conn, err := amqp.Dial(rabbitMqInfo)
	if err != nil {
		return nil, err
	}
	return conn.Channel()
}

func makeGRPCUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		resp, err = handler(ctx, req)
		return resp, server2.TranslateError(err)
	}
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
	log.WithFields(log.Fields{"http port": httpProxyPort}).Info("successfully starting http transport")
	return http.ListenAndServe(httpProxyPort, mux)
}
