package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	port := os.Getenv("PORT")
	done := make(chan os.Signal, 1)
	go http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	fmt.Printf("Server started at port %s\n", port)
	fmt.Println("The bot is running. Press CTRL-C to exit.")
	signal.Notify(done, os.Interrupt)
	<-done
	fmt.Println("Gracefully shutting down.")
}
