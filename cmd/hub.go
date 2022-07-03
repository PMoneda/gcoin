/*
Copyright © 2022 Philippe Moneda <philippe.moneda@gmail.com>

*/
package cmd

import (
	"github.com/PMoneda/gcoin/hub"
	"github.com/spf13/cobra"
)

// hubCmd represents the hub command
var hubCmd = &cobra.Command{
	Use:   "hub",
	Short: "hub cli to gcoin",
	Long: `Hub cli to gcoin
	* Start Hub`,
	Run: func(cmd *cobra.Command, args []string) {
		hub.StartHub()
	},
}

func init() {
	rootCmd.AddCommand(hubCmd)
}
