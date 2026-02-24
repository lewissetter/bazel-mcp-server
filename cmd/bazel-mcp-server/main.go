// Copyright 2025 The Bazel MCP Server Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

// The bazel-mcp-server provides MCP tools for interacting with Bazel build system.
//
// It runs over the stdio transport and provides tools for common Bazel commands:
// build, test, clean, run, query, aquery, info, and mod.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const version = "0.1.0"

func main() {
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "bazel-mcp-server",
		Version: version,
	}, &mcp.ServerOptions{
		Instructions: "Use this server to interact with Bazel build system. All commands are executed in the current working directory.",
	})

	// Add Bazel command tools
	mcp.AddTool(server, &mcp.Tool{
		Name:        "bazel_build",
		Description: "Build one or more targets using Bazel",
	}, bazelBuild)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bazel_test",
		Description: "Run tests for one or more targets using Bazel",
	}, bazelTest)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bazel_clean",
		Description: "Clean Bazel build outputs",
	}, bazelClean)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bazel_run",
		Description: "Build and run a single target using Bazel",
	}, bazelRun)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bazel_query",
		Description: "Query the Bazel dependency graph",
	}, bazelQuery)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bazel_aquery",
		Description: "Query information about actions in the build graph (action graph query)",
	}, bazelAquery)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bazel_info",
		Description: "Display runtime info about the Bazel server",
	}, bazelInfo)

	mcp.AddTool(server, &mcp.Tool{
		Name:        "bazel_mod",
		Description: "Manage Bazel modules (bzlmod) - query module graph, show repositories, and inspect extensions",
	}, bazelMod)

	// Run the server over stdio
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Printf("Server failed: %v", err)
	}
}

// BuildArgs represents arguments for bazel build command
type BuildArgs struct {
	Targets []string `json:"targets" jsonschema:"list of targets to build (e.g. //pkg:target or //...)"`
	Options []string `json:"options,omitempty" jsonschema:"additional build options (e.g. --config=opt --compilation_mode=dbg)"`
}

// TestArgs represents arguments for bazel test command
type TestArgs struct {
	Targets    []string `json:"targets" jsonschema:"list of test targets to run (e.g. //pkg:test or //...)"`
	Options    []string `json:"options,omitempty" jsonschema:"additional test options (e.g. --test_output=all --test_filter=TestName)"`
	TestArgs   []string `json:"test_args,omitempty" jsonschema:"arguments to pass to the test binary (e.g. --verbose)"`
	TestFilter string   `json:"test_filter,omitempty" jsonschema:"filter to select specific tests to run"`
}

// CleanArgs represents arguments for bazel clean command
type CleanArgs struct {
	Expunge bool     `json:"expunge,omitempty" jsonschema:"remove the entire working tree for this Bazel instance"`
	Async   bool     `json:"async,omitempty" jsonschema:"clean asynchronously"`
	Options []string `json:"options,omitempty" jsonschema:"additional clean options"`
}

// RunArgs represents arguments for bazel run command
type RunArgs struct {
	Target  string   `json:"target" jsonschema:"single target to build and run (e.g. //cmd/app:main)"`
	Options []string `json:"options,omitempty" jsonschema:"additional build/run options"`
	Args    []string `json:"args,omitempty" jsonschema:"arguments to pass to the target binary"`
}

// QueryArgs represents arguments for bazel query command
type QueryArgs struct {
	Expression string   `json:"expression" jsonschema:"query expression (e.g. deps(//pkg:target) or kind(go_library //...))"`
	Options    []string `json:"options,omitempty" jsonschema:"additional query options (e.g. --output=label --output=graph)"`
}

// AqueryArgs represents arguments for bazel aquery command
type AqueryArgs struct {
	Expression string   `json:"expression" jsonschema:"action query expression (e.g. outputs with deps function)"`
	Options    []string `json:"options,omitempty" jsonschema:"additional aquery options (e.g. --output=text --output=jsonproto)"`
}

// InfoArgs represents arguments for bazel info command
type InfoArgs struct {
	Keys    []string `json:"keys,omitempty" jsonschema:"specific info keys to retrieve (e.g. bazel-bin bazel-genfiles execution_root)"`
	Options []string `json:"options,omitempty" jsonschema:"additional info options"`
}

// ModArgs represents arguments for bazel mod command
type ModArgs struct {
	Subcommand string   `json:"subcommand" jsonschema:"mod subcommand to run (graph, show_repo, show_extension, dump_repo_mapping)"`
	Args       []string `json:"args,omitempty" jsonschema:"arguments for the subcommand (e.g. repository name for show_repo)"`
	Options    []string `json:"options,omitempty" jsonschema:"additional mod options (e.g. --base_module for graph, --extension_filter for show_extension)"`
}

// bazelBuild builds one or more targets
func bazelBuild(ctx context.Context, req *mcp.CallToolRequest, args BuildArgs) (*mcp.CallToolResult, any, error) {
	if len(args.Targets) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: at least one target is required"}},
			IsError: true,
		}, nil, nil
	}

	cmdArgs := append([]string{"build"}, args.Options...)
	cmdArgs = append(cmdArgs, args.Targets...)

	output, err := executeBazel(ctx, cmdArgs)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Build failed: %v\n\nOutput:\n%s", err, output)}},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Build successful!\n\n%s", output)}},
	}, nil, nil
}

// bazelTest runs tests for one or more targets
func bazelTest(ctx context.Context, req *mcp.CallToolRequest, args TestArgs) (*mcp.CallToolResult, any, error) {
	if len(args.Targets) == 0 {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: at least one test target is required"}},
			IsError: true,
		}, nil, nil
	}

	cmdArgs := append([]string{"test"}, args.Options...)

	if args.TestFilter != "" {
		cmdArgs = append(cmdArgs, fmt.Sprintf("--test_filter=%s", args.TestFilter))
	}

	cmdArgs = append(cmdArgs, args.Targets...)

	if len(args.TestArgs) > 0 {
		cmdArgs = append(cmdArgs, "--")
		cmdArgs = append(cmdArgs, args.TestArgs...)
	}

	output, err := executeBazel(ctx, cmdArgs)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Tests failed: %v\n\nOutput:\n%s", err, output)}},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Tests passed!\n\n%s", output)}},
	}, nil, nil
}

// bazelClean cleans build outputs
func bazelClean(ctx context.Context, req *mcp.CallToolRequest, args CleanArgs) (*mcp.CallToolResult, any, error) {
	cmdArgs := []string{"clean"}

	if args.Expunge {
		cmdArgs = append(cmdArgs, "--expunge")
	}
	if args.Async {
		cmdArgs = append(cmdArgs, "--async")
	}
	cmdArgs = append(cmdArgs, args.Options...)

	output, err := executeBazel(ctx, cmdArgs)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Clean failed: %v\n\nOutput:\n%s", err, output)}},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Clean successful!\n\n%s", output)}},
	}, nil, nil
}

// bazelRun builds and runs a target
func bazelRun(ctx context.Context, req *mcp.CallToolRequest, args RunArgs) (*mcp.CallToolResult, any, error) {
	if args.Target == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: target is required"}},
			IsError: true,
		}, nil, nil
	}

	cmdArgs := append([]string{"run"}, args.Options...)
	cmdArgs = append(cmdArgs, args.Target)

	if len(args.Args) > 0 {
		cmdArgs = append(cmdArgs, "--")
		cmdArgs = append(cmdArgs, args.Args...)
	}

	output, err := executeBazel(ctx, cmdArgs)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Run failed: %v\n\nOutput:\n%s", err, output)}},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Run completed!\n\n%s", output)}},
	}, nil, nil
}

// bazelQuery performs a query on the dependency graph
func bazelQuery(ctx context.Context, req *mcp.CallToolRequest, args QueryArgs) (*mcp.CallToolResult, any, error) {
	if args.Expression == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: query expression is required"}},
			IsError: true,
		}, nil, nil
	}

	cmdArgs := append([]string{"query"}, args.Options...)
	cmdArgs = append(cmdArgs, args.Expression)

	output, err := executeBazel(ctx, cmdArgs)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Query failed: %v\n\nOutput:\n%s", err, output)}},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: output}},
	}, nil, nil
}

// bazelAquery performs an action query on the build graph
func bazelAquery(ctx context.Context, req *mcp.CallToolRequest, args AqueryArgs) (*mcp.CallToolResult, any, error) {
	if args.Expression == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: aquery expression is required"}},
			IsError: true,
		}, nil, nil
	}

	cmdArgs := append([]string{"aquery"}, args.Options...)
	cmdArgs = append(cmdArgs, args.Expression)

	output, err := executeBazel(ctx, cmdArgs)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Aquery failed: %v\n\nOutput:\n%s", err, output)}},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: output}},
	}, nil, nil
}

// bazelInfo displays runtime info about the Bazel server
func bazelInfo(ctx context.Context, req *mcp.CallToolRequest, args InfoArgs) (*mcp.CallToolResult, any, error) {
	cmdArgs := append([]string{"info"}, args.Options...)
	cmdArgs = append(cmdArgs, args.Keys...)

	output, err := executeBazel(ctx, cmdArgs)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Info failed: %v\n\nOutput:\n%s", err, output)}},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: output}},
	}, nil, nil
}

// bazelMod manages Bazel modules (bzlmod)
func bazelMod(ctx context.Context, req *mcp.CallToolRequest, args ModArgs) (*mcp.CallToolResult, any, error) {
	if args.Subcommand == "" {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: "Error: subcommand is required (graph, show_repo, show_extension, dump_repo_mapping)"}},
			IsError: true,
		}, nil, nil
	}

	// Validate subcommand
	validSubcommands := map[string]bool{
		"graph":             true,
		"show_repo":         true,
		"show_extension":    true,
		"dump_repo_mapping": true,
	}
	if !validSubcommands[args.Subcommand] {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Error: invalid subcommand '%s'. Valid options: graph, show_repo, show_extension, dump_repo_mapping", args.Subcommand)}},
			IsError: true,
		}, nil, nil
	}

	cmdArgs := []string{"mod", args.Subcommand}
	cmdArgs = append(cmdArgs, args.Options...)
	cmdArgs = append(cmdArgs, args.Args...)

	output, err := executeBazel(ctx, cmdArgs)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("Mod command failed: %v\n\nOutput:\n%s", err, output)}},
			IsError: true,
		}, nil, nil
	}

	return &mcp.CallToolResult{
		Content: []mcp.Content{&mcp.TextContent{Text: output}},
	}, nil, nil
}

// executeBazel executes a bazel command and returns its output
func executeBazel(ctx context.Context, args []string) (string, error) {
	cmd := exec.CommandContext(ctx, "bazel", args...)
	cmd.Dir, _ = os.Getwd()

	output, err := cmd.CombinedOutput()
	outputStr := strings.TrimSpace(string(output))

	return outputStr, err
}
