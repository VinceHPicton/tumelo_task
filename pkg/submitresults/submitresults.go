package submitresults

import (
	"sync"
	"tumelo_task/pkg/mockclient"
	"tumelo_task/recommendation"
)


func SubmitRecommendations(matchedRecommendations map[string]recommendation.Recommendation) []error {
	var wg sync.WaitGroup
	errChan := make(chan error, len(matchedRecommendations))

	// I've used a concurrency approach here just to demonstrate it, it's very similar to what you'd want to do in several other places
	for proposalID, rec := range matchedRecommendations {
		wg.Add(1)
		go func(proposalID string, recommendationStr string) {
			defer wg.Done()
			err := mockclient.PostRecommendation(proposalID, recommendationStr)
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

