package main

import (
	"fmt"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	clean := run()

	defer clean()

	closeGrace()
}

func run() func() {
	newApp := BuildServer()

	go newApp.Listen(":8081")

	return func() {
		fmt.Println("Closing server...")
		newApp.Shutdown()
	}

}

func closeGrace() {
	fmt.Println("starting gracefully...")
	quit := make(chan os.Signal, 1)
	defer close(quit) //cerrar siempre el channel de tipo os.Signal
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit //hasta no recibir signal en quit no permite terminar funcion
}
