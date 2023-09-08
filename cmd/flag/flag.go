package flag

import "github.com/urfave/cli/v2"

var (
	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "log.level",
			Value:       "debug",
			Usage:       "Set logging level",
			DefaultText: "(info, error, warn, debug)",
			EnvVars:     []string{"ZSYNC_LOG_LEVEL"},
		},
	}
)

func Flags() []cli.Flag {
	return flags
}
