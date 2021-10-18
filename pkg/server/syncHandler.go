package server

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jlpadilla/search-indexer/pkg/config"
	"k8s.io/klog/v2"
)

// Resource - Describes a resource (node)
type Resource struct {
	Kind           string `json:"kind,omitempty"`
	UID            string `json:"uid,omitempty"`
	ResourceString string `json:"resourceString,omitempty"`
	Properties     map[string]interface{}
}

// Describes a relationship between resources
type Edge struct {
	SourceUID, DestUID   string
	EdgeType             string
	SourceKind, DestKind string
}

// DeleteResourceEvent - Contains the information needed to delete an existing resource.
type DeleteResourceEvent struct {
	UID string `json:"uid,omitempty"`
}

// SyncEvent - Object sent by the collector with the resources to change.
type SyncEvent struct {
	ClearAll bool `json:"clearAll,omitempty"`

	AddResources    []Resource
	UpdateResources []Resource
	DeleteResources []DeleteResourceEvent

	AddEdges    []Edge
	DeleteEdges []Edge
	RequestId   int
}

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
	klog.V(2).Infof("Processing request from cluster [%s]", clusterName)

	var syncEvent SyncEvent
	err := json.NewDecoder(r.Body).Decode(&syncEvent)
	if err != nil {
		klog.Error("Error decoding body of syncEvent: ", err)
		// respond(http.StatusBadRequest)
		return
	}

	klog.Infof("Request body(decoded): %+v \n", syncEvent)

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
