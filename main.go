package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"qasir-supplier/inventory/pb"
	"qasir-supplier/inventory/pkg/endpoint"
	"qasir-supplier/inventory/pkg/service"
	transport "qasir-supplier/inventory/pkg/transport"

	"google.golang.org/grpc"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
		grpcAddr = flag.String("grpc", ":8081", "gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()

	// define our inventory services
	srvProduct := service.NewProdutService()
	srvBrand := service.NewBrandService()

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := endpoint.Endpoints{
		// products endpoint
		GetProductsEndpoint: endpoint.MakeGetProductsEndpoint(srvProduct),

		// brands endpoint
		GetBrandsEndpoint: endpoint.MakeGetBrandsEndpoint(srvBrand),
	}

	// Run HTTP Server
	go func() {
		log.Println("Inventory Service (http) is listening on port", *httpAddr)
		handler := transport.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	// Run GRPC Server
	go func() {
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			log.Println("Error connecting grpc server : ", err)
		}

		log.Println("Inventory Service (grpc) is listening on port", *grpcAddr)

		defer grpcListener.Close()

		handler := transport.NewGRPCServer(ctx, endpoints)
		grpcServer := grpc.NewServer()

		// register products server
		pb.RegisterProductsServer(grpcServer, handler)

		// register brands server
		/*TODO*/

		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Println("Failed To Serve", err)
		}
	}()

	log.Fatalln(<-errChan)
}
