package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

func main() {
	content, err := ioutil.ReadFile("problems.csv")
	if err != nil {
		fmt.Println(err)
	}
	r := csv.NewReader(strings.NewReader(string(content)))
	questions, err := r.ReadAll()
	for _, question := range questions {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Problem #1: " + question[0] + " = ")
		answer, _ := reader.ReadString('\n')
		fmt.Println(answer)
	}
}