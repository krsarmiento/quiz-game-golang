package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

func main() {
	csvFlag := flag.String("csv", "problems.csv", "the csv file with the answers")
	limitFlag := flag.Int("limit", 30, "time limit for the quiz")
	shuffleFlag := flag.Bool("shuffle", false, "if the questions should be asked in random order") 

	flag.Parse()

	fmt.Printf("Time limit for this quiz: %d\n", *limitFlag)

	content, err := ioutil.ReadFile(*csvFlag)
	if err != nil {
		fmt.Println(err)
	}
	r := csv.NewReader(strings.NewReader(string(content)))
	questions, err := r.ReadAll()

	if *shuffleFlag {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(questions), func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
	}

	correctAnswers := 0
	timer := time.NewTimer(time.Duration(*limitFlag) * time.Second)
	go func() {
		<-timer.C
		fmt.Printf("\nYou scored %d out of %d.\n", correctAnswers, len(questions))
		os.Exit(0)
	}()
	
	for index, question := range questions {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Problem #%d: %s = ", index+1, question[0])
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSuffix(answer, "\n")
		if answer == question[1] {
			correctAnswers++
		}
	}
	fmt.Printf("You scored %d out of %d.\n", correctAnswers, len(questions))
}
