package main

import (
	"flag"
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	"log"
	"sync"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// CLI is the command line object
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run invokes the CLI with the given arguments.
func (cli *CLI) Run(args []string) int {
	var (
		d       string
		version bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(Name, flag.ContinueOnError)
	flags.SetOutput(cli.errStream)

	flags.StringVar(&d, "destination", "", "Destination")
	flags.StringVar(&d, "d", "", "(Short)")

	flags.BoolVar(&version, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if version {
		fmt.Fprintf(cli.errStream, "%s version %s\n", Name, Version)
		return ExitCodeOK
	}

	_ = d

	if d == "" {
		fmt.Println("set destination with 'd' flag")
		return ExitCodeError
	}

	if len(flags.Args()) < 1 {
		log.Fatal("required destination and file names")
		return ExitCodeError
	}
	dst := d
	var tracks []Track

	for _, arg := range flags.Args() {
		t, err := NewTrack(arg)
		if err != nil {
			return ExitCodeError
		}
		tracks = append(tracks, t)
	}

	bar := pb.New(len(tracks))
	bar.Start()

	wg := new(sync.WaitGroup)
	for _, track := range tracks {
		wg.Add(1)
		go func(track Track) {
			defer wg.Done()
			defer bar.Increment()
			track.TransferTo(dst)
		}(track)
	}
	wg.Wait()
	bar.FinishPrint("Complete Transfer")

	return ExitCodeOK
}
