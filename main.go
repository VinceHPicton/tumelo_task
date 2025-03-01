package main

import (
	"fmt"
	"log"
	"os"
	"tumelo_task/cli"
	"tumelo_task/generalmeeting"
	"tumelo_task/pkg/csv_reader"
	"tumelo_task/pkg/get_organisations"
	"tumelo_task/proposal"
	"tumelo_task/recommendation"
)

func main() {

	// Kickoff CLI
	// csvFilePath := cli.Start()
	// csvFilePath := "./ExampleRecommendationsOriginal.csv"
	// csvFilePath := "./ExampleRecommendationsClean.csv"
	csvFilePath := "./ExampleRecommendationsSmall.csv"
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

	fmt.Println(recommendations)
	for _, v := range recommendations {
		fmt.Println(v)
	}

	// At this point we should have no invalid data:
	fmt.Println("At this point we should have no invalid data")



	genMeetingIndex := generalmeeting.CreateMeetingIndex(orgNameToIDMap)

	proposalsIndex := proposal.CreateProposalsIndex(genMeetingIndex)

	// fmt.Println(genMeetingIndex)

	// fmt.Println(proposalsIndex)

	matchedRecommendations := make(map[string]recommendation.Recommendation)
	for _, rec := range recommendations {
		meetingKey := rec.OrganisationID + "|" + rec.MeetingDate
		meetingID, found := genMeetingIndex[meetingKey]
		if !found {
			continue
		}
		fmt.Println("Found meetingID: ", meetingID)

		proposalKey := meetingID + "|" + rec.ProposalText + "|" + rec.SequenceID
		fmt.Println("proposalKey: ", proposalKey)
		proposalID, found := proposalsIndex[proposalKey]
		if !found {
			continue
		}
		fmt.Println("Found proposalID: ", proposalID)

		matchedRecommendations[proposalID] = rec
	}

	println("genMeetingIndex:")
	for k , v := range genMeetingIndex {
		fmt.Println(k ,v)
	}

	println("ProposalsIndex:")
	for k , v := range proposalsIndex {
		fmt.Println(k ,v)
	}

	fmt.Println("matched recs:")
	for k, v := range matchedRecommendations {
		fmt.Println(k, v)
	}

}
