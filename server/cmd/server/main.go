package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

const defaultPort = 8000

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
