package folder1

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func F1() {
	file, err := os.Open("test_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close() // this line closes the file before exiting the program

	scanner := bufio.NewScanner(file) // create a new Scanner for the file
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // the Text() function converts the scanned bytes to a string
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

const chunkSize = 15

func F2() {
	file, err := os.Open("test_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	buf := make([]byte, chunkSize) // create a slice of bytes buffer with
	// the previously defined chunk size

	for {
		readTotal, err := file.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break // after reading the last chunk, break the loop
			}
			log.Fatal(err)
		}
		fmt.Println(string(buf[:readTotal]))
	}
}

func F3() {

	var fileName string

	fmt.Scan(&fileName)
	if fileName != "" {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}
}

//operations with maps
func Operations() {
	elements := make(map[string]string, 3)

	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	fmt.Println("element length is:", len(elements))

	for key := range elements {
		fmt.Println(key)
	}

	movieRatings := map[string]int{
		"The Matrix":          88,
		"Speed":               94,
		"The Matrix Reloaded": 73,
		"John Wick":           86,
	}

	fmt.Println("key - val")
	// Option #1 - create the 'val' variable to print the values of the map
	for key, val := range movieRatings {
		fmt.Println(key, val)
	}
}
