# Example Go Web Service

## Overview

This is an example Go web service that shows the stages of creating an application that is deployed
to a Kubernetes cluster. A real deployment is more complicated but should also involve more 
automation to help deal with the added complexity. The point of this exercise is to clearly show
how we got to using Helm charts.

## Requirements

- Go >= 1.11
- Docker >= 17.05
- Helm
- Kubernetes cluster

## Stages

### Stage One

The first stage deals with building and running the Go web service that outputs the message
`Hello, <APP_NAME>! You are running on <HOST>`. This could use any language and framework but the point
is showing initial application development before it is transitioned into a Docker container.

Run unit tests:

```sh
go test ./src
```

Build the binary:

```sh
go build -o example-go-web-service ./src
```

Run the binary:

```sh
./example-go-web-service
```

Open `http://localhost:8080/` in your browser to view the output.

### Stage Two

The second stage is where we will build the Docker container image and bring the application build
into the Dockerfile as part of a multi-stage build.

Build the container image:

```sh
docker build -t cirrocloud/example-go-web-service:v1.0.0 .
```

Run the container:

```sh
docker run --rm -p 8080:8080 cirrocloud/example-go-web-service:v1.0.0 example-go-web-service
```

Open `http://localhost:8080/` in your browser to view the output.

### Stage Three

The final stage is creating an application deployment using a Helm chart and the container image
built in the previous stage. Helm allows templating and overriding configuration to match the 
environment the container will run in. This allows the application package to remain immutable as 
it is deployed to different environments.

Install the application:

```sh
helm install --namespace example-go-web-service --name example-go-web-service charts/example-go-web-service/
```

Expose the web service:

```sh
kubectl -n example-go-web-service port-forward svc/example-go-web-service 8080:8080
```

Open `http://localhost:8080/` in your browser to view the output.
