# meepo-api

This repo stores the protobuf files for all the meepo services.

# How to add protobuf for a new service and generate code

- Install protoc as per [Instruction](https://protobuf.dev/installation/).

- Run `make init` to install all the tools.

- Create a new folder {service_name/}/service/v1, put the {service_name}.proto file in it and add the rpc and message definitions to the protobuf file.

- Run `make api` to generate the client code