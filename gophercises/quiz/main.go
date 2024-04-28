package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func loadQuiz() (io.Reader, error) {
	file := os.Getenv("QUESTIONS_PATH")
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %v", err)
	}
	return f, nil
}

func validator(userAnswer, correctAnswer int) int {
	if userAnswer == correctAnswer {
		return 1
	}
	return 0
}

func playQuiz(question []string) int {
	var answer int
	fmt.Printf("%s = ", question[0])
	
	_, err := fmt.Scan(&answer)
	if err != nil {
		log.Fatalf("failed to read user input: %v", err)
	}

	res, err := strconv.Atoi(question[1])
	if err != nil {
		log.Fatalf("failed to convert user answer to int: %v", err)
	}

	return validator(answer, res)
}

func parseQuiz(questions io.Reader) (int, int) {
	var totalScore, totalQuestions int
	r := csv.NewReader(questions)
	for {
		question, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed to read question: %v", err)
			continue
		}
		totalQuestions += 1
		totalScore += playQuiz(question)
	}
	return totalScore, totalQuestions
}

func main() {
	quiz, _ := loadQuiz()
	result, questions := parseQuiz(quiz)
	fmt.Println("Result:", result, "/", questions)
}
