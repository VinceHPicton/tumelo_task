package recommendation

import (
	"reflect"
	"testing"
)

func Test_CleanData(t *testing.T) {

	initialRec := Recommendation{
		OrganisationID: "432aaae7-500b-492c-bc87-360437b37354",
		MeetingDate: "21/07/2015",
		Recommendation: For,
	}

	testCases := map[string]struct {
		given Recommendation
		expected Recommendation
	}{
		"happy path": {
			given: Recommendation{
				OrganisationID: initialRec.OrganisationID,
				MeetingDate: initialRec.MeetingDate,
				Recommendation: "for",
			},
			expected: Recommendation{
				OrganisationID: initialRec.OrganisationID,
				MeetingDate: initialRec.MeetingDate,
				Recommendation: "For",
			},
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {

			test.given.CleanData()

			if !reflect.DeepEqual(test.given, test.expected) {
				t.Errorf("got: %v, expected: %v", test.given, test.expected)
			}
		})
	}
}

func Test_convertToCapitalisedFirstLetterString(t *testing.T) {

	testCases := map[string]struct {
		given string
		expected string
	}{
		"happy path": {
			given: "for",
			expected: "For",
		},
		"happy path 2": {
			given: "aNyThING",
			expected: "Anything",
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {

			got := convertToCapitalisedFirstLetterString(test.given)

			if got != test.expected {
				t.Errorf("got: %v, expected: %v", got, test.expected)
			}
		})
	}
}