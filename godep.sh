#!/bin/bash

go get
GO15VENDOREXPERIMENT=1 godep save ./...
