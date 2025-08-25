package main

import (
	"github.com/nfl-v3-projections/mcp-server/config"
	"github.com/nfl-v3-projections/mcp-server/models"
	tools_format "github.com/nfl-v3-projections/mcp-server/tools/format"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_format.CreateProjectedfantasydefenseseasonstatswadpTool(cfg),
		tools_format.CreateProjectedplayergamestatsbyteamwinjurieslineupsdfssalariesTool(cfg),
		tools_format.CreateDfsslatesbyweekTool(cfg),
		tools_format.CreateProjectedplayerseasonstatswadpTool(cfg),
		tools_format.CreateProjectedplayerseasonstatsbyplayerwadpTool(cfg),
		tools_format.CreateDfsslatesbydateTool(cfg),
		tools_format.CreateIdpprojectedplayergamestatsbyweekwinjurieslineupsdfssalariesTool(cfg),
		tools_format.CreateProjectedplayergamestatsbyplayerwinjurieslineupsdfssalariesTool(cfg),
		tools_format.CreateProjectedplayerseasonstatsbyteamwadpTool(cfg),
		tools_format.CreateIdpprojectedplayergamestatsbyplayerwinjurieslineupsdfssalariesTool(cfg),
		tools_format.CreateInjuredplayersTool(cfg),
		tools_format.CreateDfsslateownershipprojectionsbyslateidTool(cfg),
		tools_format.CreateIdpprojectedplayergamestatsbyteamwinjurieslineupsdfssalariesTool(cfg),
		tools_format.CreateUpcomingdfsslateownershipprojectionsTool(cfg),
		tools_format.CreateProjectedfantasydefensegamestatswdfssalariesTool(cfg),
		tools_format.CreateProjectedplayergamestatsbyweekwinjurieslineupsdfssalariesTool(cfg),
	}
}
