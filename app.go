package main

import (
	"fmt"
	"github.com/armon/go-socks5"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"os"
)

type Socks5NameserverApp struct {
	*cli.App
}

func NewApp(version string) *Socks5NameserverApp {
	app := &Socks5NameserverApp{cli.NewApp()}
	app.Name = "socks5-nameserver"
	app.Version = version
	app.Usage = "socks5-nameserver change your nameserver and resolve on sock5"
	app.ErrWriter = os.Stderr
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "0.0.0.0",
			Usage: "Set host for listen valid are 127.0.0.1 and 0.0.0.0",
		},
		cli.IntFlag{
			Name:  "port, p",
			Usage: "Listen port",
			Value: 8000,
		},
		cli.StringFlag{
			Name:  "nameserver, n",
			Usage: "Set dns nameserver",
		},
		cli.BoolFlag{
			Name:  "in-tcp",
			Usage: "Set dns nameserver to be called in tcp instead of udp",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "launch socks5-nameserver",
			Action: start,
		},
	}
	return app
}

func start(c *cli.Context) error {
	log.SetOutput(os.Stderr)
	nameserver := c.GlobalString("nameserver")
	if nameserver == "" {
		return fmt.Errorf("Nameserver must be defined.")
	}
	conf := &socks5.Config{
		Resolver: NewCustomDNSResolver(nameserver, c.GlobalBool("in-tcp")),
	}
	server, err := socks5.New(conf)
	if err != nil {
		return err
	}
	listen := fmt.Sprintf("%s:%d", c.GlobalString("host"), c.GlobalInt("port"))
	log.Infof("Sock5-nameserver is started and listening on %s", listen)
	return server.ListenAndServe("tcp", listen)
}
