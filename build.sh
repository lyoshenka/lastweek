#!/bin/bash

#GO15VENDOREXPERIMENT=1
go generate $(go list ./... | grep -v /vendor/)
godep go build -v -tags lastweek.go
npm install
node_modules/postcss-cli/bin/postcss --use autoprefixer --use cssnano --output static/css/styles.build.css  static/css/styles.css
