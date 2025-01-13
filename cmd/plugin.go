package cmd

import (
	"github.com/eddie-knight/pvtr-osps-baseline/armory"

	"github.com/privateerproj/privateer-sdk/config"
	"github.com/privateerproj/privateer-sdk/pluginkit"
)

var (
	Vessel = pluginkit.Vessel{
		PluginName:  "osps-baseline",
		Armory:      &armory.Armory,
		Initializer: initializer,
		RequiredVars: []string{
			"owner",
			"repo",
		},
	} // Used by the plugin or debug function to run the Plugin
)

type Plugin struct{}

// Plugin.Start() is called by plugin.Serve
func (p *Plugin) Start() (err error) {
	err = Vessel.Mobilize()
	return
}

func initializer(c *config.Config) (err error) {
	armory.SetupArmory(c)
	return
}
