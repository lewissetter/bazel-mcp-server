# Bazel MCP Server - Project Overview

## What Is This?

Bazel MCP Server is a complete, production-ready implementation of a Model Context Protocol (MCP) server for the Bazel build system. It enables AI assistants (like Claude) and other MCP clients to interact with Bazel projects through natural language commands.

## Project Status

✅ **Ready for GitHub** - Fully configured and committed to git
✅ **Production Ready** - Comprehensive tests and error handling
✅ **CI/CD Ready** - GitHub Actions configured for testing and releases
✅ **Well Documented** - Complete documentation and examples

## What's Included

### Core Implementation
- **7 Bazel Tools**: build, test, clean, run, query, aquery, info
- **Type-Safe API**: Strongly-typed arguments with JSON schema validation
- **Error Handling**: User-friendly error messages
- **Context Support**: Cancellation support via context.Context
- **Comprehensive Tests**: Full test coverage with table-driven tests

### CI/CD & Automation
- **GitHub Actions CI**: Tests on Linux, macOS, and Windows with Go 1.21-1.23
- **GitHub Actions Release**: Automated releases with GoReleaser
- **Multi-Platform Builds**: Automatic binaries for all major platforms
- **Dependabot**: Automated dependency updates
- **Linting**: golangci-lint configuration

### Documentation
- **README.md**: Complete project overview and usage guide
- **CONTRIBUTING.md**: Contribution guidelines
- **SECURITY.md**: Security policy and best practices
- **CODE_OF_CONDUCT.md**: Community guidelines
- **CHANGELOG.md**: Version history
- **EXAMPLES.md**: Detailed usage examples with 10+ scenarios

### Developer Experience
- **Makefile**: Convenient build, test, clean, and run commands
- **Dockerfile**: Container support for Docker deployments
- **Issue Templates**: Bug report and feature request templates
- **PR Template**: Pull request template with checklist

## Directory Structure

```
bazel-mcp-server/
├── cmd/
│   └── bazel-mcp-server/        # Main application
│       ├── main.go              # Server implementation (10.5KB)
│       └── main_test.go         # Test suite (7.3KB)
├── examples/
│   ├── EXAMPLES.md              # Detailed usage examples
│   └── claude_desktop_config.json # Sample configuration
├── .github/
│   ├── workflows/               # CI/CD pipelines
│   │   ├── ci.yml              # Test workflow
│   │   └── release.yml         # Release workflow
│   ├── ISSUE_TEMPLATE/         # Issue templates
│   │   ├── bug_report.yml
│   │   └── feature_request.yml
│   ├── PULL_REQUEST_TEMPLATE.md
│   └── dependabot.yml          # Dependency updates
├── .gitignore                   # Git ignore rules
├── .golangci.yml                # Linter configuration
├── .goreleaser.yml              # Release configuration
├── .dockerignore                # Docker ignore rules
├── Dockerfile                   # Container definition
├── Makefile                     # Build automation
├── go.mod                       # Go module definition
├── go.sum                       # Go module checksums
├── LICENSE                      # MIT License
├── README.md                    # Project documentation (283 lines)
├── CONTRIBUTING.md              # Contribution guidelines (152 lines)
├── SECURITY.md                  # Security policy (69 lines)
├── CODE_OF_CONDUCT.md          # Code of conduct (128 lines)
└── CHANGELOG.md                 # Version history

Total: 25 files, ~2,150 lines of content
```

## Key Features

### 1. Complete Bazel Integration
- All major Bazel commands supported
- Options and flags fully configurable
- Multiple targets support
- Test filtering capabilities

### 2. Production Quality
- Comprehensive error handling
- Input validation
- Type-safe APIs
- Context cancellation support
- Combined stdout/stderr output

### 3. Developer Friendly
- Clear documentation
- Example configurations
- Detailed usage examples
- Easy to extend

### 4. GitHub Ready
- Complete CI/CD setup
- Automated releases
- Issue and PR templates
- Code of conduct and security policy
- Automated dependency updates

## Quick Start

### For Users

1. **Install from source:**
   ```bash
   go install github.com/lewissetter/bazel-mcp-server/cmd/bazel-mcp-server@latest
   ```

2. **Configure Claude Desktop** (or other MCP client):
   ```json
   {
     "mcpServers": {
       "bazel": {
         "command": "bazel-mcp-server"
       }
     }
   }
   ```

3. **Use natural language commands:**
   - "Build my project"
   - "Run the tests"
   - "What are the dependencies?"

### For Developers

1. **Clone and build:**
   ```bash
   cd bazel-mcp-server
   make build
   ```

2. **Run tests:**
   ```bash
   make test
   ```

3. **Install locally:**
   ```bash
   make install
   ```

## Publishing to GitHub

The project is ready to be pushed to GitHub:

1. **Create a new GitHub repository** named `bazel-mcp-server`

2. **Add remote and push:**
   ```bash
   git remote add origin https://github.com/lewissetter/bazel-mcp-server.git
   git push -u origin main
   ```

3. **Create first release:**
   ```bash
   git tag -a v0.1.0 -m "Release v0.1.0"
   git push origin v0.1.0
   ```

The GitHub Actions workflow will automatically:
- Run tests on multiple platforms
- Build binaries for Linux, macOS, and Windows
- Create a GitHub release
- Upload release artifacts

## Technologies Used

- **Go 1.21+**: Modern Go with generics support
- **MCP Go SDK**: Official Model Context Protocol implementation
- **GitHub Actions**: CI/CD automation
- **GoReleaser**: Multi-platform releases
- **golangci-lint**: Code quality
- **Docker**: Container support

## Next Steps

### Future Enhancements
- [ ] Add more Bazel commands (coverage, mobile-install, etc.)
- [ ] Add workspace configuration support
- [ ] Implement caching for query results
- [ ] Add progress streaming for long builds
- [ ] Support for Bazel remote execution
- [ ] Integration tests with real Bazel projects
- [ ] Performance benchmarks

## Support & Community

Once published on GitHub, users can:
- 🐛 Report bugs via GitHub Issues
- ✨ Request features via GitHub Issues
- 💬 Ask questions via GitHub Discussions
- 🤝 Contribute via Pull Requests

## License

MIT License - See LICENSE file for details

## Acknowledgments

- Built with [Model Context Protocol Go SDK](https://github.com/modelcontextprotocol/go-sdk)
- Inspired by the need to integrate AI assistants with Bazel workflows
- Follows Go community best practices

---

**Project created with [Continue](https://continue.dev)**
