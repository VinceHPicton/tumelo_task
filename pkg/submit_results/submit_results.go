package submit_results

import (
	"sync"
	"tumelo_task/recommendation"
)


func SubmitRecommendations(matchedRecommendations map[string]recommendation.Recommendation) []error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(matchedRecommendations))

	for proposalID, rec := range matchedRecommendations {
		wg.Add(1)
		go func(proposalID string, recommendationStr string) {
			defer wg.Done()
			err := submitRecommendation(proposalID, recommendationStr)
			if err != nil {
				errChan <- err
			}
		}(proposalID, rec.Recommendation)
	}

	wg.Wait()
	close(errChan)

	var errors []error
	for err := range errChan {
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}

type RecommendationSubmission struct {
	ProposalIdentifier string `json:"proposal_identifier"`
	RecommendationString string `json:"recommendation"`
}

func submitRecommendation(proposalID string, recommendationString string) error {

	return nil

	// this is what the real requests could look like
	// client := http.DefaultClient

	// dataStruct := RecommendationSubmission{
	// 	ProposalIdentifier: proposalID,
	// 	RecommendationString: recommendationString,
	// }

	// headers := map[string]string{
	// 	"api_key": os.Getenv("API_KEY"),
	// }

	// _, err := api_caller.PostRequestWithHeaders(client, os.Getenv("TUMELO_API_ADDRESS"), dataStruct, headers)
	// if err != nil {
	// 	return fmt.Errorf("Proposal: %v failed with error: %v", proposalID, err.Error())
	// }

	// return nil
}

