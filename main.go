package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	//"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	//"qasir-supplier/inventory/pb"
	"qasir-supplier/order/database"
	"qasir-supplier/order/pkg/endpoint"
	"qasir-supplier/order/pkg/service"
	transport "qasir-supplier/order/pkg/transport"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	//"google.golang.org/grpc"
)

func main() {
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	db := database.DBInit()

	httpPort := os.Getenv("ORDER_HTTP_PORT")
	//grpcPort := os.Getenv("INVENTORY_GRPC_PORT")

	var (
		httpAddr = flag.String("http", ":"+httpPort, "http listen address")
		//grpcAddr = flag.String("grpc", ":"+grpcPort, "gRPC listen address")
	)
	flag.Parse()
	ctx := context.Background()

	// define our inventory services
	srvOrder := service.NewOrderService(db)

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := endpoint.Endpoints{
		// orders endpoints
		GetMerchantOrdersEndpoint: endpoint.MakeGetMerchantOrdersEndpoint(srvOrder),
		ShowOrderDetailEndpoint:   endpoint.MakeShowOrderDetailEndpoint(srvOrder),
		CreateOrderEndpoint:       endpoint.MakeCreateOrderEndpoint(srvOrder),
		ChangeOrderStatusEndpoint: endpoint.MakeChangeOrderStatusEndpoint(srvOrder),
	}

	// Run HTTP Server
	go func() {
		log.Println("Order Service (http) is listening on port", *httpAddr)
		handler := transport.NewHTTPServer(ctx, endpoints)
		handler = handlers.LoggingHandler(os.Stdout, handler)
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
