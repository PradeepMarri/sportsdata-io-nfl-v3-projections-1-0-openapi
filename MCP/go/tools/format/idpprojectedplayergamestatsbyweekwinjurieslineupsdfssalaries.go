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

func IdpprojectedplayergamestatsbyweekwinjurieslineupsdfssalariesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		seasonVal, ok := args["season"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: season"), nil
		}
		season, ok := seasonVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: season"), nil
		}
		weekVal, ok := args["week"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: week"), nil
		}
		week, ok := weekVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: week"), nil
		}
		url := fmt.Sprintf("%s/%s/IdpPlayerGameProjectionStatsByWeek/%s/%s", cfg.BaseURL, format, season, week)
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
		var result []PlayerGameProjection
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

func CreateIdpprojectedplayergamestatsbyweekwinjurieslineupsdfssalariesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_format_IdpPlayerGameProjectionStatsByWeek_season_week",
		mcp.WithDescription("IDP Projected Player Game Stats by Week (w/ Injuries, Lineups, DFS Salaries)"),
		mcp.WithString("format", mcp.Required(), mcp.Description("Desired response format. Valid entries are <code>XML</code> or <code>JSON</code>.")),
		mcp.WithString("season", mcp.Required(), mcp.Description("Year of the season and the season type. If no season type is provided, then the default is regular season.\n          <br>Examples: <code>2015REG</code>, <code>2015PRE</code>, <code>2015POST</code>.\n        ")),
		mcp.WithString("week", mcp.Required(), mcp.Description("Week of the season. Valid values are as follows: Preseason 0 to 4, Regular Season 1 to 17, Postseason 1 to 4.\n          Example: <code>1</code>\n        ")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    IdpprojectedplayergamestatsbyweekwinjurieslineupsdfssalariesHandler(cfg),
	}
}
