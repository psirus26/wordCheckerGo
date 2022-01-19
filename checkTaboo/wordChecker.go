package checkTaboo

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func ReadFileWithTabooWords() map[string]bool {

	var fileNameTaboo string
	var tabooWords = make(map[string]bool)

	fmt.Println("Enter the name of the file containing the taboo words dictionary")
	fmt.Scan(&fileNameTaboo)
	if fileNameTaboo != "" {
		file, err := os.Open(fileNameTaboo)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			tabooWords[strings.ToLower(scanner.Text())] = true
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

	}

	return tabooWords
}

func checkIfWordIsTaboo(currentWord string, tabooWords map[string]bool) string {

	var result string

	if tabooWords[strings.ToLower(currentWord)] == true {
		result = strings.Repeat("*", len(currentWord))
	} else {
		result = currentWord
	}

	return result
}

func checkIfSentenceHasTabooWords(currentSentence string, tabooWords map[string]bool) string {

	var result string

	var splitSentenceWords = strings.Split(currentSentence, " ")
	var splitSentence = make([]string, len(splitSentenceWords))
	splitSentence = splitSentenceWords

	for _, word := range splitSentence {
		result += checkIfWordIsTaboo(word, tabooWords)
	}

	return result
}

func RunCheckTaboo() {

	var tabooSentence string
	var tabooWords = make(map[string]bool)
	tabooWords = ReadFileWithTabooWords()

	fmt.Println("Enter a sentence to see if it contains a taboo word. Enter 'exit' to quit")
	for true {
		fmt.Scan(&tabooSentence)
		if tabooSentence != "exit" {
			fmt.Println(checkIfSentenceHasTabooWords(tabooSentence, tabooWords))
		} else {
			fmt.Println("Bye!")
			break
		}
	}

}
