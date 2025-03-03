package proposal

import "tumelo_task/pkg/mockclient"

// CreateProposalsIndex  gets proposals for every general meeting it's given and then indexes them
// it indexes with: meetingID|proposalText|identifier => proposal ID
func CreateProposalsIndex(genMeetingIndex map[string]string) map[string]string {
	proposalsIndex := make(map[string]string)
	for _, genMeetingId := range genMeetingIndex {
		// TODO: the much more efficient way to do this would be concurrently send all requests,
		// then after all requests return run the loop again to aggregate all the data into the index
		proposalsForMeeting, _ := mockclient.GetProposalsForGeneralMeeting(genMeetingId)

		addProposalsToIndex(proposalsForMeeting, &proposalsIndex)
	}

	return proposalsIndex
}

func addProposalsToIndex(proposals []mockclient.Proposal, proposalsIndex *map[string]string) {
	for _, proposal := range proposals {
		key := generateProposalIndexKey(proposal)
		(*proposalsIndex)[key] = proposal.ID
	}
}

func generateProposalIndexKey(proposal mockclient.Proposal) string {
	return proposal.GeneralMeetingID + "|" + proposal.Text + "|" + proposal.Identifier
}