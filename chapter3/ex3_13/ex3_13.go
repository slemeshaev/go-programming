// Exercise 3.13: Write const declarations for KB, MB, up through YB
// as compactly as you can.

package main

import "fmt"

const (
	KB float64 = 1000
	MB         = KB * 1000
	GB         = MB * 1000
	TB         = GB * 1000
	PB         = TB * 1000
	EB         = PB * 1000
	ZB         = EB * 1000
	YB         = ZB * 1000
)

func main() {
	fmt.Printf("KB = %.0f\n", KB)
	fmt.Printf("MB = %.0f\n", MB)
	fmt.Printf("GB = %.0f\n", GB)
	fmt.Printf("TB = %.0f\n", TB)
	fmt.Printf("PB = %.0f\n", PB)
	fmt.Printf("EB = %.0f\n", EB)
	fmt.Printf("ZB = %.0f\n", ZB)
	fmt.Printf("YB = %.0f\n", YB) // precision is already floating here
}
