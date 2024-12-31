package convert

import (
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

type RemoteConfigJSON struct {
	Name  string   `json:"name"`
	URLS  []string `json:"urls"`
	Fetch []string `json:"fetch"`
}

func ConvertToRemoteConfigJSON(remoteConfig *config.RemoteConfig) RemoteConfigJSON {
	fetchSpecs := make([]string, len(remoteConfig.Fetch))
	for i, fetchSpec := range remoteConfig.Fetch {
		fetchSpecs[i] = fetchSpec.String()
	}

	return RemoteConfigJSON{
		Name:  remoteConfig.Name,
		URLS:  remoteConfig.URLs,
		Fetch: fetchSpecs,
	}
}

func ConvertToRemotesConfigJSON(remotes []*git.Remote) []RemoteConfigJSON {
	remoteConfigs := []RemoteConfigJSON{}
	for _, remote := range remotes {
		config := remote.Config()
		if config.Name != "" {
			remoteConfigs = append(remoteConfigs, ConvertToRemoteConfigJSON(config))
		}
	}

	return remoteConfigs
}
