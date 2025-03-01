package generalmeeting

type GeneralMeeting struct {
	ID string
	OrganisationID string
	Date string
}

// OrgID to general meeting
var data = map[string][]GeneralMeeting{
	"fbf713f4-d13b-4dbf-8e6c-33b494bfe519": {
		GeneralMeeting{
			ID: "4a06d20c-c81c-430d-a1a8-c74e5ae323b0",
			OrganisationID: "fbf713f4-d13b-4dbf-8e6c-33b494bfe519",
			Date: "21/07/2023",
		},
		GeneralMeeting{
			ID: "7ab1edb2-0176-4f2e-a2df-ffeb946a96ee",
			OrganisationID: "fbf713f4-d13b-4dbf-8e6c-33b494bfe519",
			Date: "21/07/2023",
		},
	},
}

func GetGeneralMeetingForOrganisation(organisationID string) []GeneralMeeting{
	meetings := data[organisationID]

	return meetings
}

// CreateMeetingIndex indexes with: OrganisationID|MeetingDate => meeting ID
func CreateMeetingIndex(orgNameToIDMap map[string]string) map[string]string {
	// TODO: Issue: if 2 or more meetings are on the same day, all but one will be lost, solution: make it a []string
	genMeetingIndex := make(map[string]string)
	for _, orgID := range orgNameToIDMap {
		// TODO: the much more efficient way to do this would be concurrently send all requests,
		// then after all requests return run the loop again to aggregate all the data into the index
		generalMeetingsForOrg := GetGeneralMeetingForOrganisation(orgID)
		
		addMeetingsToIndex(generalMeetingsForOrg, &genMeetingIndex)
	}

	return genMeetingIndex
}

func addMeetingsToIndex(meetings []GeneralMeeting, meetingIndex *map[string]string) {
	for _, meeting := range meetings {
		key := generateMeetingIndexKey(meeting)
		(*meetingIndex)[key] = meeting.ID
	}
}

func generateMeetingIndexKey(meeting GeneralMeeting) string {
	return meeting.OrganisationID + "|" + meeting.Date
}