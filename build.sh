#!/bin/bash

#GO15VENDOREXPERIMENT=1
go generate $(go list ./... | grep -v /vendor/)
godep go build -v -tags lastweek.go
