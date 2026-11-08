package main

import (
	"fmt"
	"time"
	"uuid"
)

func main() {
	id := uuid.NewV4()

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		fmt.Printf("%s: %s\n", time.Now().UTC().Format(time.RFC3339Nano), id)
		<-ticker.C
	}
}
