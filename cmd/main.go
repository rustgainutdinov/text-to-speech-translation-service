package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"text-to-speech-translation-service/api"
	"text-to-speech-translation-service/pkg/infrastructure"
)

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
	}
}

func run() error {
	envConf, err := infrastructure.ParseEnv()
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

func runGRPCService(envConf *infrastructure.Config) error {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", envConf.DBUser, envConf.DBPass, envConf.DBName, envConf.DBPort, envConf.DBHost)
	db, err := sqlx.Open("postgres", dbInfo)
	if err != nil {
		return err
	}
	dependencyContainer := infrastructure.NewDependencyContainer(db, *envConf)
	lis, err := net.Listen("tcp", envConf.GRPCAddress)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	api.RegisterTranslationServiceServer(server, &infrastructure.TranslationServer{DependencyContainer: dependencyContainer})
	fmt.Println("starting grpc server at " + envConf.GRPCAddress)
	return server.Serve(lis)
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
	fmt.Println("starting http server at " + httpProxyPort)
	return http.ListenAndServe(httpProxyPort, mux)
}
