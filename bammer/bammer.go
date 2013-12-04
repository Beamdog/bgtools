package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
	"image"
	"image/gif"
	"github.com/Beamdog/bgfileformats"
)

var input = flag.String("input", "input.bam", "Source [bam, bamd, gif]")
var output = flag.String("output", "out", "Output Directory/Filename")
var palette = flag.String("palette", "palette.png", "Palette to use for the bam")
var mirror = flag.Bool("mirror", false, "Set to true to mirror all the frames")
var offset_x = flag.Int("offsetx", 0, "Offset all exported frames in X direction")
var offset_y = flag.Int("offsety", 0, "Offset all exported frames in Y direction")

func main() {
	flag.Parse()

	bamFileIn, err := os.Open(filepath.Clean(*input))
	if err != nil {
		log.Fatal(err)
	}
	defer bamFileIn.Close()
	if strings.ToLower(filepath.Ext(*input)) == ".bam" {

		bamIn, err := bg.OpenBAM(bamFileIn)
		if err != nil {
			log.Fatal(err)
		}
		outputName := strings.TrimSuffix(filepath.Base(*input), filepath.Ext(*input))
		if strings.ToLower(filepath.Ext(*output)) == ".bamd" {
			bamIn.MakeBamd(filepath.Clean(*output), outputName, *mirror, *offset_x, *offset_y)
		} else if strings.ToLower(filepath.Ext(*output)) == ".gif" {
			bamIn.MakeGif(*output, outputName)
		}
	} else if strings.ToLower(filepath.Ext(*input)) == ".bamd" {
		bamOut, err := bg.OpenBAMD(bamFileIn, *palette)
		if err != nil {
			log.Fatal(err)
		}
		outBam, err := os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
		defer outBam.Close()

		bamOut.MakeBam(outBam)
	} else if strings.ToLower(filepath.Ext(*input)) == ".gif" {
		gif, err := gif.DecodeAll(bamFileIn)
		if err != nil {
			log.Fatal(err)
		}

		outFile, err := os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()

		sequences := make([]image.Point, 1)
		sequences[0] = image.Pt(0,len(gif.Image))
		bam, err := bg.MakeBamFromGif(gif, sequences)
		if err != nil {
			log.Fatal(err)
		}

		err = bam.MakeBam(outFile)
		if err != nil {
			log.Fatal(err)
		}
	}
}


