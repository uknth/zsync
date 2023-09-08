package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/uknth/zsync/cmd/flag"
	"github.com/uknth/zsync/cmd/tag"
	"github.com/uknth/zsync/internal/zsync"
	"github.com/urfave/cli/v2"
)

func main() {
	var scr zsync.Syncer

	err := (&cli.App{
		Name:    "zsync",
		Usage:   "zsync is simple utility to sync zsh history",
		Version: tag.VERSION,
		Flags:   flag.Flags(),
		Commands: []*cli.Command{
			{
				Name:    "sync",
				Aliases: []string{"s"},
				Before: func(ctx *cli.Context) error {
					ss, err := zsync.NewDefaultSync()

					if err != nil {
						return fmt.Errorf("failed to initialise zsync: %w", err)
					}

					scr = ss
					return err
				},

				Action: func(ctx *cli.Context) error {
					return scr.Sync()
				},
			},
		},
	}).Run(os.Args)
	if err != nil {
		log.Println("Something went wrong!. Failed to run 'zsync'")
		log.Println("-----------------------------------------------")

		in := ">"
		for e := errors.Unwrap(err); e != nil; {
			in = in + ">"
			fmt.Printf("	%s %s %v \n", in, e.Error(), e)
		}

		log.Fatalf("-----------------------------------------------\n")
	}
}
