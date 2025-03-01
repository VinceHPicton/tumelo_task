package proposal

type Proposal struct {
	ID string
	GeneralMeetingID string
	Text string
	Identifier string
}

var proposalsData = map[string][]Proposal{
	"4a06d20c-c81c-430d-a1a8-c74e5ae323b0": {
		{
			ID: "4f7151ef-ea8b-4e0c-94d3-2e7b33a2e175",
			GeneralMeetingID: "4a06d20c-c81c-430d-a1a8-c74e5ae323b0",
			Text: "Elect Dominic J. Caruso",
			Identifier: "1b.",
		},
		{
			ID: "2ee01feb-f904-4f1f-b808-aabc6551dda8",
			GeneralMeetingID: "4a06d20c-c81c-430d-a1a8-c74e5ae323b0",
			Text: "Elect W. Roy Dunbar",
			Identifier: "1c.",
		},
	},
	"7ab1edb2-0176-4f2e-a2df-ffeb946a96ee": {
		{
			ID: "623a8bcb-d3a8-4f41-999b-e7497173dd02",
			GeneralMeetingID: "7ab1edb2-0176-4f2e-a2df-ffeb946a96ee",
			Text: "Elect Maria Martinez",
			Identifier: "1h.",
		},
		{
			ID: "1c39aaec-b68c-4228-b747-669b591186cf",
			GeneralMeetingID: "7ab1edb2-0176-4f2e-a2df-ffeb946a96ee",
			Text: "Elect Susan R. Salka",
			Identifier: "1i.",
		},
	},
}

func GetProposalsForGenMeeting(generalMeetingID string) []Proposal {
	proposals := proposalsData[generalMeetingID]

	return proposals
}

// CreateProposalsIndex indexes with: meetingID|proposalText|identifier => proposal ID
func CreateProposalsIndex(genMeetingIndex map[string]string) map[string]string {
	proposalsIndex := make(map[string]string)
	for _, genMeetingId := range genMeetingIndex {
		// TODO: the much more efficient way to do this would be concurrently send all requests,
		// then after all requests return run the loop to aggregate all the data
		proposalsForMeeting := GetProposalsForGenMeeting(genMeetingId)

		addProposalsToIndex(proposalsForMeeting, &proposalsIndex)
	}

	return proposalsIndex
}

func addProposalsToIndex(proposals []Proposal, proposalsIndex *map[string]string) {
	for _, proposal := range proposals {
		key := generateProposalIndexKey(proposal)
		(*proposalsIndex)[key] = proposal.ID
	}
}

func generateProposalIndexKey(proposal Proposal) string {
	return proposal.GeneralMeetingID + "|" + proposal.Text + "|" + proposal.Identifier
}