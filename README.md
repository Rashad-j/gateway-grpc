# RESTful gateway for gRPC microservices
This RESTful service serves as a gateway for two gRPC services: [File monitoring](https://github.com/Rashad-j/json-parsing-grcp) and [Remote data store](https://github.com/Rashad-j/binary-search-grpc). 

This gateway uses RESTful api patterns and exposes a number of endpoints to serve its gRPC microservices. It utilizes the use of Gin framework for routing. 

## Design patterns
Factory method pattern is used widely. E.g. the `NewServiceClient` function serves as a factory method. It encapsulates the creation of a ServiceClient instance, providing a way to instantiate the ServiceClient with a `parser.JsonParsingServiceClient` dependency. Other patterns such as Strategy pattern used in `ParserService` which enables different implementation of this interface. 

These showcases important aspects of these design patterns and demonstrates good design practices by promoting separation of concerns and modularity.

## Unit tests
Unit test examples include testing using the `httptest.NewRecorder()` to test our RESTful endpoints for each gRPC service

## Authorization
The gateway expects the users to provide a JWT for authorization header. The service has a placeholder for the auth service implementaiton, but currently will validate any random JWT as a valid for the sake of testing. 

## Docker
This gateway project is built and pushed to docker repository. Can run locally using `docker-compose`.

## How to test?
This service should be run via `docker-compose` in order to allow it running with other microservices. Otherwise, running as a standalone locally need to make sure that other services are running as well.

Run docker command as follows:
```
$ docker compose up
```

Then you can use the commands in makefile, e.g. `make jsonParser` to get all the json files from the json parser service, or use `make search`, `make delete`, or `make insert` to test the remote store service. Have a look on `makefile` to understand more.

Note: that both services are started with some random data to enable you to play with it's service and test it. 