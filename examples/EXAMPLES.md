# Example Interactions with Bazel MCP Server

This document provides example interactions showing how to use the Bazel MCP Server with various clients.

## Setup

First, configure your MCP client (e.g., Claude Desktop) to use the server:

```json
{
  "mcpServers": {
    "bazel": {
      "command": "/usr/local/bin/bazel-mcp-server"
    }
  }
}
```

## Example 1: Building a Target

**User Request:**
> Build the main application target

**Tool Call:**
```json
{
  "name": "bazel_build",
  "arguments": {
    "targets": ["//cmd/app:main"]
  }
}
```

**Response:**
```
Build successful!

INFO: Analyzed target //cmd/app:main (0 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //cmd/app:main up-to-date:
  bazel-bin/cmd/app/main
INFO: Build completed successfully, 1 total action
```

## Example 2: Running Tests with Filter

**User Request:**
> Run only the unit tests in the auth package

**Tool Call:**
```json
{
  "name": "bazel_test",
  "arguments": {
    "targets": ["//pkg/auth:auth_test"],
    "options": ["--test_output=errors"],
    "test_filter": "TestAuthentication*"
  }
}
```

**Response:**
```
Tests passed!

//pkg/auth:auth_test                                             PASSED in 0.5s

Executed 1 out of 1 test: 1 test passes.
```

## Example 3: Querying Dependencies

**User Request:**
> Show me all the dependencies of the main application

**Tool Call:**
```json
{
  "name": "bazel_query",
  "arguments": {
    "expression": "deps(//cmd/app:main)",
    "options": ["--output=label"]
  }
}
```

**Response:**
```
//cmd/app:main
//pkg/api:api
//pkg/auth:auth
//pkg/database:database
@go_sdk//go:go
@com_github_gin_gonic_gin//:gin
```

## Example 4: Finding Go Libraries

**User Request:**
> List all Go library targets in the project

**Tool Call:**
```json
{
  "name": "bazel_query",
  "arguments": {
    "expression": "kind(go_library, //...)",
    "options": ["--output=label"]
  }
}
```

**Response:**
```
//pkg/api:api
//pkg/auth:auth
//pkg/database:database
//pkg/models:models
//pkg/utils:utils
```

## Example 5: Running with Arguments

**User Request:**
> Start the server on port 9090 in debug mode

**Tool Call:**
```json
{
  "name": "bazel_run",
  "arguments": {
    "target": "//cmd/server:server",
    "args": ["--port=9090", "--debug"]
  }
}
```

**Response:**
```
Run completed!

INFO: Running command line: bazel-bin/cmd/server/server --port=9090 --debug
Starting server on port 9090...
Debug mode enabled
Server ready to accept connections
```

## Example 6: Cleaning Build Outputs

**User Request:**
> Clean all build outputs completely

**Tool Call:**
```json
{
  "name": "bazel_clean",
  "arguments": {
    "expunge": true
  }
}
```

**Response:**
```
Clean successful!

INFO: Starting clean (this may take a while).
```

## Example 7: Getting Build Information

**User Request:**
> Where are the build outputs stored?

**Tool Call:**
```json
{
  "name": "bazel_info",
  "arguments": {
    "keys": ["bazel-bin", "bazel-genfiles", "execution_root"]
  }
}
```

**Response:**
```
bazel-bin: /private/var/tmp/_bazel_user/abc123/execroot/myproject/bazel-out/darwin-fastbuild/bin
bazel-genfiles: /private/var/tmp/_bazel_user/abc123/execroot/myproject/bazel-out/darwin-fastbuild/bin
execution_root: /private/var/tmp/_bazel_user/abc123/execroot/myproject
```

## Example 8: Action Query for Compilation

**User Request:**
> Show me all the Go compilation actions for the main package

**Tool Call:**
```json
{
  "name": "bazel_aquery",
  "arguments": {
    "expression": "mnemonic('GoCompile', deps(//cmd/app:main))"
  }
}
```

**Response:**
```
action 'GoCompile cmd/app/main.a'
  Mnemonic: GoCompile
  Target: //cmd/app:main
  Configuration: darwin-fastbuild
  Execution platform: @local_config_platform//:host
  
action 'GoCompile pkg/api/api.a'
  Mnemonic: GoCompile
  Target: //pkg/api:api
  Configuration: darwin-fastbuild
  Execution platform: @local_config_platform//:host
```

## Example 9: Building with Configuration

**User Request:**
> Build the application with optimization enabled

**Tool Call:**
```json
{
  "name": "bazel_build",
  "arguments": {
    "targets": ["//..."],
    "options": ["--config=opt", "--compilation_mode=opt"]
  }
}
```

**Response:**
```
Build successful!

INFO: Build option --compilation_mode has changed, discarding analysis cache.
INFO: Analyzed 42 targets (15 packages loaded, 267 targets configured).
INFO: Found 42 targets...
INFO: Build completed successfully, 156 total actions
```

## Example 10: Running All Tests

**User Request:**
> Run all tests in the project with detailed output

**Tool Call:**
```json
{
  "name": "bazel_test",
  "arguments": {
    "targets": ["//..."],
    "options": ["--test_output=all", "--test_summary=detailed"]
  }
}
```

**Response:**
```
Tests passed!

//pkg/api:api_test                                               PASSED in 0.3s
//pkg/auth:auth_test                                             PASSED in 0.5s
//pkg/database:database_test                                     PASSED in 0.8s
//pkg/models:models_test                                         PASSED in 0.2s
//pkg/utils:utils_test                                           PASSED in 0.1s

Executed 5 out of 5 tests: 5 tests pass.
```

## Tips for Effective Use

### 1. Use Wildcards for Broader Operations
```json
{
  "name": "bazel_build",
  "arguments": {
    "targets": ["//pkg/..."]  // Builds all targets in pkg/
  }
}
```

### 2. Combine Query with Other Tools
First query to find targets, then build or test them:
```json
// Step 1: Find targets
{
  "name": "bazel_query",
  "arguments": {
    "expression": "kind(go_test, //pkg/...)"
  }
}

// Step 2: Test found targets
{
  "name": "bazel_test",
  "arguments": {
    "targets": ["//pkg/api:api_test", "//pkg/auth:auth_test"]
  }
}
```

### 3. Use Test Filters for Specific Tests
```json
{
  "name": "bazel_test",
  "arguments": {
    "targets": ["//pkg/..."],
    "test_filter": "TestIntegration*"  // Only integration tests
  }
}
```

### 4. Get Incremental Build Info
```json
{
  "name": "bazel_info",
  "arguments": {
    "keys": ["output_base", "workspace"]
  }
}
```

## Error Handling

The server returns errors as tool errors (not protocol errors) with helpful messages:

**Example Error Response:**
```
Build failed: exit status 1

Output:
ERROR: /path/to/BUILD:5:10: no such target '//pkg/nonexistent:target'
```

## Integration with AI Assistants

When using with AI assistants like Claude, you can ask natural language questions:

- "Build my project"
- "Run the tests for the authentication module"
- "What are the dependencies of the API package?"
- "Clean everything and rebuild from scratch"
- "Show me where Bazel stores its outputs"

The assistant will translate these into appropriate tool calls.
