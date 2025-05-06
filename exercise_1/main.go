package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	csv := flag.String("csv", "problems.csv", "File path to problem (default: problems.csv)")
	limit := flag.Int("limit", 30, "Quiz time limit (default: 30s)")
	flag.Parse()
	
	counter := 0
	problems := readCsvFile(*csv)
	timer := time.NewTimer(time.Duration(*limit) * time.Second)
	
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", (i+1), p.Q)
		
		answerChan := make(chan string)
		go func(){
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerChan <- answer
		}()

		select {
			case <- timer.C:
				fmt.Println("\nTimed out!")
				return
			case answer := <- answerChan:
				if answer == p.A {
					counter++
					fmt.Println("Correct!")
				}
		}
	}

	fmt.Printf("You scored %d out of %d\n", counter, len(problems))
}

func readCsvFile(fpath string) []Problem {
	f, err := os.Open(fpath)
	if err != nil {
		log.Fatal("Unable to read input file " + fpath, err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for " + fpath, err)
	}
	ret := make([]Problem, len(records))
	for i,rec := range records {
		ret[i] = Problem{
			Q: rec[0],
			A: rec[1],
		}
	}
	return ret
}

type Problem struct {
	Q string
	A string
}