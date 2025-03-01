package main

import (
	"fmt"
	"log"
	"os"
	"tumelo_task/cli"
	"tumelo_task/generalmeeting"
	"tumelo_task/pkg/csv_reader"
	"tumelo_task/pkg/get_organisations"
	"tumelo_task/pkg/submit_results"
	"tumelo_task/proposal"
	"tumelo_task/recommendation"
)

func main() {

	// Kickoff CLI
	// csvFilePath := cli.Start()
	// csvFilePath := "./ExampleRecommendationsOriginal.csv"
	csvFilePath := "./ExampleRecommendationsClean.csv"
	// csvFilePath := "./ExampleRecommendationsSmall.csv"
	// csvFilePath := "./OneRecommendation.csv"

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
		invalidDataFixed, newInvalidRecs := cli.HandleInvalidDataScenario(invalidRecommendations, &recommendations)

		if !invalidDataFixed {
			fmt.Println("----Data cleaning did not remove all invalid data, listing invalid data and stopping----")
			cli.ListInvalidRecommendations(newInvalidRecs)
			os.Exit(1)
		}
	}

	// At this point we should have no invalid CSV data

	genMeetingIndex := generalmeeting.CreateMeetingIndex(orgNameToIDMap)

	proposalsIndex := proposal.CreateProposalsIndex(genMeetingIndex)

	matchedRecommendations := recommendation.FindMatchingRecommendations(&recommendations, genMeetingIndex, proposalsIndex)

	// println("genMeetingIndex:")
	// for k , v := range genMeetingIndex {
	// 	fmt.Println(k ,v)
	// }

	// println("ProposalsIndex:")
	// for k , v := range proposalsIndex {
	// 	fmt.Println(k ,v)
	// }

	fmt.Println("matched recommendationss:")
	for proposalID, rec := range matchedRecommendations {
		fmt.Println("ProposalID: ", proposalID, "Recommendation: ", rec.Recommendation)
	}

	errorList := submit_results.SubmitRecommendations(matchedRecommendations)

	fmt.Println("Process complete, any failed submissions listed below:")
	for _, err := range errorList {
		fmt.Println(err.Error())
	}

}
