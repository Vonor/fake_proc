package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Signal-Channel erstellen, um SIGTERM zu empfangen
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM)

	// Goroutine, die auf das Beenden-Signal wartet
	go func() {
		<-sigChan
		fmt.Println("Beenden-Signal erhalten, das Programm wird beendet.")
		os.Exit(0)
	}()

	// Endlosschleife, die jede Sekunde eine Nachricht ausgibt
	for {
		time.Sleep(1 * time.Second)
	}
}
