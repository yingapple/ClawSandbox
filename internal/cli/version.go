package cli

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"

	"github.com/weiyong1024/clawsandbox/internal/version"
)

var versionShort bool

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of ClawSandbox",
	Run: func(cmd *cobra.Command, args []string) {
		if versionShort {
			fmt.Println(version.Version)
			return
		}
		fmt.Printf("clawsandbox %s\n", version.Version)
		fmt.Printf("  commit:    %s\n", version.GitCommit)
		fmt.Printf("  built:     %s\n", version.BuildDate)
		fmt.Printf("  go:        %s\n", runtime.Version())
		fmt.Printf("  platform:  %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	versionCmd.Flags().BoolVar(&versionShort, "short", false, "Print version number only")
}
