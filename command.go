package main

import (
	"GoWebServer/client"
	"GoWebServer/server"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var clientCommand = cli.Command{
	Name:  "client",
	Usage: `Run Client`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "p",
			Usage: "protocol tcp or udp",
		},
		cli.StringFlag{
			Name:  "d",
			Usage: "Download file",
		},
		cli.BoolFlag{
			Name:  "l",
			Usage: "List",
		},

		cli.BoolFlag{
			Name:  "t",
			Usage: "time",
		},
	},
	Action: func(ctx *cli.Context) error {
		protocol := ctx.String("p")
		list := ctx.Bool("l")
		time := ctx.Bool("t")
		path := ctx.String("d")

		if protocol == "tcp" {
			client.HandleTcp(list, path)
		} else if protocol == "udp" {
			log.Info("udp client\n")
			if time {
				client.HandleUdp()
			} else {
				return nil
			}
		} else {
			log.Error("Unknown protocal")
			return nil
		}
		return nil
	},
}
var serverCommand = cli.Command{
	Name:  "server",
	Usage: `Run Server`,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "p",
			Usage: "protocol tcp or udp",
		},
	},
	Action: func(ctx *cli.Context) error {
		protocol := ctx.String("p")
		if protocol == "tcp" {
			server.TcpServer()
		} else if protocol == "udp" {
			log.Info("udp server\n")
			server.UdpServer()
		} else {
			log.Error("Unknown protocal")
			return nil
		}
		return nil
	},
}
