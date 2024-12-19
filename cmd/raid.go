package cmd

import (
	"github.com/eddie-knight/raid-osps-baseline/armory"

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
	armory.SetupArmory(c)
	return
}
