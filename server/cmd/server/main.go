package main

import (
	"bufio"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"

	"github.com/Scopics/architecture-lab-3/server/db"
	"github.com/joho/godotenv"
)

const defaultPort = 8080

func NewDbConnection() (*sql.DB, error) {
	host, _ := os.LookupEnv("HOST_DB")
	port, _ := os.LookupEnv("PORT_DB")
	user, _ := os.LookupEnv("USER_DB")
	password, _ := os.LookupEnv("PASSWORD_DB")

	connection := &db.Connection{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		DbName:   "restaurant",
	}
	return connection.Open()
}

func inputLoop(quit chan string) {
	reader := bufio.NewReader(os.Stdin)
	log.Println("Enter q to quit")
	for {
		log.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if strings.Compare("q", text) == 0 {
			quit <- "Quitted using q"
		}
	}
}

func interrruptLoop(quit chan string) {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt)
	<-sigChannel
	quit <- "Interrupted server"
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	var port int
	flag.IntVar(&port, "port", defaultPort, "Specify a http port")
	flag.IntVar(&port, "p", defaultPort, "Specify a http port (shorthand)")
	flag.Parse()

	if server, err := ComposeApiServer(port); err == nil {
		log.Printf("Restaurant DBMS server running on port %d\n", port)
		go func() {
			err := server.Start()
			if err == http.ErrServerClosed {
				log.Printf("Server stopped")
			} else {
				log.Fatalf("Cannot start HTTP server: %s", err)
			}

		}()

		quit := make(chan string)
		go inputLoop(quit)
		go interrruptLoop(quit)

		quitMessage := <-quit

		if err := server.Stop(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error stopping the server: %s", err)
		} else {
			log.Println(quitMessage)
		}
	} else {
		log.Fatalf("Cannot initialize server: %s", err)
	}

}
