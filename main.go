package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/eskibars/lmsensorsbeat/beater"
)

func main() {
	err := beat.Run("lmsensorsbeat", "", beater.New())
	if err != nil {
		os.Exit(1)
	}
}
