# 🚚 Order Microservice with gRPC 🚚

## Overview

This repository contains a simple **Order Microservice** implemented using **gRPC**. The microservice provides basic order management functionality, including creating, retrieving, updating, and deleting orders. The project helped me learn how to build a **gRPC server**, work with **Protocol Buffers (protobuf)**, and implement client-server communication in a distributed system.

## What is gRPC?

[gRPC](https://grpc.io/) is a high-performance, open-source RPC (Remote Procedure Call) framework that uses HTTP/2 for transport, Protocol Buffers for serialization, and it supports multiple languages. It is ideal for building scalable, distributed systems.

## Features ✨

- **Create Order**: Create a new order with details such as product name, quantity, and customer information.
- **Get Order**: Retrieve an order by its unique identifier.
- **Update Order**: Update the details of an existing order.
- **Delete Order**: Delete an existing order by its unique identifier.
- **gRPC Communication**: The service uses gRPC for fast, efficient communication between clients and the server.
- **Protobuf Definitions**: The service defines the data structure and service methods using Protocol Buffers, which are language-neutral and platform-neutral.

## Technologies Used ⚙️

- **gRPC**: A modern open-source RPC framework.
- **Protocol Buffers (protobuf)**: A language-agnostic interface definition language for serializing structured data.
- **Go**: The backend language used to implement the gRPC server.

## Learning Goals 🎓

- **gRPC Server**: Building and running a gRPC server.
- **Protobuf Syntax**: Defining messages and service methods in `.proto` files.
- **gRPC Communication**: Understanding how gRPC clients communicate with servers.
- **Microservices Architecture**: Understanding how microservices can communicate using gRPC.

## Installation 🚀

To run the Order Microservice, follow these steps:

### Prerequisites 🔧

- **Go**: Version 1.16 or higher
- **Protocol Buffers (protoc)**: Used to generate the gRPC code from `.proto` files. [Installation Guide](https://grpc.io/docs/protoc-installation/)
- **Make**: A build automation tool to simplify the process. [Installation Guide](https://www.gnu.org/software/make/)

### Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/amal5haji/grpc-order-microservice.git
   cd grpc-order-microservice
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Generate the gRPC code from the `.proto` file using the `Makefile`:

   ```bash
   make protoc
   ```

   This will run the `protoc` command and generate the necessary Go files for the gRPC server and client based on the `order.proto` file.

4. Run the gRPC server:

   ```bash
   go run ./cmd/server/main.go
   ```

   The server will start on port `50051` by default.

### Client 💻

The project also includes a **gRPC client** to interact with the server. The client can perform the following operations:

- **Create an Order**: Create a new order with a product name, quantity, and customer.
- **Get an Order**: Retrieve an existing order by its unique ID.
- **Update an Order**: Modify an existing order's details.
- **Delete an Order**: Remove an order from the system.


### `Makefile` ⚙️

The `Makefile` simplifies the process of generating gRPC code from the `.proto` file. To regenerate the code, simply run:

```bash
make protoc
```

The `Makefile` contains the following content:

```Makefile
# Makefile for generating gRPC code from .proto files

PROTOC_GEN_GO_PATH=$(shell go list -f '{{.Dir}}' github.com/golang/protobuf/protoc-gen-go)

# The target protoc command
protoc:
	protoc --go_out=. --go-grpc_out=. order.proto
```
