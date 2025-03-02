package recommendation

// FindMatchingRecommendations returns map of proposalID => Recommendation by looking up correct meeting key using data on recommendation
// And then looking up the correct proposal for that meeting
func FindMatchingRecommendations(recommendations *[]Recommendation, genMeetingIndex map[string]string, proposalsIndex map[string]string) map[string]Recommendation {
	matchedRecommendations := make(map[string]Recommendation)
	for _, recommendation := range *recommendations {
		// TODO: could split this loop up into 2 further functions, would probably just make the code more confusing though

		meetingKey := generateMeetingIndexKey(recommendation)
		// TODO: it's at this point where some data could be missed - imagine if 2 meetings happened on the same day.
		// Here what we'd need to do is have meetingIDAtRecommendationDate be a slice of strings, not just 1 string
		// We'd then loop that slice and for each item in it run the 2nd piece to collect that proposal
		// The previous code where we create the indexes would also need to change of course
		meetingIDAtRecommendationDate, meetingFound := genMeetingIndex[meetingKey]
		if !meetingFound {
			continue
		}

		proposalKey := generateProposalKey(meetingIDAtRecommendationDate, recommendation)
		proposalIDFromMeetingWithCorrectSequenceIDAndText, proposalFound := proposalsIndex[proposalKey]
		if !proposalFound {
			continue
		}

		matchedRecommendations[proposalIDFromMeetingWithCorrectSequenceIDAndText] = recommendation
	}

	return matchedRecommendations
}

func generateMeetingIndexKey(recommendation Recommendation) string {
	return recommendation.OrganisationID + "|" + recommendation.MeetingDate
}

func generateProposalKey(meetingID string, recommendation Recommendation) string {
	return meetingID + "|" + recommendation.ProposalText + "|" + recommendation.SequenceID
}