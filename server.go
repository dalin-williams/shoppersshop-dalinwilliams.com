package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
)

var cmdServer = &cobra.Command{
	Use:     "server",
	Aliases: []string{"serve"},
	Short:   "starts the api server",
	Run: func(cmd *cobra.Command, args []string) {
		glog.Infof("Starting the magic on port 8080")
		flag.Parse()
		defer glog.Flush()

		if err := Run(":8080"); err != nil {
			glog.Fatal(err)
		}
	},
}
