package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_syncRequest(t *testing.T) {
	// TODO: Send request body.
	body, readErr := os.Open("./mocks/simple.json")
	if readErr != nil {
		t.Fatal(readErr)
	}
	fmt.Printf("Body: %+v\n", body)

	// bytes, err1 := ioutil.ReadFile("./mocks/simple.json")

	// var data map[string]interface{}
	// if err := json.Unmarshal(bytes, &data); err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Body: %+v", data)

	request := httptest.NewRequest(http.MethodPost, "/aggregator/clusters/test-cluster/sync", body)
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

// func readMock(file string) (data map[string]interface{}) {
// 	// bytes, _ := ioutil.ReadFile("./data/sno-0.json")
// 	bytes, _ := ioutil.ReadFile(file)
// 	// var data map[string]interface{}
// 	if err := json.Unmarshal(bytes, &data); err != nil {
// 		panic(err)
// 	}
// }
