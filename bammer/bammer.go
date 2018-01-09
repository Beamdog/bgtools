package main

import (
	"flag"
	"image"
	"image/gif"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Beamdog/bgfileformats"
	//	"github.com/davecheney/profile"
)

var input = flag.String("input", "input.bam", "Source [bam, bamd, gif]")
var output = flag.String("output", "out", "Output Directory/Filename")
var palette = flag.String("palette", "palette.png", "Palette to use for the bam")
var mirror = flag.Bool("mirror", false, "Set to true to mirror all the frames")
var offset_x = flag.Int("offsetx", 0, "Offset all exported frames in X direction")
var offset_y = flag.Int("offsety", 0, "Offset all exported frames in Y direction")
var mode = flag.String("mode", "bamd", "Output format[bamd, gif]")

func main() {
	flag.Parse()
	//defer profile.Start(profile.CPUProfile).Stop()

	bamFileIn, err := os.Open(filepath.Clean(*input))
	if err != nil {
		log.Fatal(err)
	}
	defer bamFileIn.Close()
	if strings.ToLower(filepath.Ext(*input)) == ".bam" {

		bamIn, err := bg.OpenBAM(bamFileIn, nil)
		if err != nil {
			log.Fatal(err)
		}
		outputName := strings.TrimSuffix(filepath.Base(*output), filepath.Ext(*output))
		if *mode == "gif" {
			err := bamIn.MakeGif(*output, outputName)
			if err != nil {
				log.Fatal(err)
			}
		} else if *mode == "bamd" {
			bamIn.MakeBamd(filepath.Clean(*output), outputName, *mirror, *offset_x, *offset_y)
		} else {
			log.Fatal("Unknown output mode: %s\n", *mode)
		}

	} else if strings.ToLower(filepath.Ext(*input)) == ".bamd" {
		//Open our output path first
		outBam, err := os.Create(*output)
		if err != nil {
			log.Fatal(err)
		}
		defer outBam.Close()

		//Chdir to the root of our .bamd so our paths are consistent
		bamdRoot := filepath.Dir(*input)
		os.Chdir(bamdRoot)
		if err != nil {
			log.Fatal(err)
		}
		bamOut, err := bg.OpenBAMD(bamFileIn, *palette)
		if err != nil {
			log.Fatal(err)
		}

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
		sequences[0] = image.Pt(0, len(gif.Image))
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
