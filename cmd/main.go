package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/EmreAyberk/goLangCrud/pkg/item"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, os.Interrupt)

	itemRepo := item.NewRepository()
	itemService := item.NewService(itemRepo)
	itemHandler := item.NewHandler(itemService)

	mux := http.NewServeMux()
	mux.HandleFunc("/items", itemHandler.Handle)

	go func() {
		log.Fatal(http.ListenAndServe(":5000", mux))
	}()
	<-sig
	fmt.Println("Graceful Shutdown!!")
}
