package mockserver

import (
	"tumelo_task/pkg/mockclient"
)

// OrgID to general meetings
var generalMeetingsData = map[string][]mockclient.GeneralMeeting{
	"fbf713f4-d13b-4dbf-8e6c-33b494bfe519": {
		mockclient.GeneralMeeting{
			ID: "4a06d20c-c81c-430d-a1a8-c74e5ae323b0",
			OrganisationID: "fbf713f4-d13b-4dbf-8e6c-33b494bfe519",
			Date: "21/07/2023",
		},
		mockclient.GeneralMeeting{
			ID: "7ab1edb2-0176-4f2e-a2df-ffeb946a96ee",
			OrganisationID: "fbf713f4-d13b-4dbf-8e6c-33b494bfe519",
			Date: "21/07/2023",
		},
		mockclient.GeneralMeeting{
			ID: "30cf6652-e9a5-4553-a108-84751131df69",
			OrganisationID: "7193d18c-a18a-465c-987f-62fd4f0b30ad",
			Date: "21/07/2023",
		},
	},
}

// general_meeting_id to proposals
var proposalsData = map[string][]mockclient.Proposal{
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
	"30cf6652-e9a5-4553-a108-84751131df69": {
		{
			ID: "United-proposal-test-ID",
			GeneralMeetingID: "30cf6652-e9a5-4553-a108-84751131df69",
			Text: "Elect Michael Lewis",
			Identifier: "10",
		},
	},
}