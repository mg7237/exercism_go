package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {

	csv := flag.String("csv", "problems.csv", "quiz csv file name ")
	// limit := flag.Int("limit", 30, "time limit for the quiz")
	flag.Parse()
	f := *csv
	// l := *limit

	qna, err := readCSV(f)
	if err != nil {
		fmt.Println("Invalid File:", f)
		os.Exit(2)
	}
	reader := bufio.NewReader(os.Stdin)
	count, qestion := 0, 1
	for i := range qna {
		fmt.Println("Question #", qestion, ": ", qna[i].q, "=")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input", f)
			os.Exit(3)
		}
		if text == qna[i].a+"\n" {
			count++
		}
		qestion++
	}
	fmt.Println("Your Score: ", count, " out of ", len(qna))
}

type quiz struct {
	q string
	a string
}

func readCSV(file string) (map[int]quiz, error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}
	qna := make(map[int]quiz)
	for i, line := range lines {
		qna[i] = quiz{line[0], line[1]}
	}

	return qna, nil
}
