package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nfl-v3-projections/mcp-server/config"
	"github.com/nfl-v3-projections/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DfsslatesbydateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		formatVal, ok := args["format"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: format"), nil
		}
		format, ok := formatVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: format"), nil
		}
		dateVal, ok := args["date"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: date"), nil
		}
		date, ok := dateVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: date"), nil
		}
		url := fmt.Sprintf("%s/%s/DfsSlatesByDate/%s", cfg.BaseURL, format, date)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("Ocp-Apim-Subscription-Key", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result []DfsSlate
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateDfsslatesbydateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_format_DfsSlatesByDate_date",
		mcp.WithDescription("DFS Slates by Date"),
		mcp.WithString("format", mcp.Required(), mcp.Description("Desired response format. Valid entries are <code>XML</code> or <code>JSON</code>.")),
		mcp.WithString("date", mcp.Required(), mcp.Description("The date of the slates.\n<br>Examples: <code>2017-SEP-25</code>, <code>2017-10-31</code>.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DfsslatesbydateHandler(cfg),
	}
}
