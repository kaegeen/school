package main

import "fmt"

func main() {
	var startYear, studyDuration int

	fmt.Println("Study Completion Year Calculator")
	fmt.Println("===============================")

	// Input the start year
	fmt.Print("Enter the year you started studying: ")
	fmt.Scan(&startYear)

	// Input the duration of study
	fmt.Print("Enter the duration of your study in years: ")
	fmt.Scan(&studyDuration)

	// Calculate the year of completion
	completionYear := startYear + studyDuration
	fmt.Printf("If you started studying in %d for %d years, you would finish in the year %d.\n", startYear, studyDuration, completionYear)
}
