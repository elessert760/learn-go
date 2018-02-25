package main

import (
	"fmt"
	"github.com/beevik/ntp" // Using pre-made NTP protocol implementation
)

func main() {
	time, err := ntp.Time("time.nist.gov")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	const layout = "3:04:05 PM (MST) on Monday, January _2, 2006"
	fmt.Println("Current Local Time:")
	fmt.Println(time.Local().Format(layout))
}
