package main

import (
	"qasir-supplier/inventory/hello"
	"qasir-supplier/inventory/router"
	"log"
)

func main() {
	log.Println(hello.Name("Refactory"))

	r := router.Router()
	r.Run()
}
