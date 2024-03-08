package main

import (
	"fmt"
	"time"
)

// Formatting / Parsing

const (
	RFC3339Milli = "2006-01-02T15:04:05.999Z07:00"
)

func main() {
	fmt.Println(time.Now().Format(RFC3339Milli))
	fmt.Println(time.Now().UnixMilli())
}
