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
	// csvFilePath := "./ExampleRecommendationsOriginal.csv"
	// csvFilePath := "./ExampleRecommendationsClean.csv"
	csvFilePath := "./OneRecommendation.csv"

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
		invalidDataFixed := cli.HandleInvalidDataScenario(invalidRecommendations, &recommendations)

		if !invalidDataFixed {
			os.Exit(1)
		}
	}

	// At this point we should have no invalid data:
	fmt.Println("At this point we should have no invalid data")
	fmt.Println(recommendations)

}
