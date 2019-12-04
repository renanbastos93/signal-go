package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Doc about signals -> http://man7.org/linux/man-pages/man7/signal.7.html
// (kill -9 <PID> == kill -SIGKILL <PID>) -> Kill signal

func handlerSignal(sig chan os.Signal, finesh_program chan int) {
	for {
		s := <-sig
		switch s {
		case syscall.SIGUSR1: // User-defined signal 1
			fmt.Println("received signal USR1 I will return my date", time.Now())
		case syscall.SIGUSR2: // User-defined signal 2
			fmt.Println("received signal USR1 I will return my unixtemp", time.Now().Unix())
		case syscall.SIGTERM: // Termination signal
			fmt.Println("I died")
			finesh_program <- 1
			return // Finesh thread
		}
	}
}

func main() {

	// Create channel to listen signal
	sig := make(chan os.Signal, 1)

	// Notify channel when receive a signal
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)

	// Show PID this program
	fmt.Println(os.Getpid())

	// Created Chanel to close APP when finesh thread
	exit_chan := make(chan int, 1)

	// Up thread to listen and treatment all signals
	go handlerSignal(sig, exit_chan)

	// Close channel to exit program
	<-exit_chan

	fmt.Println("Close program")
}
