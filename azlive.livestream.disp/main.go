package main

import (
	"fmt"
	"github.com/azlive/azlive.livestream.disp/config"
	"github.com/azlive/azlive.livestream.disp/server"
	"github.com/urfave/cli"
	"os"
)

const (
	VERSION = "v1.0.0"
)

func Actionfunc(c *cli.Context) error {
	conf,err := config.Initconf()
	if err != nil {
		fmt.Printf("INIT FAILED!")
		return err
	}
	server.Startserver(conf)
	return nil
}



func main() {
	muc := cli.NewApp()
	muc.Name = "azlive.livestream.disp"
	muc.Usage = "It's a Dispatch Program for azlive."
	muc.Version = VERSION
	muc.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "c",
			Usage: "Load configuration from `FILE`",
		},
	}
	muc.Action = Actionfunc

	muc.Run(os.Args)
}
