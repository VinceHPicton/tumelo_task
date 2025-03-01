package recommendation

// FindMatchingRecommendations returns map of proposalID => Recommendation by looking up correct meeting key using data on recommendation
// And then looking up the correct proposal for that meeting
func FindMatchingRecommendations(recommendations *[]Recommendation, genMeetingIndex map[string]string, proposalsIndex map[string]string) map[string]Recommendation {
	matchedRecommendations := make(map[string]Recommendation)
	for _, recommendation := range *recommendations {
		// TODO: could split this up into 2 further functions, would probably just make the code more confusing though
		meetingKey := generateMeetingIndexKey(recommendation)
		meetingID, meetingFound := genMeetingIndex[meetingKey]
		if !meetingFound {
			continue
		}

		proposalKey := generateProposalKey(meetingID, recommendation)
		proposalID, proposalFound := proposalsIndex[proposalKey]
		if !proposalFound {
			continue
		}

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