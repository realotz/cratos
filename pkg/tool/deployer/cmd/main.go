package cmd

import (
	"fmt"
	"github.com/realotz/cratos/pkg/tool/deployer"
	"github.com/urfave/cli"
	"log"
	"os"
)

var version = "snapshot"

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s\nversion=%s revision=%s\n", c.App.Name, c.App.Version, version)
	}

	app := cli.NewApp()
	app.Name = "deployer"
	app.Usage = "Deploy/upgrade deployments on cratos k8s.io"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.Int64SliceFlag{
			Name:  "ports,pt",
			Usage: "`ports`",
		},
		cli.StringSliceFlag{
			Name:  "env,e",
			Usage: "`env`",
		},
		cli.StringFlag{
			Name:  "config,c",
			Usage: "`config`",
		},
		cli.StringFlag{
			Name:  "file,f",
			Usage: "`config file`",
		},
		cli.StringFlag{
			Name:  "name,n",
			Value: "",
			Usage: "name",
		},
		cli.StringFlag{
			Name:  "namespace,ns",
			Value: "",
			Usage: "k8s.io namespace",
		},
		cli.StringFlag{
			Name:  "image,i",
			Value: "",
			Usage: "Override the Docker `IMAGE-URL`",
		},
	}

	app.Action = func(c *cli.Context) error {
		if c.String("token") != "" {
			_ = os.Setenv("DEPLOYER_TOKEN", c.String("token"))
		}
		deployer.Deploy(c)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
