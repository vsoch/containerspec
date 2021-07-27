package cli

import (
	"fmt"
	"github.com/DataDrake/cli-ng/v2/cmd"
	"github.com/vsoch/containerspec/spec"
)

// Args and flags for generate
type HostArgs struct{}
type HostFlags struct{}

// Host dumps out information about the host
var Host = cmd.Sub{
	Name:  "host",
	Alias: "h",
	Short: "Dump out information about the host.",
	Flags: &HostFlags{},
	Args:  &HostArgs{},
	Run:   RunHost,
}

func init() {
	cmd.Register(&Host)
}

// RunHost detects the host of the machine
func RunHost(r *cmd.Root, c *cmd.Sub) {

	// Create and load a new config
	// conf := config.Load(args.Chefyaml[0])

	// Detect the host
	arch := spec.Detect()
	fmt.Println(arch)
}
