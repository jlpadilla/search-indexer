package main

import (
	"github.com/jlpadilla/search-indexer/pkg/config"
	"github.com/jlpadilla/search-indexer/pkg/server"
	"k8s.io/klog/v2"
)

func main() {
	klog.InitFlags(nil)
	klog.Info("Starting search-indexer.")

	config := config.New()
	config.PrintConfig()

	server.StartAndListen()
}
