// Copyright 2025 The Bazel MCP Server Authors. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// TestBuildArgsValidation tests that BuildArgs validation works correctly
func TestBuildArgsValidation(t *testing.T) {
	// Test empty targets returns error
	args := BuildArgs{Targets: []string{}}
	result, _, err := bazelBuild(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if !result.IsError {
		t.Error("expected error result for empty targets")
	}

	// Valid arguments - we can't test actual execution without a Bazel workspace,
	// but we can verify the validation passes
	args = BuildArgs{Targets: []string{"//pkg:target"}}
	result, _, err = bazelBuild(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	// We expect it to fail because we're not in a Bazel workspace, but at least
	// it shouldn't be a validation error (IsError will be true from Bazel failure)
	if result != nil && len(result.Content) > 0 {
		text := result.Content[0].(*mcp.TextContent).Text
		if text == "Error: at least one target is required" {
			t.Error("validation should have passed for non-empty targets")
		}
	}
}

// TestTestArgsValidation tests that TestArgs validation works correctly
func TestTestArgsValidation(t *testing.T) {
	// Test empty targets returns error
	args := TestArgs{Targets: []string{}}
	result, _, err := bazelTest(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if !result.IsError {
		t.Error("expected error result for empty targets")
	}

	// Valid arguments
	args = TestArgs{Targets: []string{"//pkg:test"}, TestFilter: "TestMyFunction"}
	result, _, err = bazelTest(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if result != nil && len(result.Content) > 0 {
		text := result.Content[0].(*mcp.TextContent).Text
		if text == "Error: at least one test target is required" {
			t.Error("validation should have passed for non-empty targets")
		}
	}
}

// TestRunArgsValidation tests that RunArgs validation works correctly
func TestRunArgsValidation(t *testing.T) {
	// Test empty target returns error
	args := RunArgs{Target: ""}
	result, _, err := bazelRun(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if !result.IsError {
		t.Error("expected error result for empty target")
	}

	// Valid arguments
	args = RunArgs{Target: "//cmd/app:main", Args: []string{"--port=8080"}}
	result, _, err = bazelRun(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if result != nil && len(result.Content) > 0 {
		text := result.Content[0].(*mcp.TextContent).Text
		if text == "Error: target is required" {
			t.Error("validation should have passed for non-empty target")
		}
	}
}

// TestQueryArgsValidation tests that QueryArgs validation works correctly
func TestQueryArgsValidation(t *testing.T) {
	// Test empty expression returns error
	args := QueryArgs{Expression: ""}
	result, _, err := bazelQuery(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if !result.IsError {
		t.Error("expected error result for empty expression")
	}

	// Valid arguments
	args = QueryArgs{Expression: "deps(//pkg:target)", Options: []string{"--output=label"}}
	result, _, err = bazelQuery(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if result != nil && len(result.Content) > 0 {
		text := result.Content[0].(*mcp.TextContent).Text
		if text == "Error: query expression is required" {
			t.Error("validation should have passed for non-empty expression")
		}
	}
}

// TestAqueryArgsValidation tests that AqueryArgs validation works correctly
func TestAqueryArgsValidation(t *testing.T) {
	// Test empty expression returns error
	args := AqueryArgs{Expression: ""}
	result, _, err := bazelAquery(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if !result.IsError {
		t.Error("expected error result for empty expression")
	}

	// Valid arguments
	args = AqueryArgs{Expression: "deps(//pkg:target)", Options: []string{"--output=text"}}
	result, _, err = bazelAquery(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if result != nil && len(result.Content) > 0 {
		text := result.Content[0].(*mcp.TextContent).Text
		if text == "Error: aquery expression is required" {
			t.Error("validation should have passed for non-empty expression")
		}
	}
}

// TestCleanArgs tests that CleanArgs work correctly
func TestCleanArgs(t *testing.T) {
	// Clean always succeeds even with no arguments (validation-wise)
	args := CleanArgs{}
	result, _, err := bazelClean(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	// We expect an error because bazel isn't in a valid workspace, but that's OK for this test
	// We're just testing the argument handling
	if result == nil {
		t.Error("expected non-nil result")
	}
}

// TestInfoArgs tests that InfoArgs work correctly
func TestInfoArgs(t *testing.T) {
	// Info always succeeds even with no arguments (validation-wise)
	args := InfoArgs{}
	result, _, err := bazelInfo(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	// We expect an error because bazel isn't in a valid workspace, but that's OK for this test
	// We're just testing the argument handling
	if result == nil {
		t.Error("expected non-nil result")
	}
}

// TestModArgsValidation tests that ModArgs validation works correctly
func TestModArgsValidation(t *testing.T) {
	// Test empty subcommand returns error
	args := ModArgs{Subcommand: ""}
	result, _, err := bazelMod(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if !result.IsError {
		t.Error("expected error result for empty subcommand")
	}

	// Test invalid subcommand returns error
	args = ModArgs{Subcommand: "invalid_command"}
	result, _, err = bazelMod(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if !result.IsError {
		t.Error("expected error result for invalid subcommand")
	}

	// Valid arguments
	args = ModArgs{Subcommand: "graph"}
	result, _, err = bazelMod(context.Background(), &mcp.CallToolRequest{}, args)
	if err != nil {
		t.Fatalf("unexpected error from handler: %v", err)
	}
	if result != nil && len(result.Content) > 0 {
		text := result.Content[0].(*mcp.TextContent).Text
		if text == "Error: subcommand is required (graph, show_repo, show_extension, dump_repo_mapping)" {
			t.Error("validation should have passed for valid subcommand")
		}
		if text == "Error: invalid subcommand 'graph'. Valid options: graph, show_repo, show_extension, dump_repo_mapping" {
			t.Error("validation should have passed for 'graph' subcommand")
		}
	}
}

// TestJSONSerialization tests that all argument types can be properly serialized
func TestJSONSerialization(t *testing.T) {
	tests := []struct {
		name string
		args interface{}
	}{
		{"BuildArgs", BuildArgs{Targets: []string{"//..."}, Options: []string{"--config=opt"}}},
		{"TestArgs", TestArgs{Targets: []string{"//..."}, TestFilter: "Test*"}},
		{"CleanArgs", CleanArgs{Expunge: true, Async: false}},
		{"RunArgs", RunArgs{Target: "//cmd:main", Args: []string{"--help"}}},
		{"QueryArgs", QueryArgs{Expression: "deps(//...)", Options: []string{"--output=label"}}},
		{"AqueryArgs", AqueryArgs{Expression: "deps(//...)", Options: []string{"--output=text"}}},
		{"InfoArgs", InfoArgs{Keys: []string{"bazel-bin", "bazel-genfiles"}}},
		{"ModArgs", ModArgs{Subcommand: "graph", Args: []string{}, Options: []string{}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Test that we can marshal and unmarshal
			data, err := json.Marshal(tt.args)
			if err != nil {
				t.Fatalf("failed to marshal: %v", err)
			}

			// Unmarshal back to map to verify structure
			var m map[string]interface{}
			if err := json.Unmarshal(data, &m); err != nil {
				t.Fatalf("failed to unmarshal: %v", err)
			}

			// Verify we got some data
			if len(m) == 0 {
				t.Error("expected non-empty map after unmarshaling")
			}
		})
	}
}
