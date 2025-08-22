package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// DescribeServicesRequest represents the DescribeServicesRequest schema from the OpenAPI specification
type DescribeServicesRequest struct {
	Language interface{} `json:"language,omitempty"`
	Servicecodelist interface{} `json:"serviceCodeList,omitempty"`
}

// DescribeTrustedAdvisorCheckRefreshStatusesRequest represents the DescribeTrustedAdvisorCheckRefreshStatusesRequest schema from the OpenAPI specification
type DescribeTrustedAdvisorCheckRefreshStatusesRequest struct {
	Checkids interface{} `json:"checkIds"`
}

// DescribeCasesRequest represents the DescribeCasesRequest schema from the OpenAPI specification
type DescribeCasesRequest struct {
	Language interface{} `json:"language,omitempty"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Beforetime interface{} `json:"beforeTime,omitempty"`
	Includecommunications interface{} `json:"includeCommunications,omitempty"`
	Includeresolvedcases interface{} `json:"includeResolvedCases,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
	Aftertime interface{} `json:"afterTime,omitempty"`
	Caseidlist interface{} `json:"caseIdList,omitempty"`
	Displayid interface{} `json:"displayId,omitempty"`
}

// DescribeTrustedAdvisorCheckRefreshStatusesResponse represents the DescribeTrustedAdvisorCheckRefreshStatusesResponse schema from the OpenAPI specification
type DescribeTrustedAdvisorCheckRefreshStatusesResponse struct {
	Statuses interface{} `json:"statuses"`
}

// CreateCaseResponse represents the CreateCaseResponse schema from the OpenAPI specification
type CreateCaseResponse struct {
	Caseid interface{} `json:"caseId,omitempty"`
}

// AddCommunicationToCaseResponse represents the AddCommunicationToCaseResponse schema from the OpenAPI specification
type AddCommunicationToCaseResponse struct {
	Result interface{} `json:"result,omitempty"`
}

// ResolveCaseRequest represents the ResolveCaseRequest schema from the OpenAPI specification
type ResolveCaseRequest struct {
	Caseid interface{} `json:"caseId,omitempty"`
}

// AttachmentDetails represents the AttachmentDetails schema from the OpenAPI specification
type AttachmentDetails struct {
	Attachmentid interface{} `json:"attachmentId,omitempty"`
	Filename interface{} `json:"fileName,omitempty"`
}

// DescribeTrustedAdvisorCheckSummariesResponse represents the DescribeTrustedAdvisorCheckSummariesResponse schema from the OpenAPI specification
type DescribeTrustedAdvisorCheckSummariesResponse struct {
	Summaries interface{} `json:"summaries"`
}

// TrustedAdvisorCheckSummary represents the TrustedAdvisorCheckSummary schema from the OpenAPI specification
type TrustedAdvisorCheckSummary struct {
	Categoryspecificsummary interface{} `json:"categorySpecificSummary"`
	Checkid interface{} `json:"checkId"`
	Hasflaggedresources interface{} `json:"hasFlaggedResources,omitempty"`
	Resourcessummary TrustedAdvisorResourcesSummary `json:"resourcesSummary"` // Details about Amazon Web Services resources that were analyzed in a call to Trusted Advisor <a>DescribeTrustedAdvisorCheckSummaries</a>.
	Status interface{} `json:"status"`
	Timestamp interface{} `json:"timestamp"`
}

// CaseDetails represents the CaseDetails schema from the OpenAPI specification
type CaseDetails struct {
	Subject interface{} `json:"subject,omitempty"`
	Ccemailaddresses interface{} `json:"ccEmailAddresses,omitempty"`
	Displayid interface{} `json:"displayId,omitempty"`
	Servicecode interface{} `json:"serviceCode,omitempty"`
	Severitycode interface{} `json:"severityCode,omitempty"`
	Categorycode interface{} `json:"categoryCode,omitempty"`
	Timecreated interface{} `json:"timeCreated,omitempty"`
	Language interface{} `json:"language,omitempty"`
	Recentcommunications interface{} `json:"recentCommunications,omitempty"`
	Status interface{} `json:"status,omitempty"`
	Submittedby interface{} `json:"submittedBy,omitempty"`
	Caseid interface{} `json:"caseId,omitempty"`
}

// AddAttachmentsToSetRequest represents the AddAttachmentsToSetRequest schema from the OpenAPI specification
type AddAttachmentsToSetRequest struct {
	Attachments interface{} `json:"attachments"`
	Attachmentsetid interface{} `json:"attachmentSetId,omitempty"`
}

// SeverityLevel represents the SeverityLevel schema from the OpenAPI specification
type SeverityLevel struct {
	Name interface{} `json:"name,omitempty"`
	Code interface{} `json:"code,omitempty"`
}

// RefreshTrustedAdvisorCheckResponse represents the RefreshTrustedAdvisorCheckResponse schema from the OpenAPI specification
type RefreshTrustedAdvisorCheckResponse struct {
	Status interface{} `json:"status"`
}

// DescribeAttachmentRequest represents the DescribeAttachmentRequest schema from the OpenAPI specification
type DescribeAttachmentRequest struct {
	Attachmentid interface{} `json:"attachmentId"`
}

// DescribeCommunicationsRequest represents the DescribeCommunicationsRequest schema from the OpenAPI specification
type DescribeCommunicationsRequest struct {
	Aftertime interface{} `json:"afterTime,omitempty"`
	Beforetime interface{} `json:"beforeTime,omitempty"`
	Caseid interface{} `json:"caseId"`
	Maxresults interface{} `json:"maxResults,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// TrustedAdvisorCategorySpecificSummary represents the TrustedAdvisorCategorySpecificSummary schema from the OpenAPI specification
type TrustedAdvisorCategorySpecificSummary struct {
	Costoptimizing interface{} `json:"costOptimizing,omitempty"`
}

// DescribeAttachmentResponse represents the DescribeAttachmentResponse schema from the OpenAPI specification
type DescribeAttachmentResponse struct {
	Attachment interface{} `json:"attachment,omitempty"`
}

// DescribeTrustedAdvisorCheckResultResponse represents the DescribeTrustedAdvisorCheckResultResponse schema from the OpenAPI specification
type DescribeTrustedAdvisorCheckResultResponse struct {
	Result interface{} `json:"result,omitempty"`
}

// DescribeTrustedAdvisorChecksRequest represents the DescribeTrustedAdvisorChecksRequest schema from the OpenAPI specification
type DescribeTrustedAdvisorChecksRequest struct {
	Language interface{} `json:"language"`
}

// Communication represents the Communication schema from the OpenAPI specification
type Communication struct {
	Attachmentset interface{} `json:"attachmentSet,omitempty"`
	Body interface{} `json:"body,omitempty"`
	Caseid interface{} `json:"caseId,omitempty"`
	Submittedby interface{} `json:"submittedBy,omitempty"`
	Timecreated interface{} `json:"timeCreated,omitempty"`
}

// DescribeServicesResponse represents the DescribeServicesResponse schema from the OpenAPI specification
type DescribeServicesResponse struct {
	Services interface{} `json:"services,omitempty"`
}

// ResolveCaseResponse represents the ResolveCaseResponse schema from the OpenAPI specification
type ResolveCaseResponse struct {
	Finalcasestatus interface{} `json:"finalCaseStatus,omitempty"`
	Initialcasestatus interface{} `json:"initialCaseStatus,omitempty"`
}

// DescribeCommunicationsResponse represents the DescribeCommunicationsResponse schema from the OpenAPI specification
type DescribeCommunicationsResponse struct {
	Communications interface{} `json:"communications,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DescribeSupportedLanguagesResponse represents the DescribeSupportedLanguagesResponse schema from the OpenAPI specification
type DescribeSupportedLanguagesResponse struct {
	Supportedlanguages interface{} `json:"supportedLanguages,omitempty"`
}

// AddAttachmentsToSetResponse represents the AddAttachmentsToSetResponse schema from the OpenAPI specification
type AddAttachmentsToSetResponse struct {
	Attachmentsetid interface{} `json:"attachmentSetId,omitempty"`
	Expirytime interface{} `json:"expiryTime,omitempty"`
}

// DescribeCasesResponse represents the DescribeCasesResponse schema from the OpenAPI specification
type DescribeCasesResponse struct {
	Cases interface{} `json:"cases,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// TrustedAdvisorCheckDescription represents the TrustedAdvisorCheckDescription schema from the OpenAPI specification
type TrustedAdvisorCheckDescription struct {
	Name interface{} `json:"name"`
	Category interface{} `json:"category"`
	Description interface{} `json:"description"`
	Id interface{} `json:"id"`
	Metadata interface{} `json:"metadata"`
}

// Service represents the Service schema from the OpenAPI specification
type Service struct {
	Categories interface{} `json:"categories,omitempty"`
	Code interface{} `json:"code,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// Attachment represents the Attachment schema from the OpenAPI specification
type Attachment struct {
	Data interface{} `json:"data,omitempty"`
	Filename interface{} `json:"fileName,omitempty"`
}

// CreateCaseRequest represents the CreateCaseRequest schema from the OpenAPI specification
type CreateCaseRequest struct {
	Attachmentsetid interface{} `json:"attachmentSetId,omitempty"`
	Communicationbody interface{} `json:"communicationBody"`
	Ccemailaddresses interface{} `json:"ccEmailAddresses,omitempty"`
	Issuetype interface{} `json:"issueType,omitempty"`
	Categorycode interface{} `json:"categoryCode,omitempty"`
	Language interface{} `json:"language,omitempty"`
	Subject interface{} `json:"subject"`
	Servicecode interface{} `json:"serviceCode,omitempty"`
	Severitycode interface{} `json:"severityCode,omitempty"`
}

// DateInterval represents the DateInterval schema from the OpenAPI specification
type DateInterval struct {
	Enddatetime interface{} `json:"endDateTime,omitempty"`
	Startdatetime interface{} `json:"startDateTime,omitempty"`
}

// Category represents the Category schema from the OpenAPI specification
type Category struct {
	Code interface{} `json:"code,omitempty"`
	Name interface{} `json:"name,omitempty"`
}

// TrustedAdvisorCheckResult represents the TrustedAdvisorCheckResult schema from the OpenAPI specification
type TrustedAdvisorCheckResult struct {
	Status interface{} `json:"status"`
	Timestamp interface{} `json:"timestamp"`
	Categoryspecificsummary interface{} `json:"categorySpecificSummary"`
	Checkid interface{} `json:"checkId"`
	Flaggedresources interface{} `json:"flaggedResources"`
	Resourcessummary TrustedAdvisorResourcesSummary `json:"resourcesSummary"` // Details about Amazon Web Services resources that were analyzed in a call to Trusted Advisor <a>DescribeTrustedAdvisorCheckSummaries</a>.
}

// TrustedAdvisorResourcesSummary represents the TrustedAdvisorResourcesSummary schema from the OpenAPI specification
type TrustedAdvisorResourcesSummary struct {
	Resourcessuppressed interface{} `json:"resourcesSuppressed"`
	Resourcesflagged interface{} `json:"resourcesFlagged"`
	Resourcesignored interface{} `json:"resourcesIgnored"`
	Resourcesprocessed interface{} `json:"resourcesProcessed"`
}

// DescribeCreateCaseOptionsRequest represents the DescribeCreateCaseOptionsRequest schema from the OpenAPI specification
type DescribeCreateCaseOptionsRequest struct {
	Categorycode interface{} `json:"categoryCode"`
	Issuetype interface{} `json:"issueType"`
	Language interface{} `json:"language"`
	Servicecode interface{} `json:"serviceCode"`
}

// RecentCaseCommunications represents the RecentCaseCommunications schema from the OpenAPI specification
type RecentCaseCommunications struct {
	Communications interface{} `json:"communications,omitempty"`
	Nexttoken interface{} `json:"nextToken,omitempty"`
}

// DescribeCreateCaseOptionsResponse represents the DescribeCreateCaseOptionsResponse schema from the OpenAPI specification
type DescribeCreateCaseOptionsResponse struct {
	Communicationtypes interface{} `json:"communicationTypes,omitempty"`
	Languageavailability interface{} `json:"languageAvailability,omitempty"`
}

// DescribeTrustedAdvisorCheckResultRequest represents the DescribeTrustedAdvisorCheckResultRequest schema from the OpenAPI specification
type DescribeTrustedAdvisorCheckResultRequest struct {
	Checkid interface{} `json:"checkId"`
	Language interface{} `json:"language,omitempty"`
}

// SupportedLanguage represents the SupportedLanguage schema from the OpenAPI specification
type SupportedLanguage struct {
	Code interface{} `json:"code,omitempty"`
	Display interface{} `json:"display,omitempty"`
	Language interface{} `json:"language,omitempty"`
}

// DescribeSeverityLevelsRequest represents the DescribeSeverityLevelsRequest schema from the OpenAPI specification
type DescribeSeverityLevelsRequest struct {
	Language interface{} `json:"language,omitempty"`
}

// TrustedAdvisorCheckRefreshStatus represents the TrustedAdvisorCheckRefreshStatus schema from the OpenAPI specification
type TrustedAdvisorCheckRefreshStatus struct {
	Checkid interface{} `json:"checkId"`
	Millisuntilnextrefreshable interface{} `json:"millisUntilNextRefreshable"`
	Status interface{} `json:"status"`
}

// DescribeTrustedAdvisorCheckSummariesRequest represents the DescribeTrustedAdvisorCheckSummariesRequest schema from the OpenAPI specification
type DescribeTrustedAdvisorCheckSummariesRequest struct {
	Checkids interface{} `json:"checkIds"`
}

// DescribeSeverityLevelsResponse represents the DescribeSeverityLevelsResponse schema from the OpenAPI specification
type DescribeSeverityLevelsResponse struct {
	Severitylevels interface{} `json:"severityLevels,omitempty"`
}

// AddCommunicationToCaseRequest represents the AddCommunicationToCaseRequest schema from the OpenAPI specification
type AddCommunicationToCaseRequest struct {
	Caseid interface{} `json:"caseId,omitempty"`
	Ccemailaddresses interface{} `json:"ccEmailAddresses,omitempty"`
	Communicationbody interface{} `json:"communicationBody"`
	Attachmentsetid interface{} `json:"attachmentSetId,omitempty"`
}

// SupportedHour represents the SupportedHour schema from the OpenAPI specification
type SupportedHour struct {
	Endtime interface{} `json:"endTime,omitempty"`
	Starttime interface{} `json:"startTime,omitempty"`
}

// DescribeTrustedAdvisorChecksResponse represents the DescribeTrustedAdvisorChecksResponse schema from the OpenAPI specification
type DescribeTrustedAdvisorChecksResponse struct {
	Checks interface{} `json:"checks"`
}

// RefreshTrustedAdvisorCheckRequest represents the RefreshTrustedAdvisorCheckRequest schema from the OpenAPI specification
type RefreshTrustedAdvisorCheckRequest struct {
	Checkid interface{} `json:"checkId"`
}

// TrustedAdvisorCostOptimizingSummary represents the TrustedAdvisorCostOptimizingSummary schema from the OpenAPI specification
type TrustedAdvisorCostOptimizingSummary struct {
	Estimatedpercentmonthlysavings interface{} `json:"estimatedPercentMonthlySavings"`
	Estimatedmonthlysavings interface{} `json:"estimatedMonthlySavings"`
}

// TrustedAdvisorResourceDetail represents the TrustedAdvisorResourceDetail schema from the OpenAPI specification
type TrustedAdvisorResourceDetail struct {
	Issuppressed interface{} `json:"isSuppressed,omitempty"`
	Metadata interface{} `json:"metadata"`
	Region interface{} `json:"region,omitempty"`
	Resourceid interface{} `json:"resourceId"`
	Status interface{} `json:"status"`
}

// DescribeSupportedLanguagesRequest represents the DescribeSupportedLanguagesRequest schema from the OpenAPI specification
type DescribeSupportedLanguagesRequest struct {
	Categorycode interface{} `json:"categoryCode"`
	Issuetype interface{} `json:"issueType"`
	Servicecode interface{} `json:"serviceCode"`
}

// CommunicationTypeOptions represents the CommunicationTypeOptions schema from the OpenAPI specification
type CommunicationTypeOptions struct {
	TypeField interface{} `json:"type,omitempty"`
	Dateswithoutsupport interface{} `json:"datesWithoutSupport,omitempty"`
	Supportedhours interface{} `json:"supportedHours,omitempty"`
}
