package mockclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"tumelo_task/organisation"
)

type MockClient struct {
	httpClient *http.Client
	baseURL    string
}

// Why are these definitions here? I'm just doing this in a rush and Go prevents import cycles, so they cant be in their correct locations
// This would need refactor for real
type GeneralMeeting struct {
	ID string
	OrganisationID string
	Date string
}

type Proposal struct {
	ID string
	GeneralMeetingID string
	Text string
	Identifier string
}

type RecommendationSubmission struct {
	ProposalIdentifier string `json:"proposal_identifier"`
	RecommendationString string `json:"recommendation"`
}

var (
	client = MockClient{
		httpClient: &http.Client{},
		baseURL: "http://localhost:8080",
	}
)

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

	var organisations []organisation.Organisation
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

func GetGeneralMeetingForOrganisation(organisationID string) ([]GeneralMeeting, error) {
	resp, err := client.httpClient.Get(client.baseURL + "/generalmeetings?organisation_id=" + organisationID)
	if err != nil {
		return []GeneralMeeting{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []GeneralMeeting{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var meetings []GeneralMeeting
	err = json.NewDecoder(resp.Body).Decode(&meetings)
	return meetings, err
}

func GetProposalsForGeneralMeeting(meetingID string) ([]Proposal, error) {
	resp, err := client.httpClient.Get(client.baseURL + "/proposals?general_meeting_id=" + meetingID)
	if err != nil {
		return []Proposal{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []Proposal{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var proposalsForMeeting []Proposal
	err = json.NewDecoder(resp.Body).Decode(&proposalsForMeeting)
	return proposalsForMeeting, err
}

func PostRecommendation(proposalID string, recommendationStr string) error {
	recommendation := RecommendationSubmission{
		ProposalIdentifier: proposalID,
		RecommendationString: recommendationStr,
	}
	
	body, err := json.Marshal(recommendation)
	if err != nil {
		return err
	}

	resp, err := client.httpClient.Post(client.baseURL + "/recommendations", "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}
