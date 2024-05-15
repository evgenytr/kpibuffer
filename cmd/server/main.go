package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/evgenytr/kpibuffer/internal/agent"
	"github.com/evgenytr/kpibuffer/internal/handlers"
	"github.com/evgenytr/kpibuffer/internal/storage/memstorage"
)

func main() {

	appStorage := memstorage.NewStorage()
	storageHandler := handlers.NewStorageHandler(appStorage)

	http.HandleFunc(`/`, storageHandler.GetFactsHandler)
	http.HandleFunc(`/fact`, storageHandler.PostFactHandler)
	http.HandleFunc(`/facts`, storageHandler.PostFactsHandler)

	go http.ListenAndServe(`:8080`, nil)
	go agent.SendToAPI(appStorage)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	<-sigChan
	fmt.Println("shutdown signal")
}
