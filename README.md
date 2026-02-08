# Bazel MCP Server

[![Go Reference](https://pkg.go.dev/badge/github.com/yourusername/bazel-mcp-server.svg)](https://pkg.go.dev/github.com/yourusername/bazel-mcp-server)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/bazel-mcp-server)](https://goreportcard.com/report/github.com/yourusername/bazel-mcp-server)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

An [MCP (Model Context Protocol)](https://modelcontextprotocol.io/) server for interacting with the [Bazel](https://bazel.build/) build system. This server enables AI assistants and other MCP clients to build, test, query, and manage Bazel projects.

## Features

The server provides the following MCP tools:

- **bazel_build** - Build one or more targets
- **bazel_test** - Run tests with filtering and custom arguments
- **bazel_clean** - Clean build outputs (with expunge/async options)
- **bazel_run** - Build and run a single target with arguments
- **bazel_query** - Query the dependency graph
- **bazel_aquery** - Query the action graph
- **bazel_info** - Display Bazel server runtime info

## Installation

### From Source

```bash
go install github.com/yourusername/bazel-mcp-server/cmd/bazel-mcp-server@latest
```

### From Release

Download the appropriate binary for your platform from the [releases page](https://github.com/yourusername/bazel-mcp-server/releases).

## Usage

### With Claude Desktop

Add the server to your Claude Desktop configuration:

**macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows**: `%APPDATA%/Claude/claude_desktop_config.json`
**Linux**: `~/.config/Claude/claude_desktop_config.json`

```json
{
  "mcpServers": {
    "bazel": {
      "command": "/path/to/bazel-mcp-server"
    }
  }
}
```

### With Other MCP Clients

The server communicates over stdin/stdout using the MCP protocol. Run it directly and connect your client to its stdio streams.

```bash
bazel-mcp-server
```

## Tools Reference

### bazel_build

Build one or more targets using Bazel.

**Arguments:**
- `targets` (required): List of targets to build (e.g., `["//pkg:target"]` or `["//..."]`)
- `options` (optional): Additional build options (e.g., `["--config=opt", "--compilation_mode=dbg"]`)

**Example:**
```json
{
  "name": "bazel_build",
  "arguments": {
    "targets": ["//cmd/app:main", "//pkg/lib:library"],
    "options": ["--config=release"]
  }
}
```

### bazel_test

Run tests for one or more targets using Bazel.

**Arguments:**
- `targets` (required): List of test targets to run (e.g., `["//pkg:test"]` or `["//..."]`)
- `options` (optional): Additional test options (e.g., `["--test_output=all"]`)
- `test_args` (optional): Arguments to pass to the test binary
- `test_filter` (optional): Filter to select specific tests to run

**Example:**
```json
{
  "name": "bazel_test",
  "arguments": {
    "targets": ["//pkg:test"],
    "options": ["--test_output=all"],
    "test_filter": "TestMyFunction"
  }
}
```

### bazel_clean

Clean Bazel build outputs.

**Arguments:**
- `expunge` (optional): Remove the entire working tree for this Bazel instance
- `async` (optional): Clean asynchronously
- `options` (optional): Additional clean options

**Example:**
```json
{
  "name": "bazel_clean",
  "arguments": {
    "expunge": true
  }
}
```

### bazel_run

Build and run a single target using Bazel.

**Arguments:**
- `target` (required): Single target to build and run (e.g., `"//cmd/app:main"`)
- `options` (optional): Additional build/run options
- `args` (optional): Arguments to pass to the target binary

**Example:**
```json
{
  "name": "bazel_run",
  "arguments": {
    "target": "//cmd/server:server",
    "args": ["--port=8080", "--verbose"]
  }
}
```

### bazel_query

Query the Bazel dependency graph.

**Arguments:**
- `expression` (required): Query expression (e.g., `"deps(//pkg:target)"` or `"kind(go_library, //...)"`)
- `options` (optional): Additional query options (e.g., `["--output=label"]`)

**Example:**
```json
{
  "name": "bazel_query",
  "arguments": {
    "expression": "deps(//cmd/app:main)",
    "options": ["--output=label"]
  }
}
```

### bazel_aquery

Query information about actions in the build graph (action graph query).

**Arguments:**
- `expression` (required): Action query expression (e.g., `"mnemonic('GoCompile', deps(//...))"`)
- `options` (optional): Additional aquery options (e.g., `["--output=jsonproto"]`)

**Example:**
```json
{
  "name": "bazel_aquery",
  "arguments": {
    "expression": "mnemonic('GoCompile', deps(//...))",
    "options": ["--output=jsonproto"]
  }
}
```

### bazel_info

Display runtime info about the Bazel server.

**Arguments:**
- `keys` (optional): Specific info keys to retrieve (e.g., `["bazel-bin", "execution_root"]`)
- `options` (optional): Additional info options

**Example:**
```json
{
  "name": "bazel_info",
  "arguments": {
    "keys": ["bazel-bin", "bazel-genfiles"]
  }
}
```

## Requirements

- Go 1.21 or later (for building from source)
- Bazel must be installed and available in your PATH
- The server executes Bazel commands in the current working directory

## Development

### Building

```bash
go build -o bazel-mcp-server ./cmd/bazel-mcp-server
```

### Testing

```bash
go test ./...
```

### Running Locally

```bash
go run ./cmd/bazel-mcp-server
```

## How It Works

The server:

1. Starts and listens for MCP requests over stdin/stdout
2. Validates incoming tool call arguments
3. Executes the corresponding Bazel command in the current working directory
4. Returns the command output (both stdout and stderr) to the client
5. Reports errors as tool errors (not protocol errors) for better user experience

All commands use `context.Context` for proper cancellation support.

## Example Interactions

**Building a target:**
```
Tool: bazel_build
Arguments: {"targets": ["//cmd/app:main"]}
Result: "Build successful!\n\nINFO: Build completed successfully"
```

**Running tests:**
```
Tool: bazel_test
Arguments: {"targets": ["//..."], "options": ["--test_output=errors"]}
Result: "Tests passed!\n\n//pkg:test PASSED"
```

**Querying dependencies:**
```
Tool: bazel_query
Arguments: {"expression": "deps(//cmd/app:main)", "options": ["--output=label"]}
Result: "//cmd/app:main\n//pkg/lib:library\n//external:go_sdk"
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built with the [Model Context Protocol Go SDK](https://github.com/modelcontextprotocol/go-sdk)
- Inspired by the need to integrate AI assistants with Bazel build workflows

## Support

- 🐛 [Report a bug](https://github.com/yourusername/bazel-mcp-server/issues/new?labels=bug)
- ✨ [Request a feature](https://github.com/yourusername/bazel-mcp-server/issues/new?labels=enhancement)
- 💬 [Ask a question](https://github.com/yourusername/bazel-mcp-server/discussions)
