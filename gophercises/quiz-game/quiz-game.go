package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := flag.String("csv", "problems.csv", "file name to read")
	timeLimit := flag.Int("limit", 30, "time limit")
	flag.Parse()

	file, err := os.Open(*fileName)
	if err != nil {
		exit("cannot open the csv file")
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		exit("error while reading csv file")
	}

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	counter := 0

problemLoop:
	for _, record := range records {
		question := record[0]
		answer := record[1]

		fmt.Printf("%s: ", question)
		answerCh := make(chan string)

		go func() {
			var respond string
			if _, err := fmt.Scanf("%s\n", &respond); err != nil {
				exit("error while scan f")
			}
			answerCh <- respond
		}()

		select {
		case <-timer.C:
			break problemLoop
		case respond := <-answerCh:
			if respond == answer {
				counter++
			}
		}
	}

	fmt.Printf("correct answers %d from %d\n", counter, len(records))
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
