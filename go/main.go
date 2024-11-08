package main

import (
	"encoding/csv"
	"strconv"
	"strings"
	"syscall/js"
)

func processCsv(this js.Value, args []js.Value) interface{} {
	if len(args) < 1 {
		return 0
	}

	csvContent := args[0].String()
	reader := csv.NewReader(strings.NewReader(csvContent))
	
	var total float64
	
	records, err := reader.ReadAll()
	if err != nil {
		return 0
	}

	for _, record := range records {
		if len(record) > 0 {
			if price, err := strconv.ParseFloat(record[0], 64); err == nil {
				total += price
			}
		}
	}

	return total
}

func main() {
	c := make(chan struct{}, 0)
	js.Global().Set("processCsv", js.FuncOf(processCsv))
	<-c
}