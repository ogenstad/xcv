package internal

import (
	"fmt"
	"io"
	"os"
)

func Run() {
	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	var certificates []Certificate

	err = ProcessCertificates(&certificates, bytes)
	if err != nil {
		panic(err)
	}

	output := ""
	for i := range certificates {
		output += FormatCertificate(certificates[i])
	}

	fmt.Println(output)
}
