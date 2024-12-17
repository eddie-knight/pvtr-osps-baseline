package cmd

import (
	"github.com/eddie-knight/raid-osps-baseline/armory"
	"github.com/eddie-knight/raid-osps-baseline/transport"

	"github.com/privateerproj/privateer-sdk/config"
	"github.com/privateerproj/privateer-sdk/raidengine"
)

var (
	Vessel = raidengine.Vessel{
		RaidName:    "osps-baseline",
		Armory:      &armory.Armory,
		Initializer: initializer,
		RequiredVars: []string{
			"owner",
			"repo",
		},
	} // Used by the plugin or debug function to run the Raid
)

type Raid struct{}

// Raid.Start() is called by plugin.Serve
func (r *Raid) Start() (err error) {
	err = Vessel.Mobilize()
	return
}

func initializer(c *config.Config) (err error) {
	armory.Config = c // for strikes to reference. TODO: not sure yet whether this mitigates the need for armory.Armory.Config
	repoData, err := transport.GetRepoData(c)
	if err != nil {
		return
	}
	c.Vars["repo_data"] = repoData
	return
}
