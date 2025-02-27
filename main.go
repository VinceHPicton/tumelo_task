package main

import (
	"fmt"
	"log"
	"os"
	"tumelo_task/cli"
	"tumelo_task/pkg/csv_reader"
	"tumelo_task/pkg/get_organisations"
	"tumelo_task/recommendation"
)

func main() {

	// Kickoff CLI
	//csvFilePath := cli.Start()
	csvFilePath := "./ExampleRecommendations.csv"

	// Read CSV data step
	recommendationsData, err := csv_reader.ReadIgnoringHeader(csvFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Parse CSV data into structs step
	recommendations, err := recommendation.ParseCSVDataToRecommendations(&recommendationsData)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Get organisations +ids step
	orgNameToIDMap, err := get_organisations.GetOrganisations()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Add IDs to orgs step
	recommendation.AddIDs(&recommendations, &orgNameToIDMap)

	// Validate and give user options step
	invalidRecommendations := recommendation.FindInvalidRecommendations(&recommendations)

	if len(invalidRecommendations) > 0 {
		HandleInvalidDataScenario(invalidRecommendations, &recommendations)
	}
}

func HandleInvalidDataScenario(invalidRecommendations []recommendation.InvalidRecommendation, recommendationsPtr *[]recommendation.Recommendation) {

	if cli.ShouldAttemptFix() {

		recommendation.CleanAllRecommendations(recommendationsPtr)
		newInvalidRecs := recommendation.FindInvalidRecommendations(recommendationsPtr)
		if len(newInvalidRecs) > 0 {
			fmt.Println("----Data cleaning did not remove all invalid data, listing invalid data and stopping----")
			cli.ListInvalidRecommendations(newInvalidRecs)
			os.Exit(1)
		} else {
			return
		}

	} else {
		cli.ListInvalidRecommendations(invalidRecommendations)
		os.Exit(1)
	}
}