package csv_reader

import (
	"reflect"
	"testing"
	"tumelo_task/recommendation"
)

var expectedSimpleData = []recommendation.Recommendation{
	{
		Name: "Organisation Name",
		MeetingDate: "Meeting Date",
		SequenceID: "Sequence Identifier",
		ProposalText: "Proposal Text",
		Recommendation: "Recommendation",
	},
	{
		Name: "Mckesson Corporation",
		MeetingDate: "21/07/2023",
		SequenceID: "1a.",
		ProposalText: "Elect Richard H. Carmona",
		Recommendation: "For",
	},
}
func Test_Read(t *testing.T) {

	testCases := map[string]struct {
		filepath string
		expected []recommendation.Recommendation
	}{
		"happy path": {
			filepath: "../ExampleRecommendationsSimple.csv",
			expected: expectedSimpleData,
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := Read[recommendation.Recommendation](test.filepath)
			if err != nil {
				t.Errorf(err.Error())
			}

			if false == reflect.DeepEqual(got, test.expected) {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}



func Test_ReadIgnoringHeader(t *testing.T) {

	testCases := map[string]struct {
		filepath string
		expected []recommendation.Recommendation
	}{
		"happy path": {
			filepath: "../ExampleRecommendationsSimple.csv",
			expected: expectedSimpleData[1:],
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := ReadIgnoringHeader[recommendation.Recommendation](test.filepath)
			if err != nil {
				t.Errorf(err.Error())
			}

			if false == reflect.DeepEqual(got, test.expected) {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
