package main

import (
	"os"

	"github.com/tanya-saroha/taskmanager-service/app"
	"github.com/tanya-saroha/taskmanager-service/config"
	"github.com/tanya-saroha/taskmanager-service/db"
	"github.com/tanya-saroha/taskmanager-service/server"
	"github.com/urfave/cli"
)

func main() {
	config.Load("application")
	app.Init()
	defer app.Close()

	cliApp := cli.NewApp()
	cliApp.Name = "taskmanager"
	cliApp.Version = "1.0.0"
	cliApp.Commands = []cli.Command{
		{
			Name:  "start",
			Usage: "start server",
			Action: func(c *cli.Context) error {
				server.StartAPIServer()
				return nil
			},
		},
		{
			Name:  "create_migration",
			Usage: "create migration file",
			Action: func(c *cli.Context) error {
				return db.CreateMigrationFile(c.Args().Get(0))
			},
		},
		{
			Name:  "migrate",
			Usage: "run db migrations",
			Action: func(c *cli.Context) error {
				err := db.RunMigrations()
				return err
			},
		},
		{
			Name:  "rollback",
			Usage: "rollback migrations",
			Action: func(c *cli.Context) error {
				return db.RollbackMigrations(c.Args().Get(0))
			},
		},
	}
	if err := cliApp.Run(os.Args); err != nil {
		panic(err)
	}
}
