# Copyright Contributors to the Open Cluster Management project


setup:
	cd sslcert; openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -config req.conf -extensions 'v3_req'

run:
	go run main.go

test:
	go test ./... -v -coverprofile cover.out