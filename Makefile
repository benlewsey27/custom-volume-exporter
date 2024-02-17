version := $(shell cat VERSION)

build:
	go build -o bin/customvolumeexporter

compile:
	GOOS=darwin GOARCH=amd64 go build -o "bin/customvolumeexporter-darwin-amd64-$(version)"
	GOOS=linux GOARCH=amd64 go build -o "bin/customvolumeexporter-linux-amd64-$(version)"
