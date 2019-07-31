package ch5

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func Wait() {
	if err := WaitForServer(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "site is down: %v\n", err)
		log.Fatalf("site is down: %v\n", err)
		os.Exit(1)
	}
}

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil
		}
		fmt.Println("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries))
	}
	return fmt.Errorf("server %s failed to respond after %s", url)
}
