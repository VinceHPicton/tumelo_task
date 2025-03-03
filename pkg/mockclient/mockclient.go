package mockclient

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// MockClient struct to encapsulate the HTTP client
type MockClient struct {
	httpClient *http.Client
	baseURL    string
}

type Organisation struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

type GeneralMeeting struct {
	ID string
	OrganisationID string
	Date string
}

var (
	client = MockClient{
		httpClient: &http.Client{},
		baseURL: "http://localhost:8080",
	}
)

// GetOrganisations fetches organisations from the API
func GetOrganisations() (map[string]string, error) {
	resp, err := client.httpClient.Get(client.baseURL + "/organisations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read body: %v", err.Error())
	}

	var organisations []Organisation
	err = json.Unmarshal(bytes, &organisations)
	if err != nil {
		return map[string]string{}, err
	}

	orgNamesToID := map[string]string{}
	for _, org := range organisations {
		orgNamesToID[org.Name] = org.ID
	}

	return orgNamesToID, nil
}

// GetGeneralMeetings fetches general meetings
func GetGeneralMeetingForOrganisation(organisationID string) ([]GeneralMeeting, error) {
	resp, err := client.httpClient.Get(client.baseURL + "/generalmeetings?organisation_id=" + organisationID)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var meetings []GeneralMeeting
	err = json.NewDecoder(resp.Body).Decode(&meetings)
	return meetings, err
}

// PostRecommendation sends a recommendation
// func PostRecommendation(recommendation map[string]interface{}) error {
	// body, err := json.Marshal(recommendation)
	// if err != nil {
	// 	return err
	// }

	// resp, err := c.httpClient.Post(c.baseURL+"/recommendations", "application/json", bytes.NewReader(body))
	// if err != nil {
	// 	return err
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusCreated {
	// 	return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	// }
	// return nil
// }

// func main() {
// 	// Use the local mock server
// 	client := NewClient("http://localhost:8080")

// 	// Fetch Organisations
// 	orgs, err := client.GetOrganisations()
// 	if err != nil {
// 		log.Fatalf("Error fetching organisations: %v", err)
// 	}
// 	fmt.Println("Organisations:", orgs)

// 	// Fetch General Meetings
// 	meetings, err := client.GetGeneralMeetings()
// 	if err != nil {
// 		log.Fatalf("Error fetching general meetings: %v", err)
// 	}
// 	fmt.Println("General Meetings:", meetings)

// 	// Send a Recommendation
// 	rec := map[string]interface{}{
// 		"proposal_id":   1001,
// 		"recommendation": "For",
// 	}
// 	err = client.PostRecommendation(rec)
// 	if err != nil {
// 		log.Fatalf("Error sending recommendation: %v", err)
// 	}
// 	fmt.Println("Recommendation sent successfully!")
// }
