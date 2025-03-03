package mockserver

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

const (
	organisationsFilePath = "./pkg/mockserver/organisations.json"
)

func Start() {
	http.HandleFunc("/organisations", handleOrganisations)
	http.HandleFunc("/generalmeetings", handleGeneralMeetings)
	http.HandleFunc("/proposals", handleProposals)
	http.HandleFunc("/recommendations", handleRecommendations)

	port := ":8080"
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleOrganisations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	data, err := os.ReadFile(organisationsFilePath)
	if err != nil {
		http.Error(w, "Failed to read organisations data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func handleGeneralMeetings(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	orgID := r.URL.Query().Get("organisation_id")

	meetingsForOrg, ok := generalMeetingsData[orgID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(meetingsForOrg)
}

func handleProposals(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	genMeetingID := r.URL.Query().Get("general_meeting_id")

	proposalsForMeeting, ok := proposalsData[genMeetingID]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(proposalsForMeeting)
}

func handleRecommendations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
