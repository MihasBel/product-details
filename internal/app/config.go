package app

// Config exported variable to contain config values
var Config configuration

type configuration struct {
	ConnectionString string
	Database         string
	Collection       string
	APIKey           string
}
