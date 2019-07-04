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

	"qasir-supplier/inventory/pkg/endpoint"
	httpserver "qasir-supplier/inventory/pkg/server/http"
	"qasir-supplier/inventory/pkg/service"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8080", "http listen address")
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

	// HTTP transport
	go func() {
		log.Println("inventory service is listening on port:", *httpAddr)
		handler := httpserver.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
