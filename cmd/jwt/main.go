package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	flagAlg     = flag.String("alg", "", algHelp())
	flagKey     = flag.String("key", "", "path to key file or '-' to read from stdin")
	flagCompact = flag.Bool("compact", false, "output compact JSON")
	flagDebug   = flag.Bool("debug", false, "print out all kinds of debug data")
	flagClaims  = make(ArgList)
	flagHead    = make(ArgList)

	flagSign   = flag.String("sign", "", "path to claims file to sign, '-' to read from stdin, or '+' to use only -claim args")
	flagVerify = flag.String("verify", "", "path to JWT token file to verify or '-' to read from stdin")
	flagShow   = flag.String("show", "", "path to JWT token file to show without verification or '-' to read from stdin")
)

func main() {

	flag.Var(flagClaims, "claim", "add additional claims. may be used more than once")
	flag.Var(flagHead, "header", "add additional header params. may be used more than once")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  One of the following flags is required: sign, verify or show\n")
		flag.PrintDefaults()
	}

	flag.Parse()

	if err := start(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func start() error { _ = "STUB: not implemented"; return nil }

func loadData(p string) (_ []byte, retErr error) { _ = "STUB: not implemented"; return nil, nil }

func printJSON(j any) error { _ = "STUB: not implemented"; return nil }

func verifyToken() error { _ = "STUB: not implemented"; return nil }

func signToken() error { _ = "STUB: not implemented"; return nil }

func showToken() error { _ = "STUB: not implemented"; return nil }

func isEs() bool { _ = "STUB: not implemented"; return false }

func isRs() bool { _ = "STUB: not implemented"; return false }

func isEd() bool { _ = "STUB: not implemented"; return false }

func isNone() bool { _ = "STUB: not implemented"; return false }

func algHelp() string { _ = "STUB: not implemented"; return "" }

type ArgList map[string]string

func (l ArgList) String() string { _ = "STUB: not implemented"; return "" }

func (l ArgList) Set(arg string) error { _ = "STUB: not implemented"; return nil }
