package app

import (
	"github.com/Haiss2/binance-futu-mm/pkg/dbutil"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" //nolint sql driver name: "postgres"
	"github.com/urfave/cli"
)

const (
	PostgresHostFlag      = "postgres-host"
	PostgresPortFlag      = "postgres-port"
	PostgresUserFlag      = "postgres-user"
	PostgresPasswordFlag  = "postgres-password"
	PostgresDatabaseFlag  = "postgres-database"
	PostgresMigrationPath = "migration-path"
)

// PostgresSQLFlags creates new cli flags for PostgreSQL client.
func PostgresSQLFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   PostgresHostFlag,
			Usage:  "PostgresSQL host to connect",
			EnvVar: "POSTGRES_HOST",
		},
		cli.IntFlag{
			Name:   PostgresPortFlag,
			Usage:  "PostgresSQL port to connect",
			EnvVar: "POSTGRES_PORT",
		},
		cli.StringFlag{
			Name:   PostgresUserFlag,
			Usage:  "PostgresSQL user to connect",
			EnvVar: "POSTGRES_USER",
		},
		cli.StringFlag{
			Name:   PostgresPasswordFlag,
			Usage:  "PostgresSQL password to connect",
			EnvVar: "POSTGRES_PASSWORD",
		},
		cli.StringFlag{
			Name:   PostgresDatabaseFlag,
			Usage:  "Postgres database to connect",
			EnvVar: "POSTGRES_DATABASE",
		},
		cli.StringFlag{
			Name:   PostgresMigrationPath,
			Value:  "migrations",
			EnvVar: "MIGRATION_PATH",
		},
	}
}

// NewDBFromContext creates a DB instance from cli flags configuration.
func NewDBFromContext(c *cli.Context) (*sqlx.DB, error) {
	const driverName = "postgres"

	connStr := dbutil.FormatDSN(map[string]string{
		"host":     c.String(PostgresHostFlag),
		"port":     c.String(PostgresPortFlag),
		"user":     c.String(PostgresUserFlag),
		"password": c.String(PostgresPasswordFlag),
		"dbname":   c.String(PostgresDatabaseFlag),
		"sslmode":  "disable",
	})

	return sqlx.Connect(driverName, connStr)
}
