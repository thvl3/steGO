package main

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:
  $ source <(steGo completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ steGo completion bash > /etc/bash_completion.d/steGo
  # macOS:
  $ steGo completion bash > /usr/local/etc/bash_completion.d/steGo

Zsh:
  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:
  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ steGo completion zsh > "${fpath[1]}/_steGo"

  # You will need to start a new shell for this setup to take effect.

fish:
  $ steGo completion fish | source

  # To load completions for each session, execute once:
  $ steGo completion fish > ~/.config/fish/completions/steGo.fish

PowerShell:
  PS> steGo completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> steGo completion powershell > steGo.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
