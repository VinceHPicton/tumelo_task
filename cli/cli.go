package cli

import (
	"fmt"
	"os"
	"slices"
	"tumelo_task/recommendation"
)

const (
	listInvalidDataAndStopOption = "1. List invalid data and stop"
	runAFixOption                = "2. Run a fix on the data and retry"
)

// TODO: this really isn't great, multiple soruces of truth for allowed numbers etc, needs refactor if I had time
var allowedChoices = []int{1, 2}

var options = []string{
	listInvalidDataAndStopOption,
	runAFixOption,
}

func Start() string {
	var fileName string
	fmt.Print("Enter a CSV file path: ")
	fmt.Scanln(&fileName)

	return fileName
}

func ListOptionsForInvalidData() int {
	fmt.Println("Invalid data was detected, select from the following options")

	for _, option := range options {
		fmt.Printf("%s\n", option)
	}

	var choice int
	fmt.Print("Enter the number of your choice: ")
	_, err := fmt.Scanln(&choice)

	choiceAllowed := slices.Contains(allowedChoices, choice)

	if err != nil || !choiceAllowed {
		// TODO: I did originally put a retry here - where you could try again to select your number, but 
		// This was a little buggy and so for time reasons I just exit 
		fmt.Println("\n---Invalid choice. Please try again.---")
		os.Exit(1)
	}

	return choice
}

func ShouldAttemptFix() bool {

	chosenOption := ListOptionsForInvalidData()
	chosenOptionStr := options[chosenOption-1]

	switch chosenOptionStr {

	case listInvalidDataAndStopOption:
		return false
	case runAFixOption:
		return true
	default:
		return true
	}
}

func ListInvalidRecommendations(invalidRecommendations []recommendation.InvalidRecommendation) {
	fmt.Println("Invalid data:")
	for _, invalidRec := range invalidRecommendations {
		fmt.Printf("csv line: %d, reason: %s\n", invalidRec.OriginalIndex, invalidRec.Reason)
	}
}

func HandleInvalidDataScenario(invalidRecommendations []recommendation.InvalidRecommendation, recommendationsPtr *[]recommendation.Recommendation) (success bool, newInvalidRecommendations []recommendation.InvalidRecommendation){

	if !ShouldAttemptFix() {
		ListInvalidRecommendations(invalidRecommendations)
		os.Exit(1)
	}

	successfullyFixed, newInvalidRecs := HandleFixDataAttempt(recommendationsPtr)

	if !successfullyFixed {
		return false, newInvalidRecs
	}

	return true, newInvalidRecs
}

func HandleFixDataAttempt(recommendationsPtr *[]recommendation.Recommendation) (success bool, newInvalidRecommendations []recommendation.InvalidRecommendation) {

	recommendation.CleanAllRecommendations(recommendationsPtr)
	newInvalidRecs := recommendation.FindInvalidRecommendations(recommendationsPtr)

	if len(newInvalidRecs) == 0 {
		return true, newInvalidRecs
	}

	return false, newInvalidRecs
}
