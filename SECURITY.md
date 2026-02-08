# Security Policy

## Supported Versions

We release patches for security vulnerabilities for the following versions:

| Version | Supported          |
| ------- | ------------------ |
| 0.1.x   | :white_check_mark: |

## Reporting a Vulnerability

If you discover a security vulnerability within Bazel MCP Server, please send an email to lewissetter@gmail.com. All security vulnerabilities will be promptly addressed.

**Please do not report security vulnerabilities through public GitHub issues.**

### What to Include

When reporting a vulnerability, please include:

- Type of issue (e.g., buffer overflow, SQL injection, cross-site scripting, etc.)
- Full paths of source file(s) related to the manifestation of the issue
- The location of the affected source code (tag/branch/commit or direct URL)
- Any special configuration required to reproduce the issue
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the issue, including how an attacker might exploit the issue

### Response Timeline

- We will acknowledge receipt of your vulnerability report within 3 business days
- We will provide a more detailed response within 7 business days
- We will work on a fix and coordinate the release of any necessary patches

## Security Best Practices

When using Bazel MCP Server:

1. **Keep Dependencies Updated**: Regularly update to the latest version to receive security patches
2. **Limit Permissions**: Run the server with the minimum necessary permissions
3. **Validate Input**: The server validates Bazel commands, but be aware of command injection risks
4. **Secure Configuration**: Ensure your MCP client configuration is properly secured
5. **Monitor Activity**: Keep logs and monitor for unusual Bazel command patterns

## Known Security Considerations

### Command Execution

The server executes Bazel commands on your system. Ensure:
- The server runs in a trusted environment
- Input validation is in place (which it is)
- Access is restricted to authorized MCP clients only

### File System Access

The server operates in the current working directory and has access to:
- Bazel workspace files
- Build outputs
- Source code

Ensure the server runs with appropriate file system permissions.

## Updates

Security updates will be released as soon as possible after a vulnerability is discovered and verified. Updates will be announced:

- In the CHANGELOG.md file
- As GitHub releases
- Through security advisories on the GitHub repository
