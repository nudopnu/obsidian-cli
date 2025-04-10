package commands

const (
	localCertFile = "obsidian-local-rest-api.crt"
)

type State struct {
	ApiKey string
	Plain  bool
}
