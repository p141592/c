package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func _init() error {
	log.Fatal("Установка компонентов cli_home")
	return nil
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "install",
				Usage: "Установка cli_home на систему",
				Action: func(c *cli.Context) error {return _init()},
			}
		}
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
