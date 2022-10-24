package internal

import (
	"encoding/asn1"
	"errors"
	"fmt"

	"golang.org/x/crypto/cryptobyte"
	cryptobyte_asn1 "golang.org/x/crypto/cryptobyte/asn1"
)

var (
	ErrParseASNFromEKU    = errors.New("unable to read asn1 data from extended key usage")
	ErrParseASNOidFromEKU = errors.New("unable to parse oid from extended key usage")
)

func parseExtendedKeyUsage(der cryptobyte.String) ([]ExtendedKeyUsage, bool, error) {
	var usages []ExtendedKeyUsage

	foundUnsupported := false

	ekuMap := map[string]string{
		"1.3.6.1.5.5.7.3.1": "Server Authentication",
		"1.3.6.1.5.5.7.3.2": "Client Authentication",
	}

	if !der.ReadASN1(&der, cryptobyte_asn1.SEQUENCE) {
		return usages, false, ErrParseASNFromEKU
	}

	for !der.Empty() {
		var eku asn1.ObjectIdentifier
		if !der.ReadASN1ObjectIdentifier(&eku) {
			return usages, false, ErrParseASNOidFromEKU
		}

		oid := eku.String()

		var usage ExtendedKeyUsage

		if v, found := ekuMap[oid]; found {
			usage = ExtendedKeyUsage{Oid: oid, Name: v}
		} else {
			usage = ExtendedKeyUsage{Oid: oid, Name: fmt.Sprintf("Undefined (%s)", oid)}
			foundUnsupported = true
		}

		usages = append(usages, usage)
	}

	return usages, foundUnsupported, nil
}
