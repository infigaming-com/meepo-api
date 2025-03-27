# meepo-api

This repo stores the protobuf files for all the meepo services

# How to add protobuf for the new service

Let's say the new service is user.

```bash
// Create user.proto under api/user/service/v1 folder
$ kratos proto add api/user/service/v1/user.proto
// generate the proto client code with either of these two commands
$ kratos proto client api/user/service/v1/user.proto
$ make api
```