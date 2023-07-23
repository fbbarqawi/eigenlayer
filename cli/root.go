package cli

import (
	"github.com/NethermindEth/eigenlayer/cli/prompter"
	"github.com/NethermindEth/eigenlayer/pkg/daemon"
	"github.com/spf13/cobra"
)

func RootCmd(d daemon.Daemon, p prompter.Prompter) *cobra.Command {
	cmd := cobra.Command{
		Use: "eigenlayer",
	}
	cmd.AddCommand(
		InstallCmd(d, p),
		LocalInstallCmd(d),
		StopCmd(d),
		UninstallCmd(d),
		PluginCmd(d),
		RunCmd(d),
		ListCmd(d),
		LogsCmd(d),
	)
	cmd.CompletionOptions.DisableDefaultCmd = true
	return &cmd
}
