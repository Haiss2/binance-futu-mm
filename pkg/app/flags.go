package app

import (
	"github.com/urfave/cli"
)

const (
	BindAddressFlag = "bind-address"
)

func NewAppFlag() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:   BindAddressFlag,
			Usage:  "application host and port",
			EnvVar: "BIND_ADDRESS",
		},
	}
}
