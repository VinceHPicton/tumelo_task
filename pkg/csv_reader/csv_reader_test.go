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

	// TODO: a few more test cases with varying files - eg different numbers of columns
	// This package mostly just calls encoding/csv though so no need to go mad testing a dependency.
	// TODO: including csv files in the repo for testing isnt ideal, however when you build the application with go build
	// These files are ignored by the compiler so it won't affect prod in any way.
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
