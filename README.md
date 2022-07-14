# grpc-envoy-template

This repo is a template/example for implementing a protobuf-base microservice running on the `grpc-web`/`envoy proxy` stack.

## Motivation

When first implementing this stack, I've found the gRPC and grpc-web's site a bit outdated/lacking with regards to implementing the grpc-web and envoy proxy stack. In particular, the fussy nature of protoc can make stitching each of these components together a bit time-consuming.

This repo is a standard-ish template that I've built on to impement this sort of microservice stack.

## Prerequsites

- Docker