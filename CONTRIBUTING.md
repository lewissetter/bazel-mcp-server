# Contributing to Bazel MCP Server

Thank you for your interest in contributing to Bazel MCP Server! This document provides guidelines and instructions for contributing.

## Code of Conduct

Please be respectful and considerate in your interactions with other contributors.

## How to Contribute

### Reporting Bugs

Before creating bug reports, please check existing issues to avoid duplicates. When creating a bug report, include:

- A clear and descriptive title
- Steps to reproduce the issue
- Expected behavior
- Actual behavior
- Your environment (OS, Go version, Bazel version)
- Any relevant logs or error messages

### Suggesting Enhancements

Enhancement suggestions are welcome! Please provide:

- A clear and descriptive title
- Detailed description of the proposed feature
- Use cases and examples
- Any potential drawbacks or considerations

### Pull Requests

1. Fork the repository
2. Create a new branch from `main`:
   ```bash
   git checkout -b feature/your-feature-name
   ```
3. Make your changes
4. Add or update tests as needed
5. Ensure all tests pass:
   ```bash
   go test ./...
   ```
6. Format your code:
   ```bash
   go fmt ./...
   ```
7. Commit your changes with a clear message:
   ```bash
   git commit -m "Add feature: description of your changes"
   ```
8. Push to your fork:
   ```bash
   git push origin feature/your-feature-name
   ```
9. Open a Pull Request

## Development Setup

### Prerequisites

- Go 1.21 or later
- Bazel (for testing with actual Bazel projects)
- Git

### Building

```bash
go build -o bazel-mcp-server ./cmd/bazel-mcp-server
```

### Running Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with race detector
go test -race ./...
```

### Project Structure

```
bazel-mcp-server/
├── cmd/
│   └── bazel-mcp-server/    # Main application
│       ├── main.go           # Server implementation
│       └── main_test.go      # Tests
├── .github/
│   └── workflows/            # GitHub Actions CI/CD
├── .goreleaser.yml           # Release configuration
├── go.mod                    # Go module definition
├── LICENSE                   # MIT License
└── README.md                 # Project documentation
```

## Coding Guidelines

### Style

- Follow standard Go conventions and idioms
- Use `gofmt` to format your code
- Run `go vet` to catch common mistakes
- Keep functions focused and reasonably sized
- Add comments for exported types and functions

### Testing

- Write tests for new functionality
- Ensure existing tests still pass
- Aim for good test coverage
- Use table-driven tests where appropriate

### Commits

- Write clear, concise commit messages
- Use present tense ("Add feature" not "Added feature")
- Reference issues and pull requests where appropriate

### Documentation

- Update README.md if adding new features
- Add inline documentation for complex logic
- Update examples if behavior changes

## Release Process

Releases are automated using GitHub Actions and GoReleaser:

1. Create and push a new tag:
   ```bash
   git tag -a v0.2.0 -m "Release v0.2.0"
   git push origin v0.2.0
   ```

2. GitHub Actions will:
   - Run all tests
   - Build binaries for multiple platforms
   - Create a GitHub release
   - Upload release artifacts

## Questions?

Feel free to open an issue with your question or start a discussion.

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
