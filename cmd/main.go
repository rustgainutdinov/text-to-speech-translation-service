package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"text-to-speech-translation-service/api"
	"text-to-speech-translation-service/pkg/infrastructure"
)

const httpProxyPort = ":8010"
const serviceAddr = "127.0.0.1:8011"

func main() {
	if err := run(); err != nil {
		fmt.Println(err.Error())
	}
}

func run() error {
	go func() {
		if err := runGRPCService(serviceAddr); err != nil {
			fmt.Println(err.Error())
		}
	}()
	return runHTTPProxy(serviceAddr, httpProxyPort)
}

func runGRPCService(serviceAddr string) error {
	lis, err := net.Listen("tcp", serviceAddr)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	api.RegisterTranslationServiceServer(server, &infrastructure.TranslationServer{})
	fmt.Println("starting grpc server at " + serviceAddr)
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
