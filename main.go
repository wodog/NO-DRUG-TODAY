package main

import "time"

func main() {
	drug()
	ticker := time.NewTicker(24 * time.Hour)
	for _ = range ticker.C {
		drug()
	}
}
