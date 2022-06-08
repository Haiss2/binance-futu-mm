package main

import (
	"log"
	"os"

	libapp "github.com/Haiss2/binance-futu-mm/pkg/app"
	"github.com/Haiss2/binance-futu-mm/pkg/dbutil"
	"github.com/Haiss2/binance-futu-mm/pkg/hunter"
	"github.com/Haiss2/binance-futu-mm/pkg/server"
	"github.com/Haiss2/binance-futu-mm/pkg/storage"
	"github.com/urfave/cli"
)

func main() {
	app := libapp.NewApp()
	app.Name = "binance future market marker"
	app.Action = run

	if err := app.Run(os.Args); err != nil {
		log.Panic(err)
	}
}

func run(c *cli.Context) error {
	logger, _ := libapp.NewLogger(c)
	l := logger.Sugar()

	l.Info("app starting ..")

	db, err := libapp.NewDBFromContext(c)
	if err != nil {
		l.Panicw("cannot init DB connection", "err", err)
	}

	// Load cli value
	migrationPath := c.String(libapp.PostgresMigrationPath)
	pgDatabase := c.String(libapp.PostgresDatabaseFlag)
	bindAddress := c.String(libapp.BindAddressFlag)

	_, err = dbutil.RunMigrationUp(db.DB, migrationPath, pgDatabase)
	if err != nil {
		l.Panicw("cannot init DB", "err", err)
	}

	store := storage.New(db)
	_ = store

	hunter := hunter.NewHunter(l, store)

	server := server.New(bindAddress, hunter)

	return server.Run()
}
