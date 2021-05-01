# Golang Example Application

# Table of Contents

- [Overview](#overview)
- [Package list](#package-list)
- [Installing](#installing)
    * [Local environment](#local-environment)
    * [Docker environment](#docker-enviroment)
- [Run services](#run-services)
    * [Start in local](#start-in-local-machine)
    * [Start in docker](#start-in-docker)
- [Getting started](#getting-started)
    * [Jaeger](#jaeger)
    * [Http with gRPC](#http-example-with-grpc)
    * [Graphql with gRPC](#graphql-example-with-grpc)
- [Deprecate version](#deprecated-version)
- [Testing](#testing)

# Overview

This is an example golang application.
Commands list:
1. Daemon - main service
2. Product service - service that returns Product, an example of gRPC client/server interaction
3. Health check service - this service is needed to show how convenient to understand on which of the services an error occurred in jaeger
4. Migrate - commands for migration
5. JWT - commands for generate JWT token