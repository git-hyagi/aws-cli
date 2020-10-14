package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "aws-test",
	Short: "This is a very simple program to manage some aws instances.",
	Long: `Very simple program to start, stop or print the state of some instances running on aws
the flags "name" and "owner" are used as filters when making the ec2 actions (poweron/poweroff/state).

For this program to run properly it is necessary to have the aws credentials stored in your home directory ($HOME/.aws/credentials).
More information about aws credentials: https://docs.aws.amazon.com/cli/latest/userguide/cli-configure-files.html`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().StringP("name", "n", "", "Instance name (required if INSTANCE_NAME env var not defined), for example: hyagi-*")
	rootCmd.PersistentFlags().StringP("owner", "o", "504887044062", "AWS Account Owner (can be suppressed if INSTANCE_OWNER env var is defined)")

	if os.Getenv("INSTANCE_NAME") == "" {
		rootCmd.MarkPersistentFlagRequired("name")
	}

	rootCmd.AddCommand(state())
	rootCmd.AddCommand(powerOn())
	rootCmd.AddCommand(powerOff())
}
