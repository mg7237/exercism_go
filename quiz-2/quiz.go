package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

var count, lenqa, question int = 0, 0, 1

func main() {

	csv := flag.String("csv", "problems.csv", "quiz csv file name ")
	limit := flag.Int("limit", 30, "time limit for the quiz")
	flag.Parse()
	f := *csv
	l := *limit

	qna, err := readCSV(f)
	if err != nil {
		fmt.Println("Invalid File:", f)
		os.Exit(2)
	}
	go timer(l)
	reader := bufio.NewReader(os.Stdin)
	for i := range qna {
		fmt.Println("Question #", question, ": ", qna[i].q, "=")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input", f)
			os.Exit(3)
		}
		if text == qna[i].a+"\n" {
			count++
		}
		question++
	}
	fmt.Println("Your Score: ", count, " out of ", lenqa)
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
	lenqa = len(qna)
	return qna, nil
}

func timer(seconds int) {
	s, _ := time.ParseDuration(strconv.Itoa(seconds) + "s")
	t := time.NewTimer(s)
	<-t.C
	if question != lenqa {
		fmt.Println("\nYour Score: ", count, " out of ", lenqa)
	}
	os.Exit(0)
}
