package config

const (
	ModeClient = "client"
	ModeKeeper = "keeper"
	ModeServer = "server"
)

type StartupConfig struct {
	drivers map[string]StartupConfigDriverSection
	mode string
	routes []StartupConfigRouteSection
}

type StartupConfigDriverSection struct {
	kind string
	options map[string] interface{}
	poolSize int
}

type StartupConfigRouteSection struct {

}
