package main

import (
	"encoding/base64"
	"flag"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/snappy"
)

var (
	isDecode  bool
	outBinary bool
)

func init() {
	flag.BoolVar(&isDecode, "d", false, "Decode from stdin")
	flag.BoolVar(&outBinary, "b", false, "Output compressed data as binary instead of a base64-encoded string")
	flag.Parse()
}

func main() {
	var (
		out []byte
		err error
	)
	r, _ := ioutil.ReadAll(os.Stdin)
	if isDecode {
		// First attempt to decode base64
		out, err = base64.StdEncoding.DecodeString(string(r))
		if err != nil {
			// Assume it isn't base64 and just move on
			out = r
		}
		out, err = snappy.Decode(nil, out)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		out = snappy.Encode(nil, r)
		if !outBinary {
			out = []byte(base64.StdEncoding.EncodeToString(out))
		}
	}
	os.Stdout.Write(out)
}
