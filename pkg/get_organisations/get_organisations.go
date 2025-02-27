package get_organisations

import (
	"encoding/json"
	"os"
)

type Organisation struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

const (
	organisationsFilePath = "./pkg/get_organisations/organisations.json"
)

// GetOrganisations returns a map of organisation names to IDs
func GetOrganisations() (map[string]string, error) {
	// Since we are not implementing the API here, here's how you would call the real api:
	// This function acts as a wrapper around what would call the endpoint, you could just add another private
	// func to do the HTTP logic. This way the caller of this package doesn't need to know it comes from an HTTP
	// endpoint, the endpoint could change in future, like to a graphQL endpoint, and calling code doesnt need
	// to change.
	bytes, err := os.ReadFile(organisationsFilePath)
	if err != nil {
		return map[string]string{}, err
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