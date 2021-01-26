package dhcp

import (
	"github.com/spf13/cobra"
)

var (
	baseCmd = &cobra.Command{
		Use:   "dhcp",
		Short: "Manage and initialize DHCP",
		Long: "DHCP is a dependent service to Warewulf. This command will configure DHCP as defined\n" +
			"in the warewulf.conf file.",
		RunE: CobraRunE,
	}
	SetShow    bool
	SetPersist bool
)

func init() {
	baseCmd.PersistentFlags().BoolVarP(&SetShow, "show", "s", false, "Show configuration (don't update)")
	baseCmd.PersistentFlags().BoolVar(&SetPersist, "persist", false, "Persist the configuration and initialize the service")
}

// GetRootCommand returns the root cobra.Command for the application.
func GetCommand() *cobra.Command {
	return baseCmd
}
