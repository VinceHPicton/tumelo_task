package recommendation

// FindMatchingRecommendations returns proposalID => Recommendation (struct)
func FindMatchingRecommendations(recommendations *[]Recommendation, genMeetingIndex map[string]string, proposalsIndex map[string]string) map[string]Recommendation {
	matchedRecommendations := make(map[string]Recommendation)
	for _, recommendation := range *recommendations {
		// TODO: could split this up into 2 further functions, would probably just make the code more confusing though
		meetingKey := generateMeetingIndexKey(recommendation)
		meetingID, meetingFound := genMeetingIndex[meetingKey]
		if !meetingFound {
			continue
		}
		// fmt.Println("Found meetingID: ", meetingID)

		proposalKey := generateProposalKey(meetingID, recommendation)
		// fmt.Println("proposalKey: ", proposalKey)
		proposalID, proposalFound := proposalsIndex[proposalKey]
		if !proposalFound {
			continue
		}
		// fmt.Println("Found proposalID: ", proposalID)

		matchedRecommendations[proposalID] = recommendation
	}

	return matchedRecommendations
}

func generateMeetingIndexKey(recommendation Recommendation) string {
	return recommendation.OrganisationID + "|" + recommendation.MeetingDate
}

func generateProposalKey(meetingID string, recommendation Recommendation) string {
	return meetingID + "|" + recommendation.ProposalText + "|" + recommendation.SequenceID
}