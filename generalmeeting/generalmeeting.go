package generalmeeting

import "tumelo_task/pkg/mockclient"

// func GetGeneralMeetingForOrganisation(organisationID string) []GeneralMeeting{
// 	// TODO: this should be an actual API call in the real app but for the same reason as get_organisations I've simplified to doing this
// 	meetings := generalMeetingsData[organisationID]

// 	return meetings
// }

// CreateMeetingIndex gets general meetings for every organisation it's given and then indexes them
// with key: OrganisationID|MeetingDate => meeting ID
func CreateMeetingIndex(orgNameToIDMap map[string]string) map[string]string {
	// TODO: Issue: if 2 or more meetings are on the same day, all but one will be lost, solution: make it a []string and loop through that later in the process
	// This is a bit of a edge case so i've acknowledged it and moved on due to time.
	genMeetingIndex := make(map[string]string)
	for _, orgID := range orgNameToIDMap {
		// TODO: the much more efficient way to do this would be concurrently send all requests,
		// then after all requests return run the loop again to aggregate all the data into the index
		generalMeetingsForOrg, _ := mockclient.GetGeneralMeetingForOrganisation(orgID)

		//TODO: we should be collecting these errors and possibly logging them, not time right now
		// It getting an err for  one of these calls will usually mean there were just no meetings for that org, not a failure
		
		addMeetingsToIndex(generalMeetingsForOrg, &genMeetingIndex)
	}

	return genMeetingIndex
}

func addMeetingsToIndex(meetings []mockclient.GeneralMeeting, meetingIndex *map[string]string) {
	for _, meeting := range meetings {
		key := generateMeetingIndexKey(meeting)
		(*meetingIndex)[key] = meeting.ID
	}
}

func generateMeetingIndexKey(meeting mockclient.GeneralMeeting) string {
	return meeting.OrganisationID + "|" + meeting.Date
}