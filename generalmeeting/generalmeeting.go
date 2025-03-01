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

func CreateMeetingIndex(orgNameToIDMap map[string]string) map[string]string {
	// Index meetings
	// TODO: Issue: if 2 meetings are on the same day, one will be lost, solution make it a []string
	// OrgID + Date = MeetingID
	genMeetingIndex := make(map[string]string)
	for _, orgID := range orgNameToIDMap {
		generalMeetingsForOrg := GetGeneralMeetingForOrganisation(orgID)
		
		for _, meeting := range generalMeetingsForOrg {
			key := meeting.OrganisationID + "|" + meeting.Date
			genMeetingIndex[key] = meeting.ID
		}
	}

	return genMeetingIndex
}


func GetGeneralMeetingForOrganisation(organisationID string) []GeneralMeeting{
	meetings := data[organisationID]

	return meetings
}