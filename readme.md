# Ports

A gRPC service to store port data.

And a client api to read in a list of ports from a json file and save to the service.

## How to Use

Run `make start logs`.

This should pull images from Docker hub and run them. If this doesn't work run `tag=0.0.1 make build` to build the Docker images.

To get ports from the client, browse to http://localhost:8080/ports.

To stop press `Ctrl-C` or run `make stop`.

To run tests, run `make test`.

# Roadmap

Here are things I'd like to do in the future.

1. Refactor domains in portclient to a `ports` domain instead of `reader` and `repo`.

2. Save more than one port at a time to the gRPC service. Probably best way to do this is make use of the gRPC `stream` keyword on methods to send ports to service.

3. Returning a lot of data with an HTTP GET could be slow - do some kind of pagination?

4. Get specific port from client with an `id` query parameter.

5. Move code to get environment variables into a `config` package. Configure things like the port numbers for the services for more flexibility.

6. `make proto` command - write Dockerfile and run this in a container so it doesn't rely on the user installing `protobuf-compiler` or its plugins.

7. Use https://github.com/go-swagger/go-swagger to document REST API. (Maybe overkill, with only one endpoint.)

8. Finish off the repo domain, could do with a refactor to ease testing.

9. Add tests to `portsvc`, look at any other possible tests for `portclient`.

10. Add integration tests for each service and end-to-end tests for both services.
