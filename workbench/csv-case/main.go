package main

import (
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	RFC3339Milli = "2006-01-02T15:04:05.000Z07:00"
)

func main() {

	f, err := os.Create("time2.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write([]string{
		"index",
		"score",
	}); err != nil {
		log.Fatalln("error writing field record to file", err)
	}

	ticker := time.NewTicker(time.Millisecond * 1)
	count := 0

	for now := range ticker.C {
		if err := w.Write([]string{
			// strconv.FormatInt(now.UnixMilli(), 10) // grafana NOT compatibility
			// now.Format(time.RFC3339) // grafana NOT compatibility
			// strconv.Itoa(count) // grafana compatibility
			// now.Format(RFC3339Milli) // grafana compatibility
			now.Format(RFC3339Milli),
			strconv.Itoa(rand.Intn(100)),
		}); err != nil {
			log.Fatalln("error writing record to file", now, err)
		}
		count++
		if count > 1000 {
			break
		}
	}
}
