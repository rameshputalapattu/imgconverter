package main

import (
	"context"
	"errors"
	"flag"

	"github.com/genuinetools/pkg/cli"
	"github.com/sirupsen/logrus"
)

func main() {

	p := cli.NewProgram()
	p.Name = "imgconverter"
	p.Description = "utility to convert images from jpg to png and vice-versa"

	params := &CmdParams{}

	p.Commands = []cli.Command{&convertCommand{params}}

	p.FlagSet = flag.NewFlagSet("global", flag.ExitOnError)
	p.FlagSet.StringVar(&params.SrcImageFile, "from", "", "source image file")
	p.FlagSet.StringVar(&params.DstImageFile, "to", "", "destination image file")

	p.Before = func(ctx context.Context) error {

		if len(params.SrcImageFile) == 0 || len(params.DstImageFile) == 0 {
			return errors.New("Both source and destination image file paths must be provided")
		}

		return nil

	}

	p.Run()
	logrus.Info("executed the command successfully")
}
