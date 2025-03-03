package recommendation

import (
	"fmt"
	"time"
)

type InvalidRecommendation struct {
	Recommendation Recommendation
	Reason string
	// Although there is some coupling to the concept of recommendation being from a CSV file here with OriginalIndex, if this
	// struct was used for something else the field can just be ignored
	OriginalIndex int
}

func FindInvalidRecommendations(recommendationsPtr *[]Recommendation) ([]InvalidRecommendation) {
	// Why are we taking a pointer here? Performance: input may be very large and there's no need to copy it for this by passing by value

	invalidRecommendations := []InvalidRecommendation{}
	for index, recommendation := range *recommendationsPtr {
		validationError := recommendation.Validate()
		if validationError != nil {
			invalidRecommendations = append(invalidRecommendations, InvalidRecommendation{
				Recommendation: recommendation,
				Reason: validationError.Error(),
				OriginalIndex: index,
			})
		}
	}

	return invalidRecommendations
}

// Validate tests the data in the recommendation, returns nil if nothing invalid, or an error containing the fail reason if not.
func (r *Recommendation) Validate() error {
	if r.OrganisationID == "" {
		return fmt.Errorf("organisation ID not found: %v", *r)
	}

	if !validateRecommendationString(r.Recommendation) {
		return fmt.Errorf("recommendation string was invalid: %v", *r)
	}

	if !validateDate(r.MeetingDate) {
		return fmt.Errorf("date was invalid: %v", *r)
	}

	return nil
}

func validateDate(dateStr string) bool {
	_, err := time.Parse(validDateFormat, dateStr)
	return err == nil
}

func validateRecommendationString(recommendationString string) bool {
	switch recommendationString {
	case For, Abstain, Against:
		return true
	default:
		return false
	}
}