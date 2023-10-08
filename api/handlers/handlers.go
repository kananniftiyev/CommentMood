package handlers

import (
	"Sentitube/api/models"
	"Sentitube/app"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func PostVideoID(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}

	var videoReq models.VideoRequest
	err = json.Unmarshal(requestBody, &videoReq)
	if err != nil {
		http.Error(w, "Failed to parse JSON request", http.StatusBadRequest)
	}
	requestData := videoReq.VideoID
	positive, neutral, negative, not := app.GetYoutubeAnalysis(requestData)

	response := models.Response{
		VideoID:  requestData,
		Positive: positive,
		Neutral:  neutral,
		Negative: negative,
		Not:      not,
	}
	// Marshal the response struct to JSON
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
	}

	// Write the JSON response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func PostInstagramPostID(w http.ResponseWriter, r *http.Request) {

}
