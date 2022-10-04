package app

// Config exported variable to contain config values
var Config Configuration

// Configuration exported type for config
type Configuration struct {
	ConnectionString string
	Database         string
	Collection       string
	APIKey           string
	Address          string
	StartTimeout     int
	StopTimeout      int
}
