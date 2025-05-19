package server

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/tenkoh/go-mcp-hands-on/caesar-mcp/caesar"
)

func New() *server.MCPServer {
	s := server.NewMCPServer(
		"caesar Cipher",
		"1.0.0",
	)

	tool := mcp.NewTool("caesar_rotate",
		mcp.WithDescription("Rotate a string by a given number of positions. It is used to encrypt or decrypt text of caesar Cipher."),
		mcp.WithString("text",
			mcp.Description("Text to rotate"),
			mcp.Required(),
		),
		mcp.WithNumber("shift",
			mcp.Description("Number of positions to rotate"),
			mcp.DefaultNumber(13),
		),
	)

	s.AddTool(tool, rotateHandler)
	return s
}

func rotateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	text, ok := request.Params.Arguments["text"].(string)
	if !ok {
		return nil, fmt.Errorf("text must be a string")
	}

	var shift int = 13
	if s, ok := request.Params.Arguments["shift"]; ok {
		if f, ok := s.(float64); ok {
			shift = int(f)
		} else {
			return nil, fmt.Errorf("shift must be a number")
		}
	}

	result := caesar.RotN(text, shift)
	return mcp.NewToolResultText(fmt.Sprint(result)), nil
}
