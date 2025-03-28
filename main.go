package main

import (
	"fmt"
	"log"
	"os"
	"tumelo_task/cli"
	"tumelo_task/generalmeeting"
	"tumelo_task/pkg/csvreader"
	"tumelo_task/pkg/mockclient"
	"tumelo_task/pkg/mockserver"
	"tumelo_task/pkg/submitresults"
	"tumelo_task/proposal"
	"tumelo_task/recommendation"
)

func main() {

	fmt.Println("Starting Mock HTTP server")
	go mockserver.Start()

	// Kickoff CLI
	csvFilePath := cli.Start()

	// Uncomment these lines and comment above for easier testing of my submission (you don't have to type out the filepath)
	// csvFilePath := "./ExampleRecommendationsOriginal.csv"
	// csvFilePath := "./ExampleRecommendationsClean.csv"

	// Read CSV data step
	recommendationsData, err := csvreader.ReadIgnoringHeader(csvFilePath)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Parse CSV data into structs step
	recommendations, err := recommendation.ParseCSVDataToRecommendations(&recommendationsData)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Get organisations +ids step
	orgNameToIDMap, err := mockclient.GetOrganisations()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Add IDs to orgs step
	recommendation.AddOrganisationIDsToRecommendations(&recommendations, &orgNameToIDMap)

	// Validate step
	invalidRecommendations := recommendation.FindInvalidRecommendations(&recommendations)

	// Use a CLI to give user options on how to process invalid data
	// TODO: this CLI has been refactored a little, so it's not hideous, but would need a lot more cleaning up before production 
	// I don't love it the way it is
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

	fmt.Println("\n---- The following matched recommendations will be submitted ----")
	for proposalID, rec := range matchedRecommendations {
		fmt.Println("ProposalID: ", proposalID, "Recommendation: ", rec.Recommendation)
	}

	errorList := submitresults.SubmitRecommendations(matchedRecommendations)

	fmt.Println("\n---- Process complete, any failed submissions listed below ----")
	for _, err := range errorList {
		fmt.Println(err.Error())
	}

	os.Exit(0)
}
