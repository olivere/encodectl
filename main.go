package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/olivere/encodectl/transformer"
)

func main() {
	var (
		infile      = flag.String("i", "", "Input file")
		inencoding  = flag.String("ie", "utf-8", "Input file encoding")
		outfile     = flag.String("o", "", "Output file")
		outencoding = flag.String("oe", "utf-8", "Output file encoding")
	)
	flag.Parse()
	log.SetFlags(0)

	if *inencoding == "" {
		log.Fatal("no input file encoding specified (-ie=...)")
	}
	if *outencoding == "" {
		log.Fatal("no output file encoding specified (-oe=...)")
	}

	var r io.Reader
	if *infile != "" {
		f, err := os.Open(*infile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		r = f
	} else {
		r = os.Stdin
	}

	var w io.Writer
	if *outfile != "" {
		f, err := os.Create(*outfile)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		w = f
	} else {
		w = os.Stdout
	}

	t, err := transformer.New(r, w, *inencoding, *outencoding)
	if err != nil {
		log.Fatal(err)
	}

	if err := t.Transform(); err != nil {
		log.Fatal(err)
	}
}
