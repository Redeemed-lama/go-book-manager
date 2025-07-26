package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	support := 90.0
	resistance := 110.0
	price := 100.0

	for step := 1; step <= 50; step++ {
		// Random price movement between -2 and +2
		move := (rand.Float64() * 4) - 2 // -2 <= move < +2
		price += move

		// Bounce off support
		if price < support {
			price = support + (support - price) // bounce back above support
		}
		// Bounce off resistance
		if price > resistance {
			price = resistance - (price - resistance) // bounce back below resistance
		}

		fmt.Printf("Step %d: Price = %.2f\n", step, price)
		time.Sleep(100 * time.Millisecond) // slow down the output for readability
	}
}
