package app

var Config Configuration

type Configuration struct {
	ConnectionString string
	Database         string
	Collection       string
	ApiKey           string
}
