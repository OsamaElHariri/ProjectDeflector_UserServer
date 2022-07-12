# ProjectDeflector UserServer

This repo holds the code to manage authentication and user users for the `Hit Bounce` mobile game.


## Mobile App

The mobile app, and a high level introduction to this project can be found in the [JsClientGame](https://github.com/OsamaElHariri/ProjectDeflector_JsClientGame) repo, which is intended to be the entry point to understanding this project.


## Overview of This Project

This is a Go server that uses the [Fiber](https://gofiber.io/) web framework. The routes are just in `main.go`, and the function of the routes is to only validate the inputs, then call a use case (which is what runs the business logic). The use cases can be found in the `use_cases.go` file, and these should incapsulate all the functions that this server can do.

Note that this project has a `.devcontainer` and is meant to be run inside a dev container.


## Outputting a Binary

To output the binary of this Go code, run the VSCode task using `CTRL+SHIFT+B`. This should be done while inside the dev container.


Once you have the binary, you need to build the docker image _outside_ the dev container. I use this command and just overwrite the image everytime. This keeps the [Infra](https://github.com/OsamaElHariri/ProjectDeflector_Infra) repo simpler.

```
docker build -t project_deflector/user_server:1.0 .
```


## Authentication Overview

This server handles multiple kinds authentications.

### Public Users

These users are identified by requests that do not have an `x-user-id` field. These public users can become authenticated by calling the `POST` `/public/user` endpoint to get a UUID. They can then use this UUID to get a JWT token by calling the `POST` `/public/access` endpoint.

### Authenticated Users

User authentication is handled via a JWT on the client device. The client sends it's JWT as a Bearer token. The way this JWT becomes an ID in the `x-user-id` is via the load balancer. The load balancer verifies requests by calling the `/auth/check` endpoint on this server. This server then adds the `x-user-id` header to the request if the JWT is valid, or it throws an error.

### Internal Requests

There are some requests that are meant to only be triggered by other servers. The way this is handled is by adding the same secret token to each server, then validating that this secret token exists at the load balancer level whenever an internal route is called. The load balancer does this by calling the `/internal/auth/check` endpoint for internal routes and verifies with this server that the request is valid.

