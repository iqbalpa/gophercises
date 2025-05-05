package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"main/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	csvVar = "csv"
	limitVar = "limit"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	args := os.Args[1:]
	kv := utils.ArgParser(args)
	
	limit := 30
	if kv[limitVar] != "" {
		limit, _ = strconv.Atoi(kv[limitVar])
	}
	csvPath := kv[csvVar]
	if csvPath == "" {
		csvPath = "./problems.csv"
	}

	counter := 0
	csvRecords := readCsvFile(csvPath)

	// init timer
	timer := time.NewTimer(time.Duration(limit) * time.Second)

	loop: for i, rec := range csvRecords {
		select {
			case <- timer.C:
				break loop
			default:
				q, ca := rec[0], rec[1]
				fmt.Printf("Problem #%d: %s = ", (i+1), q)
				a, _ := reader.ReadString('\n')
				a = strings.Trim(a, "\n")

				cai, _ := strconv.Atoi(ca)
				ai, _ := strconv.Atoi(a)
				if cai == ai {
					counter += 1
				} 
		}
	}

	fmt.Printf("You scored %d out of %d\n", counter, len(csvRecords))
}


// helper fucntion
func readCsvFile(filePath string) [][]string {
	if filePath == "" {
		filePath = "./problems.csv"
	}
	f, err := os.Open(filePath)
	if err != nil {
			log.Fatal("Unable to read input file " + filePath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
			log.Fatal("Unable to parse file as CSV for " + filePath, err)
	}
	return records
}