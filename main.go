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

	"qasir-supplier/merchant/pkg/endpoint"
	"qasir-supplier/merchant/pkg/service"
	transport "qasir-supplier/merchant/pkg/transport"
)

func main() {
	var (
		httpAddr = flag.String("http", ":8090", "http listen address")
	)
	flag.Parse()
	ctx := context.Background()

	// define our merchant services
	srvMerchant := service.NewMerchantService()

	errChan := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	// mapping endpoints
	endpoints := endpoint.Endpoints{
		// merchants endpoint
		GetMerchantsEndpoint: endpoint.MakeGetMerchantsEndpoint(srvMerchant),
	}

	// HTTP transport
	go func() {
		log.Println("merchant service is listening on port:", *httpAddr)
		handler := transport.NewHTTPServer(ctx, endpoints)
		errChan <- http.ListenAndServe(*httpAddr, handler)
	}()

	log.Fatalln(<-errChan)
}
