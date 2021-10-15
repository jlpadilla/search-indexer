package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jlpadilla/search-indexer/pkg/config"
	"k8s.io/klog/v2"
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

	// klog.Infof("request: %+v \n", r)
	// klog.Infof("body: %+v \n", r.Body)
	// klog.Info("params:", params)
	// klog.Infof("Received request from cluster [%s]", clusterName)

	response := &SyncResponse{Version: config.AGGREGATOR_API_VERSION}
	w.WriteHeader(http.StatusOK)
	encodeError := json.NewEncoder(w).Encode(response)
	if encodeError != nil {
		klog.Error("Error responding to SyncEvent:", encodeError, response)
	}

	// Record metrics.
	OpsProcessed.WithLabelValues(clusterName, r.RequestURI).Inc()
	HttpDuration.WithLabelValues(clusterName, r.RequestURI).Observe(float64(time.Since(start).Milliseconds()))
}
