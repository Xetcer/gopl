package main

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})
	cashResult := make(chan bool)

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	go func() {
		cashResult <- Withdraw(400)
	}()

	if got, want := <-cashResult, false; got != want {
		t.Errorf("Cash result = %v, want %v", got, want)
	}
}
