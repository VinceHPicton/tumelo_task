package csv_reader

import (
	"reflect"
	"testing"
)

var expectedSimpleData = [][]string{
	{
		"Organisation Name",
		"Meeting Date",
		"Sequence Identifier",
		"Proposal Text",
		"Recommendation",
	},
	{
		"Mckesson Corporation",
		"21/07/2023",
		"1a.",
		"Elect Richard H. Carmona",
		"For",
	},
}

const (
	simpleExampleDataFilePath = "./ExampleRecommendationsSimple.csv"
	emptyFilePath = "./EmptyFile.csv"
)

func Test_Read(t *testing.T) {

	testCases := map[string]struct {
		filepath string
		expected [][]string
	}{
		"happy path": {
			filepath: simpleExampleDataFilePath,
			expected: expectedSimpleData,
		},
		"empty file": {
			filepath: emptyFilePath,
			expected: [][]string{},
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := Read(test.filepath)
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
		expected [][]string
	}{
		"happy path": {
			filepath: simpleExampleDataFilePath,
			expected: expectedSimpleData[1:],
		},
		"empty file": {
			filepath: emptyFilePath,
			expected: [][]string{},
		},
	}

	for name, test := range testCases {
		t.Run(name, func(t *testing.T) {
			got, err := ReadIgnoringHeader(test.filepath)
			if err != nil {
				t.Errorf(err.Error())
			}

			if false == reflect.DeepEqual(got, test.expected) {
				t.Errorf("expected: %v, got: %v", test.expected, got)
			}
		})
	}
}
