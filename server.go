package main

import (
	"flag"

	"github.com/spf13/cobra"
	"github.com/golang/glog"


	)

var cmdServer = &cobra.Command{
	Use: "server",
	Aliases: []string{"serve"},
	Short: "starts the api server",
	Run: func(cmd *cobra.Command, args []string){
		flag.Parse()
		defer glog.Flush()

		if err := Run(":8080"); err != nil {
			glog.Fatal(err)
		}
	},
}