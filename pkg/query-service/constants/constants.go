package constants

import (
	"os"
	"strconv"

	"go.signoz.io/query-service/model"
)

const (
	HTTPHostPort    = "0.0.0.0:8080" // Address to serve http (query service)
	PrivateHostPort = "0.0.0.0:8085" // Address to server internal services like alert manager
	DebugHttpPort   = "0.0.0.0:6060" // Address to serve http (pprof)
)

var DEFAULT_TELEMETRY_ANONYMOUS = false

func IsTelemetryEnabled() bool {
	isTelemetryEnabledStr := os.Getenv("TELEMETRY_ENABLED")
	isTelemetryEnabledBool, err := strconv.ParseBool(isTelemetryEnabledStr)
	if err != nil {
		return true
	}
	return isTelemetryEnabledBool
}

const TraceTTL = "traces"
const MetricsTTL = "metrics"

func GetAlertManagerApiPrefix() string {
	if os.Getenv("ALERTMANAGER_API_PREFIX") != "" {
		return os.Getenv("ALERTMANAGER_API_PREFIX")
	}
	return "http://alertmanager:9093/api/"
}

// Alert manager channel subpath
var AmChannelApiPath = GetOrDefaultEnv("ALERTMANAGER_API_CHANNEL_PATH", "v1/routes")

var RELATIONAL_DATASOURCE_PATH = GetOrDefaultEnv("SIGNOZ_LOCAL_DB_PATH", "/var/lib/signoz/signoz.db")

const (
	ServiceName                    = "serviceName"
	HttpRoute                      = "httpRoute"
	HttpCode                       = "httpCode"
	HttpHost                       = "httpHost"
	HttpUrl                        = "httpUrl"
	HttpMethod                     = "httpMethod"
	Component                      = "component"
	OperationDB                    = "name"
	OperationRequest               = "operation"
	Status                         = "status"
	Duration                       = "duration"
	DBName                         = "dbName"
	DBOperation                    = "dbOperation"
	DBSystem                       = "dbSystem"
	MsgSystem                      = "msgSystem"
	MsgOperation                   = "msgOperation"
	Timestamp                      = "timestamp"
	RPCMethod                      = "rpcMethod"
	ResponseStatusCode             = "responseStatusCode"
	Descending                     = "descending"
	Ascending                      = "ascending"
	ContextTimeout                 = 60 // seconds
	StatusPending                  = "pending"
	StatusFailed                   = "failed"
	StatusSuccess                  = "success"
	ExceptionType                  = "exceptionType"
	ExceptionCount                 = "exceptionCount"
	LastSeen                       = "lastSeen"
	FirstSeen                      = "firstSeen"
	Attributes                     = "attributes"
	Resources                      = "resources"
	Static                         = "static"
	DefaultLogSkipIndexType        = "bloom_filter(0.01)"
	DefaultLogSkipIndexGranularity = 64
)
const (
	SIGNOZ_METRIC_DBNAME        = "signoz_metrics"
	SIGNOZ_SAMPLES_TABLENAME    = "samples_v2"
	SIGNOZ_TIMESERIES_TABLENAME = "time_series_v2"
)

var TimeoutExcludedRoutes = map[string]bool{
	"/api/v1/logs/tail": true,
}

// alert related constants
const (
	// AlertHelpPage is used in case default alert repo url is not set
	AlertHelpPage = "https://signoz.io/docs/userguide/alerts-management/#generator-url"
)

func GetOrDefaultEnv(key string, fallback string) string {
	v := os.Getenv(key)
	if len(v) == 0 {
		return fallback
	}
	return v
}

var StaticInterestingLogFields = []model.LogField{
	{
		Name:     "trace_id",
		DataType: "String",
		Type:     Static,
	},
	{
		Name:     "span_id",
		DataType: "String",
		Type:     Static,
	},
	{
		Name:     "trace_flags",
		DataType: "UInt32",
		Type:     Static,
	},
	{
		Name:     "severity_text",
		DataType: "LowCardinality(String)",
		Type:     Static,
	},
	{
		Name:     "severity_number",
		DataType: "Int32",
		Type:     Static,
	},
}

var StaticSelectedLogFields = []model.LogField{
	{
		Name:     "timestamp",
		DataType: "UInt64",
		Type:     Static,
	},
	{
		Name:     "id",
		DataType: "String",
		Type:     Static,
	},
}
