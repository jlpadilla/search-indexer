package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_syncRequest(t *testing.T) {
	// TODO: Send request body.
	request := httptest.NewRequest(http.MethodPost, "/aggregator/clusters/test-cluster/sync", nil)
	responseRecorder := httptest.NewRecorder()

	SyncResources(responseRecorder, request)

	expected := SyncResponse{Version: "TBD"}

	if responseRecorder.Code != http.StatusOK {
		t.Errorf("Want status '%d', got '%d'", http.StatusOK, responseRecorder.Code)
	}

	var decodedResp SyncResponse
	err := json.NewDecoder(responseRecorder.Body).Decode(&decodedResp)
	if err != nil {
		t.Error("Unable to decode respoonse body.")
	}

	if fmt.Sprintf("%+v", decodedResp) != fmt.Sprintf("%+v", expected) {
		// if decodedResp != expected {
		t.Errorf("Incorrect response body.\n expected '%+v'\n received '%+v'", expected, decodedResp)
	}
}
