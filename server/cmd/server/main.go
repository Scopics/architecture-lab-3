package main

import (
	"bufio"
	"database/sql"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/Scopics/architecture-lab-3/server/db"
	"github.com/joho/godotenv"
)

const defaultPort = 8000

func NewDbConnection() (*sql.DB, error) {
	host, _ := os.LookupEnv("HOST")
	port, _ := os.LookupEnv("PORT")
	user, _ := os.LookupEnv("USER")
	password, _ := os.LookupEnv("PASSWORD")

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
	fmt.Println("Enter q to quit")
	for {
		fmt.Print("-> ")
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
		fmt.Print("No .env file found")
	}
}

func main() {
	port := flag.Int("port", defaultPort, "Specify a http port")
	flag.Parse()

	fmt.Printf("Restaurant DBMS server running on port %d\n", *port)

	quit := make(chan string)
	go inputLoop(quit)
	go interrruptLoop(quit)

	quitMessage := <-quit
	fmt.Println("Stopping server...")
	fmt.Println(quitMessage)
}
