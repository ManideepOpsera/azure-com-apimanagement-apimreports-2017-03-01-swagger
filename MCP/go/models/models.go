package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ReportCollection represents the ReportCollection schema from the OpenAPI specification
type ReportCollection struct {
	Count int64 `json:"count,omitempty"` // Total record count number across all pages.
	Nextlink string `json:"nextLink,omitempty"` // Next page link if any.
	Value []ReportRecordContract `json:"value,omitempty"` // Page values.
}

// ReportRecordContract represents the ReportRecordContract schema from the OpenAPI specification
type ReportRecordContract struct {
	Apiid string `json:"apiId,omitempty"` // API identifier path. /apis/{apiId}
	Cachehitcount int `json:"cacheHitCount,omitempty"` // Number of times when content was served from cache policy.
	Servicetimemin float64 `json:"serviceTimeMin,omitempty"` // Minimum time it took to process request on backend.
	Cachemisscount int `json:"cacheMissCount,omitempty"` // Number of times content was fetched from backend.
	Interval string `json:"interval,omitempty"` // Length of aggregation period. Interval must be multiple of 15 minutes and may not be zero. The value should be in ISO 8601 format (http://en.wikipedia.org/wiki/ISO_8601#Durations).
	Subscriptionid string `json:"subscriptionId,omitempty"` // Subscription identifier path. /subscriptions/{subscriptionId}
	Callcountfailed int `json:"callCountFailed,omitempty"` // Number of calls failed due to proxy or backend errors. This includes calls returning HttpStatusCode.BadRequest(400) and any Code between HttpStatusCode.InternalServerError (500) and 600
	Userid string `json:"userId,omitempty"` // User identifier path. /users/{userId}
	Apiregion string `json:"apiRegion,omitempty"` // API region identifier.
	Apitimeavg float64 `json:"apiTimeAvg,omitempty"` // Average time it took to process request.
	Callcountsuccess int `json:"callCountSuccess,omitempty"` // Number of successful calls. This includes calls returning HttpStatusCode <= 301 and HttpStatusCode.NotModified and HttpStatusCode.TemporaryRedirect
	Bandwidth int64 `json:"bandwidth,omitempty"` // Bandwidth consumed.
	Callcountother int `json:"callCountOther,omitempty"` // Number of other calls.
	Apitimemin float64 `json:"apiTimeMin,omitempty"` // Minimum time it took to process request.
	Callcountblocked int `json:"callCountBlocked,omitempty"` // Number of calls blocked due to invalid credentials. This includes calls returning HttpStatusCode.Unauthorized and HttpStatusCode.Forbidden and HttpStatusCode.TooManyRequests
	Callcounttotal int `json:"callCountTotal,omitempty"` // Total number of calls.
	Country string `json:"country,omitempty"` // Country to which this record data is related.
	Productid string `json:"productId,omitempty"` // Product identifier path. /products/{productId}
	Apitimemax float64 `json:"apiTimeMax,omitempty"` // Maximum time it took to process request.
	Region string `json:"region,omitempty"` // Country region to which this record data is related.
	Servicetimemax float64 `json:"serviceTimeMax,omitempty"` // Maximum time it took to process request on backend.
	Name string `json:"name,omitempty"` // Name depending on report endpoint specifies product, API, operation or developer name.
	Operationid string `json:"operationId,omitempty"` // Operation identifier path. /apis/{apiId}/operations/{operationId}
	Servicetimeavg float64 `json:"serviceTimeAvg,omitempty"` // Average time it took to process request on backend.
	Timestamp string `json:"timestamp,omitempty"` // Start of aggregation period. The date conforms to the following format: `yyyy-MM-ddTHH:mm:ssZ` as specified by the ISO 8601 standard.
	Zip string `json:"zip,omitempty"` // Zip code to which this record data is related.
}

// RequestReportCollection represents the RequestReportCollection schema from the OpenAPI specification
type RequestReportCollection struct {
	Count int64 `json:"count,omitempty"` // Total record count number across all pages.
	Value []RequestReportRecordContract `json:"value,omitempty"` // Page values.
}

// RequestReportRecordContract represents the RequestReportRecordContract schema from the OpenAPI specification
type RequestReportRecordContract struct {
	Servicetime float64 `json:"serviceTime,omitempty"` // he time it took to forward this request to the backend and get the response back.
	Apiid string `json:"apiId,omitempty"` // API identifier path. /apis/{apiId}
	Apitime float64 `json:"apiTime,omitempty"` // The total time it took to process this request.
	Userid string `json:"userId,omitempty"` // User identifier path. /users/{userId}
	Ipaddress string `json:"ipAddress,omitempty"` // The client IP address associated with this request.
	Timestamp string `json:"timestamp,omitempty"` // The date and time when this request was received by the gateway in ISO 8601 format.
	Method string `json:"method,omitempty"` // The HTTP method associated with this request..
	Subscriptionid string `json:"subscriptionId,omitempty"` // Subscription identifier path. /subscriptions/{subscriptionId}
	Requestsize int `json:"requestSize,omitempty"` // The size of this request..
	Responsesize int `json:"responseSize,omitempty"` // The size of the response returned by the gateway.
	Productid string `json:"productId,omitempty"` // Product identifier path. /products/{productId}
	Url string `json:"url,omitempty"` // The full URL associated with this request.
	Requestid string `json:"requestId,omitempty"` // Request Identifier.
	Cache string `json:"cache,omitempty"` // Specifies if response cache was involved in generating the response. If the value is none, the cache was not used. If the value is hit, cached response was returned. If the value is miss, the cache was used but lookup resulted in a miss and request was fulfilled by the backend.
	Apiregion string `json:"apiRegion,omitempty"` // Azure region where the gateway that processed this request is located.
	Responsecode int `json:"responseCode,omitempty"` // The HTTP status code returned by the gateway.
	Backendresponsecode string `json:"backendResponseCode,omitempty"` // The HTTP status code received by the gateway as a result of forwarding this request to the backend.
	Operationid string `json:"operationId,omitempty"` // Operation identifier path. /apis/{apiId}/operations/{operationId}
}
