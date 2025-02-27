package recommendation

import (
	"strings"
	"unicode"
)

func CleanAllRecommendations(recommendations *[]Recommendation) {
	// Why are we taking a pointer here? Performance: input may be very large and there's no need to copy it for this by passing by value
	for i, _ := range *recommendations {
		(*recommendations)[i].CleanData()
	}
}

// CleanData does some very basic data cleaning, such as removing trailing/preceeding whitespace, making strings all lowercase
// For the purposes of validation
func (r *Recommendation) CleanData() {
	r.Name = strings.TrimSpace(r.Name)
	r.Recommendation = convertToCapitalisedFirstLetterString(r.Recommendation)
}

// convertToCapitalisedFirstLetterString converts any string like "aNyTHing" to "Anything"
func convertToCapitalisedFirstLetterString(givenString string) string {
	if givenString == "" {
		return givenString
	}

	lowerCaseString := strings.ToLower(givenString)

	runes := []rune(lowerCaseString)
	runes[0] = unicode.ToUpper(runes[0])

	return string(runes)
}
