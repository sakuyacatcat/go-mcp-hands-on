package server

import (
	"context"
	"fmt"

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

	// HINT: use mcp.WithString and mcp.WithNumber to define the parameters.
	// HINT: use mcp.DefaultNumber to set the default value of the number parameter.

	// parameters:
	// * text: 文字列
	// 	- 必須
	// 	- 説明: "Text to rotate"
	// * shift: 数値
	// 	- デフォルト値: 13
	// 	- 説明: "Number of positions to rotate"

	// register a handler into the tool
	return s
}

func rotateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// TODO: implement
	// HINT: use request.Params.Arguments to get the parameters. then, use type assertion to convert them to the correct type like `.(string)`.
	// HINT: "shift" is a number, so use `.(float64)` to convert it.

	// HINT: use caesar.RotN

	// TODO: input the result into `fmt.Sprint()`
	return mcp.NewToolResultText(fmt.Sprint()), nil
}
