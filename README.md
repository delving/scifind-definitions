# README

This repository contains the `protobuf` definitions and services for the "Vind het uit" project that aims to bring together the collections of Dutch Science museums.

## Useful commands

- lint the repository:

  `buf lint proto`

- generate the implementations

  buf generate proto


## Requirements

  - Php plugin: 
- Protoc
  - docs: https://github.com/protobuf-php/protobuf-plugin
  - installation command: `composer require "protobuf-php/protobuf-plugin"`
- Go plugin: 
  - quickstart: https://grpc.io/docs/languages/go/quickstart/
