package server

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func New() *server.MCPServer {
	s := server.NewMCPServer(
		"caesar Cipher",
		"1.0.0",
	)
	// Add tool here
	// name: "caesar_rotate"
	// description: "Rotate a string by a given number of positions. It is used to encrypt or decrypt text of caesar Cipher."

	// register a handler into the tool
	return s
}

func rotateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// TODO: implement
	// HINT: use caesar.RotN

	// TODO: return the result. the below code is just a placeholder.
	return nil, nil
}
