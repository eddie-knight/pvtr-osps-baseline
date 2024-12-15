package cmd

import (
	"github.com/privateerproj/privateer-raid-osps-baseline/armory"

	"github.com/privateerproj/privateer-sdk/raidengine"
)

var (
	Vessel = raidengine.Vessel{
		RaidName: "osps-baseline", // Double check that this is what you want the raid to be named
		Armory:   &armory.Armory,
		RequiredVars: []string{
			"repo_url",
		},
	} // Used by the plugin or debug function to run the Raid
)

type Raid struct{}

// Raid.Start() is called by plugin.Serve
func (r *Raid) Start() (err error) {
	err = Vessel.Mobilize() // Replace nil with a slice of your required var names
	return
}
