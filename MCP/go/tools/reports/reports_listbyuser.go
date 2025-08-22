package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/apimanagementclient/mcp-server/config"
	"github.com/apimanagementclient/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Reports_listbyuserHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["$filter"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("$filter=%v", val))
		}
		if val, ok := args["$top"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("$top=%v", val))
		}
		if val, ok := args["$skip"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("$skip=%v", val))
		}
		if val, ok := args["api-version"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("api-version=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/reports/byUser%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("Authorization", cfg.APIKey)
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
		var result models.ReportCollection
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

func CreateReports_listbyuserTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_reports_byUser",
		mcp.WithDescription("Lists report records by User."),
		mcp.WithString("$filter", mcp.Required(), mcp.Description("The filter to apply on the operation.")),
		mcp.WithNumber("$top", mcp.Description("Number of records to return.")),
		mcp.WithNumber("$skip", mcp.Description("Number of records to skip.")),
		mcp.WithString("api-version", mcp.Required(), mcp.Description("Version of the API to be used with the client request.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Reports_listbyuserHandler(cfg),
	}
}
