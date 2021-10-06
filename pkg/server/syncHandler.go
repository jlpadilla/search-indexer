package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// SyncResponse - Response to a SyncEvent
type SyncResponse struct {
	TotalAdded        int
	TotalUpdated      int
	TotalDeleted      int
	TotalResources    int
	TotalEdgesAdded   int
	TotalEdgesDeleted int
	TotalEdges        int
	AddErrors         []SyncError
	UpdateErrors      []SyncError
	DeleteErrors      []SyncError
	AddEdgeErrors     []SyncError
	DeleteEdgeErrors  []SyncError
	Version           string
	RequestId         int
}

// SyncError is used to respond with errors.
type SyncError struct {
	ResourceUID string
	Message     string // Often comes out of a golang error using .Error()
}

func SyncResources(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	clusterName := params["id"]

	fmt.Printf("request: %+v \n", r)
	fmt.Println("params:", params)
	log.Printf("Received request from [%s]", clusterName)

	response := &SyncResponse{Version: "TBD"} // TODO
	w.WriteHeader(http.StatusOK)
	encodeError := json.NewEncoder(w).Encode(response)
	if encodeError != nil {
		fmt.Println("Error responding to SyncEvent:", encodeError, response)
	}

	// Record metrics.
	OpsProcessed.WithLabelValues(clusterName, r.RequestURI).Inc()
	HttpDuration.WithLabelValues(clusterName, r.RequestURI).Observe(float64(time.Since(start).Milliseconds()))
}
