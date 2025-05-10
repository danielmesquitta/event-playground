# Event Playground

A Go-based event-driven application playground that demonstrates the implementation of event publishing and listening using AWS SNS/SQS through LocalStack for local development.

## Overview

This project serves as a playground for experimenting with event-driven architecture using Go. It implements a publisher-subscriber pattern using AWS SNS (Simple Notification Service) and SQS (Simple Queue Service) for message handling. The project uses LocalStack to provide a local AWS cloud stack for development and testing.

## Prerequisites

- Go 1.24.1 or later
- Docker and Docker Compose
- Make (optional, for using Makefile commands)

## Project Structure

```
.
├── cmd/                   # Application entry points
│   ├── publisher/         # Event publisher service
│   └── listener/          # Event listener service
├── internal/              # Private application code
│   ├── app/               # Application-specific code
│   ├── domain/            # Domain models and business logic
│   └── pkg/               # Shared packages
├── docker-compose.yml     # Docker services configuration
├── go.mod                 # Go module definition
└── go.sum                 # Go module checksums
```

## Getting Started

1. Clone the repository:

   ```bash
   git clone https://github.com/danielmesquitta/event-playground.git
   cd event-playground
   ```

2. Start the services (LocalStack SNS & SQS) using Docker Compose:

   ```bash
   docker-compose up -d
   ```

3. Run the listener:

   ```bash
   make listener
   ```

4. Run the publisher:

   ```bash
   make publisher
   ```

## Infrastructure

The project uses the following infrastructure components:

- **LocalStack**: Provides a local AWS cloud stack for development
  - SNS (Simple Notification Service)
  - SQS (Simple Queue Service)
- **Go**: Main application runtime
- **Docker**: Containerization for consistent development environment

## Dependencies

Main dependencies include:

- github.com/ThreeDotsLabs/watermill: Event-driven architecture library
- github.com/ThreeDotsLabs/watermill-aws: AWS integration for Watermill
