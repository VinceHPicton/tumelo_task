package recommendation

import (
	"fmt"
)

func ParseCSVDataToRecommendations(csvData *[][]string) ([]Recommendation, error) {
	// Why are we taking a pointer here? Performance: input may be very large and there's no need to copy it for this by passing by value

	// TODO: unit tests for this and mapCSVLineToRecommendation
	recommendations := []Recommendation{}

	for _, csvLine := range *csvData {
		newRecommendation, err := mapCSVLineToRecommendation(csvLine)
		if err != nil {
			return []Recommendation{}, err
		}
		recommendations = append(recommendations, newRecommendation)
	}

	return recommendations, nil
}

func mapCSVLineToRecommendation(csvLine []string) (Recommendation, error) {
	if len(csvLine) != expectedNumberOfColumnsPerCSVLine {
		return Recommendation{}, fmt.Errorf("expected %d columns in record, actual: %d", expectedNumberOfColumnsPerCSVLine, len(csvLine))
	}

	// There is some potential for a bug to appear here, if someone haphazardly changed the csv file input and expectedNumberOfColumnsPerCSVLine but not the mapper
	return Recommendation{
		Name: csvLine[0],
		MeetingDate: csvLine[1],
		SequenceID: csvLine[2],
		ProposalText: csvLine[3],
		Recommendation: csvLine[4],
	}, nil
}