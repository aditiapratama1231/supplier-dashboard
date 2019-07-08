package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	// "net"

	// "qasir-supplier/inventory/pb"
	"qasir-supplier/inventory/database"
	"qasir-supplier/inventory/pkg/endpoint"
	"qasir-supplier/inventory/pkg/service"
	transport "qasir-supplier/inventory/pkg/transport"

	"github.com/joho/godotenv"
	// "google.golang.org/grpc"
)

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	// Connect to database
	db := database.DBInit()

	httpPort := os.Getenv("INVENTORY_HTTP_PORT")
	// grpcPort := os.Getenv("INVENTORY_GRPC_PORT")

	var (
		httpAddr = flag.String("http", ":"+httpPort, "http listen address")
		// grpcAddr = flag.String("grpc", ":"+grpcPort, "gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()

	// define our inventory services
	srvProduct := service.NewProdutService(db)
	srvBrand := service.NewBrandService(db)

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := endpoint.Endpoints{
		// products endpoint
		GetProductsEndpoint:   endpoint.MakeGetProductsEndpoint(srvProduct),
		CreateProductEndpoint: endpoint.MakeCreateProductEndpoint(srvProduct),
		ShowProductsEndpoint:  endpoint.MakeShowProductEndpoint(srvProduct),
		UpdateProductEnpoint:  endpoint.MakeUpdateProductEndpoint(srvProduct),
		DeleteProductEnpoint:  endpoint.MakeDeleteProductEndpoint(srvProduct),

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
	// go func() {
	// 	grpcListener, err := net.Listen("tcp", *grpcAddr)
	// 	if err != nil {
	// 		log.Println("Error connecting grpc server : ", err)
	// 	}

	// 	log.Println("Inventory Service (grpc) is listening on port", *grpcAddr)

	// 	defer grpcListener.Close()

	// 	handler := transport.NewGRPCServer(ctx, endpoints)
	// 	grpcServer := grpc.NewServer()

	// 	// register products server
	// 	pb.RegisterProductsServer(grpcServer, handler)

	// 	// register brands server
	// 	/*TODO*/

	// 	if err := grpcServer.Serve(grpcListener); err != nil {
	// 		log.Println("Failed To Serve", err)
	// 	}
	// }()

	log.Fatalln(<-errChan)
}
