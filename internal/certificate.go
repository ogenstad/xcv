package internal

import (
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"time"

	"golang.org/x/crypto/cryptobyte"
)

type Certificate struct {
	cert                  *x509.Certificate
	Subject               pkix.Name
	Issuer                pkix.Name
	Version               int
	SerialNumber          *big.Int
	PublicKey             any
	NotBefore             time.Time
	NotAfter              time.Time
	KeyUsage              []string
	ExtendedKeyUsage      []ExtendedKeyUsage
	UnsupportedExtensions []string
	UnsupportedProperties bool
}

type ExtendedKeyUsage struct {
	Oid  string
	Name string
}

func (c *Certificate) init() error {
	err := c.processExtensions()
	if err != nil {
		return fmt.Errorf("unable to parse certificate extensions: %w", err)
	}

	if c.cert.KeyUsage != 0 {
		c.processKeyUsage()
	}

	return nil
}

func (c *Certificate) processKeyUsage() {
	if c.cert.KeyUsage&x509.KeyUsageDigitalSignature != 0 {
		c.KeyUsage = append(c.KeyUsage, "Digital signature")
	}

	if c.cert.KeyUsage&x509.KeyUsageContentCommitment != 0 {
		c.KeyUsage = append(c.KeyUsage, "Non-repudiation")
	}

	if c.cert.KeyUsage&x509.KeyUsageKeyEncipherment != 0 {
		c.KeyUsage = append(c.KeyUsage, "Key encipherment")
	}

	if c.cert.KeyUsage&x509.KeyUsageDataEncipherment != 0 {
		c.KeyUsage = append(c.KeyUsage, "Data encipherment")
	}

	if c.cert.KeyUsage&x509.KeyUsageKeyAgreement != 0 {
		c.KeyUsage = append(c.KeyUsage, "Key agreement")
	}

	if c.cert.KeyUsage&x509.KeyUsageCertSign != 0 {
		c.KeyUsage = append(c.KeyUsage, "Certificate signing")
	}

	if c.cert.KeyUsage&x509.KeyUsageCRLSign != 0 {
		c.KeyUsage = append(c.KeyUsage, "CRL signing")
	}

	if c.cert.KeyUsage&x509.KeyUsageEncipherOnly != 0 {
		c.KeyUsage = append(c.KeyUsage, "Encipher only")
	}

	if c.cert.KeyUsage&x509.KeyUsageDecipherOnly != 0 {
		c.KeyUsage = append(c.KeyUsage, "Decipher only")
	}
}

func (c *Certificate) processExtensions() error {
	for _, extension := range c.cert.Extensions {
		objectIdentifier := extension.Id.String()
		switch objectIdentifier {
		case "2.5.29.14", "2.5.29.15", "2.5.29.19", "2.5.29.31", "2.5.29.35":
			// Subject Key Identifier (SKI) is ignored, handled elsewhere
			// Key Usage is ignored, handled elsewhere
			// Basic constraints is ignored, handled elsewhere
			// CRL Distribution Points (CDP) are ignored, handled elsewhere
			// Authority Key Identifier (AKI) is ignored, handled elsewhere
		case "2.5.29.37":
			usages, foundUnsupported, err := parseExtendedKeyUsage(cryptobyte.String(extension.Value))
			if err != nil {
				return err
			}

			if foundUnsupported {
				c.UnsupportedProperties = true
			}

			c.ExtendedKeyUsage = usages
		default:
			c.UnsupportedExtensions = append(c.UnsupportedExtensions, objectIdentifier)
			c.UnsupportedProperties = true
		}
	}

	return nil
}

func New(crt *x509.Certificate) (Certificate, error) {
	certificate := Certificate{
		cert:                  crt,
		Subject:               crt.Subject,
		Issuer:                crt.Issuer,
		Version:               crt.Version,
		SerialNumber:          crt.SerialNumber,
		PublicKey:             crt.PublicKey,
		NotBefore:             crt.NotBefore,
		NotAfter:              crt.NotAfter,
		KeyUsage:              []string{},
		ExtendedKeyUsage:      []ExtendedKeyUsage{},
		UnsupportedExtensions: []string{},
		UnsupportedProperties: false,
	}

	err := certificate.init()
	if err != nil {
		return Certificate{}, fmt.Errorf("unable to parse certificate: %w", err)
	}

	return certificate, nil
}

func ProcessCertificates(certificates *[]Certificate, certData []byte) error {
	block, rest := pem.Decode(certData)
	if block == nil {
		return nil
	}

	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return fmt.Errorf("unable to parse certificate: %w", err)
	}

	crt, err := New(cert)
	if err != nil {
		return fmt.Errorf("unable to initialize certificate: %w", err)
	}

	*certificates = append(*certificates, crt)

	if len(rest) > 0 {
		err = ProcessCertificates(certificates, rest)
		if err != nil {
			return fmt.Errorf("error while processing certificate(s): %w", err)
		}
	}

	return nil
}
