#! /usr/bin/make -f

# GOROOT may need to be set if Go's installed somewhere unusual, e.g.
#     GOROOT=/home/stephen/go/go1.5.1.linux-amd64
# PATH needs to contain $GOROOT/bin.
# GOPATH needs to be set, e.g. export GOPATH=/Users/stephen/Documents/Development/translating_nature/jertranslate
#
# To build from source for local architecture do:
#     make var check all
#
# To build from source for new executable suitable for deployment, e.g.
# stage2's architecture, do:
#     make var check deploy

SHELL = /bin/bash

all:
	go install jertranslate   # Native executable in ../bin/jertranslate.

var:
	[[ -d $$GOPATH ]]
	[[ -f $$GOPATH/src/jertranslate/main.go ]]

check:
	go vet  jertranslate
	go test jertranslate

bench:
	go test -bench ^ -benchmem jertranslate

fmt:
	go fmt jertranslate

deploy:
	GOOS=linux GOARCH=amd64 go install -installsuffix deploy jertranslate

clean:
	rm -f ../bin/jertranslate
	rm -f ../bin/*/jertranslate
