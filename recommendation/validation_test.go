package recommendation

import "testing"

func Test_Validate(t *testing.T) {

	initialRec := Recommendation{
		OrganisationID: "432aaae7-500b-492c-bc87-360437b37354",
		MeetingDate: "21/07/2015",
		Recommendation: For,
	}

	// TODOS: add more exhaustive date tests
	// Potentially add some other tcs against Recommendation string but these would add little value
	testCases := map[string]struct {
		given Recommendation
		expectErr bool
	}{
		"happy path recommend For": {
			given: Recommendation{
				OrganisationID: initialRec.OrganisationID,
				MeetingDate: initialRec.MeetingDate,
				Recommendation: For,
			},
			expectErr: false,
		},
		"happy path recommend Abstain": {
			given: Recommendation{
				OrganisationID: initialRec.OrganisationID,
				MeetingDate: initialRec.MeetingDate,
				Recommendation: Abstain,
			},
			expectErr: false,
		},
		"happy path recommend Against": {
			given: Recommendation{
				OrganisationID: initialRec.OrganisationID,
				MeetingDate: initialRec.MeetingDate,
				Recommendation: Against,
			},
			expectErr: false,
		},
		"recommendation error": {
			given: Recommendation{
				OrganisationID: initialRec.OrganisationID,
				MeetingDate: initialRec.MeetingDate,
				Recommendation: "anything",
			},
			expectErr: true,
		},
		"date error 1 digit month": {
			given: Recommendation{
				OrganisationID: initialRec.OrganisationID,
				MeetingDate: "21/7/2015",
				Recommendation: initialRec.Recommendation,
			},
			expectErr: true,
		},
		"date error 1 digit day": {
			given: Recommendation{
				OrganisationID: initialRec.OrganisationID,
				MeetingDate: "1/07/2015",
				Recommendation: initialRec.Recommendation,
			},
			expectErr: true,
		},
		"date error day 00": {
			given: Recommendation{
				OrganisationID: initialRec.OrganisationID,
				MeetingDate: "00/07/2015",
				Recommendation: initialRec.Recommendation,
			},
			expectErr: true,
		},
		"date error day invalid": {
			given: Recommendation{
				OrganisationID: initialRec.OrganisationID,
				MeetingDate: "30/02/2015",
				Recommendation: initialRec.Recommendation,
			},
			expectErr: true,
		},
		"No Org ID": {
			given: Recommendation{
				OrganisationID: "",
				MeetingDate: "21/07/2015",
				Recommendation: For,
			},
			expectErr: true,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {

			err := test.given.Validate()

			if err != nil && false == test.expectErr{
				t.Errorf("unexpected error: %v", err.Error())
			}

			if err == nil && test.expectErr {
				t.Errorf("no error when expected")
			}
		})
	}
}