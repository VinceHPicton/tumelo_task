package recommendation

type Recommendation struct {
	OrganisationID string
	Name string
	MeetingDate string
	SequenceID string
	ProposalText string
	Recommendation string
}

const (
	For = "For"
	Against = "Against"
	Abstain = "Abstain"

	validDateFormat = "02/01/2006"

	expectedNumberOfColumnsPerCSVLine = 5
)

func AddOrganisationIDsToRecommendations(recommendations *[]Recommendation, orgNameToIDMap *map[string]string) {
	// Why are we taking pointers here? Performance: input may be very large and there's no need to copy it for this by passing by value

	// TODO: unit test this
	for i, recommendation := range *recommendations {
		id, ok := (*orgNameToIDMap)[recommendation.Name]
		if ok {
			(*recommendations)[i].OrganisationID = id
		}
	}
}