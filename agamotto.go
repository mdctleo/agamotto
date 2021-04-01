package main

import (
	"fmt"
	"os"
	"time"
)

// agamotto project start -> writes to file
// agamotto project stop -> writes to file
// agamotto startDate endDate -> displays all projects worked between the two dates
// agamotto project date -> displays total hour worked on the project on that day
// agamotto project -> displays total hour worked on the project
// agamotto date -> displays all projects worked on that date
// agamotto project startDate endDate -> displays total hour worked on project between two dates

// date -> 20/11/2020
// date -> 11/20/2020
// date -> January 10, 2020
// date -> Jan 10, 2020

//
func main() {
	fmt.Println("Hello, 世界")
	args := os.Args[1:]

	fmt.Println(args)

	switch len(args) {
	case 1:
		fmt.Println("1 argument provided")
		fmt.Println(args[0])
		start := time.Now()
		fmt.Println(start)
	case 2:
		fmt.Println("2 arguments provided")
	case 3:
		fmt.Println("3 arguments provided")
	default:
		fmt.Println("print help info")
	}
}

func isDate(date string) {

}


func parseDate(date string) {

}

func write(projectName string) {

}

func read(projectName string, startDate string, endDate string) {

}

func displayHelp() {
	// display help menu when invalid command
}
