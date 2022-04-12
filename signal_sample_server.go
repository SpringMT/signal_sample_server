package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	http.HandleFunc("/", hello)
	http.HandleFunc("/healthz", health)
	server := &http.Server{
		Addr: ":"+port,
		Handler: nil,
	}
	go signals()

	log.Printf("Server listening on port %s and PID %d", port, os.Getpid())
	log.Fatal(server.ListenAndServe())
}

func signals() {
	captureSignal := make(chan os.Signal, 1)
	signal.Notify(captureSignal, syscall.SIGTERM, syscall.SIGKILL, os.Interrupt)
	sig := <-captureSignal
	switch sig {
	case syscall.SIGHUP:
		log.Printf("SIGHUP")
	case syscall.SIGINT:
		log.Printf("SIGINT")
	case syscall.SIGTERM:
		log.Printf("START SIGTERM ")
		time.Sleep(30 * time.Second)
		log.Printf("END SIGTERM")
	case syscall.SIGKILL:
		log.Printf("SIGKILL")
	default:
		log.Printf("SIG %v", sig)
	}
	log.Printf("server shutdown now")
	os.Exit(0)
}

func hello(w http.ResponseWriter, r *http.Request) {
	host, _ := os.Hostname()
	fmt.Fprintf(w, "Hello, world!\n")
	fmt.Fprintf(w, "Version: 0.1.0\n")
	fmt.Fprintf(w, "Hostname: %s\n", host)
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

