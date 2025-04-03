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

func main2() {
	go startMockServer()
	csvFilePath := getCSVFilePath()

	recommendations := parseCSV(csvFilePath)

	orgNameToIDMap := fetchOrganisations()

	recommendation.AddOrganisationIDsToRecommendations(&recommendations, &orgNameToIDMap)

	handleInvalidRecommendations(&recommendations)

	genMeetingIndex := generalmeeting.CreateMeetingIndex(orgNameToIDMap)
	proposalsIndex := proposal.CreateProposalsIndex(genMeetingIndex)

	matchedRecommendations := recommendation.FindMatchingRecommendations(&recommendations, genMeetingIndex, proposalsIndex)

	submitResults(matchedRecommendations)
}

func startMockServer() {
	fmt.Println("Starting Mock HTTP server")
	mockserver.Start()
}

func getCSVFilePath() string {
	return cli.Start()
}

func parseCSV(csvFilePath string) []recommendation.Recommendation {
	data, err := csvreader.ReadIgnoringHeader(csvFilePath)
	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	recommendations, err := recommendation.ParseCSVDataToRecommendations(&data)
	if err != nil {
		log.Fatalf("Error parsing CSV data: %v", err)
	}

	return recommendations
}

func fetchOrganisations() map[string]string {
	orgs, err := mockclient.GetOrganisations()
	if err != nil {
		log.Fatalf("Error fetching organisations: %v", err)
	}
	return orgs
}

func handleInvalidRecommendations(recommendations *[]recommendation.Recommendation) {
	invalidRecommendations := recommendation.FindInvalidRecommendations(recommendations)
	if len(invalidRecommendations) == 0 {
		return
	}

	fixed, remainingInvalid := cli.HandleInvalidDataScenario(invalidRecommendations, recommendations)
	if !fixed {
		fmt.Println("----Data cleaning did not remove all invalid data, listing invalid data and stopping----")
		cli.ListInvalidRecommendations(remainingInvalid)
		os.Exit(1)
	}
}

func submitResults(matchedRecommendations map[string]recommendation.Recommendation) {
	fmt.Println("\n---- The following matched recommendations will be submitted ----")
	for proposalID, rec := range matchedRecommendations {
		fmt.Printf("ProposalID: %s, Recommendation: %s\n", proposalID, rec.Recommendation)
	}

	errors := submitresults.SubmitRecommendations(matchedRecommendations)

	if len(errors) > 0 {
		fmt.Println("\n---- Process complete, failed submissions listed below ----")
		for _, err := range errors {
			fmt.Println(err.Error())
		}
	}
}
