package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/aws-support/mcp-server/config"
	"github.com/aws-support/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func DescribecasesHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["maxResults"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("maxResults=%v", val))
		}
		if val, ok := args["nextToken"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("nextToken=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.DescribeCasesRequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/#X-Amz-Target=AWSSupport_20130415.DescribeCases%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Target"]; ok {
			req.Header.Set("X-Amz-Target", fmt.Sprintf("%v", val))
		}

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
		var result models.DescribeCasesResponse
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

func CreateDescribecasesTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=AWSSupport_20130415_DescribeCases",
		mcp.WithDescription("<p>Returns a list of cases that you specify by passing one or more case IDs. You can use the <code>afterTime</code> and <code>beforeTime</code> parameters to filter the cases by date. You can set values for the <code>includeResolvedCases</code> and <code>includeCommunications</code> parameters to specify how much information to return.</p> <p>The response returns the following in JSON format:</p> <ul> <li> <p>One or more <a href="https://docs.aws.amazon.com/awssupport/latest/APIReference/API_CaseDetails.html">CaseDetails</a> data types.</p> </li> <li> <p>One or more <code>nextToken</code> values, which specify where to paginate the returned records represented by the <code>CaseDetails</code> objects.</p> </li> </ul> <p>Case data is available for 12 months after creation. If a case was created more than 12 months ago, a request might return an error.</p> <note> <ul> <li> <p>You must have a Business, Enterprise On-Ramp, or Enterprise Support plan to use the Amazon Web Services Support API. </p> </li> <li> <p>If you call the Amazon Web Services Support API from an account that doesn't have a Business, Enterprise On-Ramp, or Enterprise Support plan, the <code>SubscriptionRequiredException</code> error message appears. For information about changing your support plan, see <a href="http://aws.amazon.com/premiumsupport/">Amazon Web Services Support</a>.</p> </li> </ul> </note>"),
		mcp.WithString("maxResults", mcp.Description("Pagination limit")),
		mcp.WithString("nextToken", mcp.Description("Pagination token")),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("beforeTime", mcp.Description("")),
		mcp.WithString("includeCommunications", mcp.Description("")),
		mcp.WithString("includeResolvedCases", mcp.Description("")),
		mcp.WithString("nextToken", mcp.Description("")),
		mcp.WithString("afterTime", mcp.Description("")),
		mcp.WithString("caseIdList", mcp.Description("")),
		mcp.WithString("displayId", mcp.Description("")),
		mcp.WithString("language", mcp.Description("")),
		mcp.WithString("maxResults", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    DescribecasesHandler(cfg),
	}
}
