package constants

// Constant values required by Pi-CLI
const (
	// Port that the Pi-Hole API is defaulted to
	DefaultPort = 80
	// The default refresh rate of the data in seconds
	DefaultRefreshS = 1
	// The name of the configuration file
	ConfigFileName = "picli-config.json"
	// The starting setting for the number of queries that are included in the live log
	DefaultAmountOfQueries = 10
	// The default refresh rate for the UI
	DefaultUIFramesPerSecond = 30
)

// Constant values required for use in authentication and API key management
const (
	KeyringService = "PiCLI"
	KeyringUsr     = "api-key"
)

// Database constants
const (
	DBDriverName = "sqlite3"
)

// Keys that can be used to index JSON responses from the Pi-Hole's API
const (
	// Summary
	DNSQueriesTodayKey     = "dns_queries_today"
	AdsBlockedTodayKey     = "ads_blocked_today"
	PercentBlockedTodayKey = "ads_percentage_today"
	DomainsOnBlockListKey  = "domains_being_blocked"
	StatusKey              = "status"
	PrivacyLevelKey        = "privacy_level"
	TotalClientsSeenKey    = "clients_ever_seen"
	// TopItems
	TopQueriesTodayKey = "top_queries"
	TopAdsTodayKey     = "top_ads"
	// GetAllQueries
	AllQueryDataKey = "data"
)
