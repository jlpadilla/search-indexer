package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jlpadilla/search-indexer/pkg/server"
	// "github.com/open-cluster-management/insights-client/pkg/config"
)

func main() {

	fmt.Println("Search indexer.")

	router := mux.NewRouter()

	// router.HandleFunc("/liveness", handlers.LivenessProbe).Methods("GET")
	// router.HandleFunc("/readiness", handlers.ReadinessProbe).Methods("GET")
	router.HandleFunc("/aggregator/clusters/{id}/sync", server.SyncResources).Methods("POST")
	// router.HandleFunc("/sync", handlers.SyncResources).Methods("POST")

	// Configure TLS
	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		},
	}
	srv := &http.Server{
		Addr:      ":3010", //config.Cfg.AggregatorAddress,
		Handler:   router,
		TLSConfig: cfg,
		// ReadHeaderTimeout: time.Duration(config.Cfg.HTTPTimeout) * time.Millisecond,
		// ReadTimeout:       time.Duration(config.Cfg.HTTPTimeout) * time.Millisecond,
		// WriteTimeout:      time.Duration(config.Cfg.HTTPTimeout) * time.Millisecond,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler)),
	}

	fmt.Println("Listening on: ", ":3010")
	// glog.Info("Listening on: ", ":3010") // config.Cfg.AggregatorAddress)
	log.Fatal(srv.ListenAndServeTLS("./sslcert/tls.crt", "./sslcert/tls.key"),
		" Use ./setup.sh to generate certificates for local development.")
}
