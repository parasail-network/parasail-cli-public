# Parasail CLI

A command-line tool for generating code and managing staking technology stack projects.

## Overview

Parasail CLI is a powerful command-line interface tool designed to streamline the development of staking technology stack projects. It provides various utilities for project generation, deployment, and management.

## Features

- Project template generation
- Deployment management
- Tool utilities
- Local database management
- Consensus protocol support

## Installation

### Prerequisites

- Go 1.22 or later
- Git

### Building from Source

```bash
git clone https://github.com/your-username/parasail-cli.git
cd parasail-cli
go build
```

## Usage

### Basic Commands

```bash
# Show version
./parasail-cli version

# Generate a new project
./parasail-cli new <project-name>

# Deploy project
./parasail-cli deploy [options]

# Access tool utilities
./parasail-cli tool [subcommand]
```

## Project Structure

```
parasail-cli/
├── generate/         # Code generation utilities
├── contract/         # Contract-related code
├── main.go          # Main entry point
├── service.go       # Service implementation
├── deploy.go        # Deployment logic
├── tool.go          # Tool utilities
├── local_db.go      # Local database management
├── consensus.go     # Consensus protocol implementation
└── Dockerfile.parasail # Docker configuration
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Logrus](https://github.com/sirupsen/logrus) - Logging
- [Gin](https://github.com/gin-gonic/gin) - Web framework
- [Badger](https://github.com/dgraph-io/badger) - Key-value store
- [Ethereum](https://github.com/ethereum/go-ethereum) - Ethereum client

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

[Add your license information here]

## Version

Current version: v0.1.0 