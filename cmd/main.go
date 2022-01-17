package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	cache "github.com/EmreAyberk/go-ddd-example/pkg/cache"
	"github.com/EmreAyberk/go-ddd-example/pkg/item"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, os.Interrupt)

	itemHandler := ItemContainer()

	mux := http.NewServeMux()
	mux.HandleFunc("/items", itemHandler.Handle)

	go func() {
		log.Fatal(http.ListenAndServe(":5000", mux))
	}()
	<-sig
	fmt.Println("Graceful Shutdown!!")
}

func ItemContainer() item.Handle {
	memoryCache := cache.NewMemoryCache()

	itemRepo := item.NewRepository()
	itemService := item.NewService(itemRepo)
	itemHandler := item.NewHandler(itemService, memoryCache)
	return itemHandler
}
