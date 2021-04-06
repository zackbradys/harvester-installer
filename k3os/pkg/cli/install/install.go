package install

import (
	"fmt"
	"os"

	"github.com/rancher/k3os/pkg/mode"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/rancher/harvester-installer/pkg/console"
)

func Command() cli.Command {
	mode, _ := mode.Get()
	return cli.Command{
		Name:  "install",
		Usage: "install k3OS",
		Flags: []cli.Flag{},
		Before: func(c *cli.Context) error {
			if os.Getuid() != 0 {
				return fmt.Errorf("must be run as root")
			}
			return nil
		},
		Action: func(*cli.Context) {
			if err := console.RunConsole(); err != nil {
				logrus.Error(err)
			}
		},
		Hidden: mode == "local",
	}
}
