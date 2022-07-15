# Makefile for protobuf ops
export GOPATH = ${HOME}/go/pkg
export GOBIN = ${PWD}/go/bin
export PATH = ${GOROOT}/bin:${GOPATH}/bin:$(shell printenv PATH):${GOBIN}
export PROTOC_URL_PREFIX = https://github.com/protocolbuffers/protobuf/releases/download/
export PROTOC_VERSION = 21.1
export PROTOC_OS = $(shell uname -s)
export PROTOC_ARCH = $(shell uname -m)
export PROTOC_URL_LINUX = ${PROTOC_URL_PREFIX}v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-${PROTOC_OS}-${PROTOC_ARCH}.zip
export PROTOC_URL_OSX = ${PROTOC_URL_PREFIX}v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-osx-universal_binary.zip

.PHONY=prereq
prereq:
	@mkdir -p ${PWD}/go
	@mkdir -p ${PWD}/go/bin

.PHONY=check-if-go-installed
check-if-go-installed:
ifeq (,$(shell which go))
    $(error Go is not installed. Please install Go to build the files.)
endif

.PHONY=intall-node-libraries
install-node-libraries:
	npm i
	cd webapp-template && npm i

.PHONY=install-protoc
install-protoc: prereq
ifeq (${PROTOC_OS}, Darwin)
	@echo "Downloading protoc for OSX"
	$(shell curl -sL ${PROTOC_URL_OSX}  | tar xvf - -C ${PWD}/go/)

else
	@echo "Downloading protoc for Linux"
	$(shell -sL ${PROTOC_URL_LINUX}  | tar xvf - -C ${PWD}/go/)
endif

.PHONY=install-protoc-compiler
install-protoc-compiler:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY=compile-go
compile-go:
	protoc --go_out=grpc_server protobufs/template.proto --go-grpc_out=grpc_server protobufs/template.proto
	protoc --go_out=grpc_client_test protobufs/template.proto --go-grpc_out=grpc_client_test protobufs/template.proto

.PHONY=compile-ts
compile-ts:
	npx protoc --ts_out webapp-template/src --proto_path protobufs protobufs/template.proto --ts_opt ts_nocheck --ts_opt long_type_number --ts_opt generate_dependencies

.PHONY=go-init
go-init:
	cd grpc_server && go mod tidy

.PHONY=build
build: check-if-go-installed prereq install-node-libraries install-protoc install-protoc-compiler compile-go compile-ts


.PHONY=run-server
start-server:
	docker-compose up -d
	echo "Web Server is running at http://localhost:8080"

stop-server:
	docker-compose down