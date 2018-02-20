package config

const (
	ModeClient = "client"
	ModeKeeper = "keeper"
	ModeServer = "server"
)

type StartupConfig struct {
	mode string
}
