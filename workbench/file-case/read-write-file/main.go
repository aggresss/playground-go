package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("REDEME.md")

	b1 := make([]byte, 371)
	f.Read(b1)

	for i := 0; i < len(b1); i++ {
		if i%16 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("0x%02X, ", b1[i])
	}

	err := os.WriteFile("README.copy.md", b1, 0644)
	if err != nil {
		fmt.Println("write error", err)
	}
}
