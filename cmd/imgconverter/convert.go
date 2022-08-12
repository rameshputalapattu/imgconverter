package main

import (
	"context"
	"errors"
	"flag"
	"os"
	"path/filepath"

	ew "github.com/pkg/errors"
	imgconverter "github.com/rameshputalapattu/imgconverter"
)

//convertCommand Command to convert images from png format to kmeg binary format
type convertCommand struct {
	Params *CmdParams
}

const convertHelp = `convert a image from jpeg/jpg to png and vice-versa`

//Name Gives the command Name
func (cmd *convertCommand) Name() string { return "convert" }

//Args returns the command args
func (cmd *convertCommand) Args() string { return "" }

//ShortHelp returns short help text
func (cmd *convertCommand) ShortHelp() string { return convertHelp }

//LongHelp returns long help text
func (cmd *convertCommand) LongHelp() string { return convertHelp }

//Hidden returns whether it is a hidden command
func (cmd *convertCommand) Hidden() bool { return false }

//Register Registers the flag set
func (cmd *convertCommand) Register(fs *flag.FlagSet) {

}

//Run the convert command to convert images from jpg/jpeg to png and vice versa
func (cmd *convertCommand) Run(ctx context.Context, args []string) error {

	if len(cmd.Params.SrcImageFile) == 0 {
		return errors.New("Source Image file should be provided for convert")
	}

	if len(cmd.Params.DstImageFile) == 0 {
		return errors.New("Destination Image file should be provided for convert")
	}

	return convert(cmd.Params.SrcImageFile, cmd.Params.DstImageFile)

}

func convert(srcFile, destFile string) error {
	r, err := os.Open(srcFile)

	if err != nil {
		return ew.Wrapf(err, "error opening the original image file %s\n", srcFile)
	}

	defer r.Close()

	ext := filepath.Ext(srcFile)

	img, err := imgconverter.ReadImage(r, ext)

	if err != nil {
		return ew.Wrap(err, "error decoding the image")
	}

	imgRGBA := imgconverter.ConvertToRGBA(img)

	w, err := os.Create(destFile)

	if err != nil {
		return ew.Wrap(err, "creating the file to write converted image failed")
	}

	defer w.Close()

	ext = filepath.Ext(destFile)

	err = imgconverter.WriteImage(imgRGBA, w, ext)

	if err != nil {
		return ew.Wrap(err, "error converting the image")
	}

	return nil

}
