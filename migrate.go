package main

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/mattes/migrate"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
	"github.com/spf13/cobra"
)

var cmdMigrate = &cobra.Command{
	Use:   "migrate",
	Short: "run database migrations",
}

func init() {
	var options struct {
		dir string
	}
	fs := cmdMigrate.Flags()
	fs.StringVar(&options.dir, "dir", "file://db/migrate", "directory containing db migration files")

	cmdMigrate.Run = func(cmd *cobra.Command, args []string) {
		//c := make(chan interface{})
		cs := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=%s", config.Database.Username, config.Database.Password, config.Database.DatabaseName, config.Database.Sslmode)
		glog.Infof("migrations db: %s", config.Database.DatabaseName)
		glog.Infof("migrations user: %s", config.Database.Username)
		glog.Infof("migrations sslmode: %s", config.Database.Sslmode)
		m, err := migrate.New(
			options.dir,
			cs,
		)
		if err != nil {
			glog.Fatalf("migration error %v", err)
		}
		m.Steps(2)

		/*for v := range c {
			if err, ok := v.(error); ok {
				glog.Errorf("migration error: %v", err)
			} else {
				glog.Info(v)
			}
		}*/
	}
}
