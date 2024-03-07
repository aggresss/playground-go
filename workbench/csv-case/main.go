package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	f, err := os.Create("users.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	for i := 0; i < 1000; i++ {
		if err := w.Write([]string{
			strconv.Itoa(i),
			strconv.Itoa(rand.Intn(100)),
		}); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
