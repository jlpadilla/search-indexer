# Copyright Contributors to the Open Cluster Management project

run:
	go run main.go

test:
	go test ./... -v -coverprofile cover.out