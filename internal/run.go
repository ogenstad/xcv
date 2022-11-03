package internal

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func Run() {
	version := flag.Bool("v", false, "Display version")

	flag.Parse()

	if *version {
		fmt.Printf("xcv %s\n", Version)
		os.Exit(0)
	}

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	var certificates []Certificate

	err = ProcessCertificates(&certificates, bytes)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	output := ""
	for i := range certificates {
		output += FormatCertificate(certificates[i])
	}

	if len(output) == 0 {
		fmt.Println("Unable to parse certificate from provided input")
		os.Exit(1)
	}

	fmt.Println(output)
}
