package main

import (
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"os"

	"math/big"
	"time"

	"sigs.k8s.io/cluster-api/util/certs"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	privKey, err := certs.NewPrivateKey()
	check(err)

	// https://github.com/kubernetes-sigs/cluster-api/blob/25a933a307b1495edc13f6bf7186c9d3328415e8/util/secret/certificates.go#L433-L456
	now := time.Now().UTC()

	tmpl := x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(0),
		Subject: pkix.Name{
			CommonName: "kubernetes",
		},
		NotBefore:             now.Add(time.Minute * -5),
		NotAfter:              now.Add(time.Hour * 24 * 365 * 100), // 100 years
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		MaxPathLenZero:        true,
		BasicConstraintsValid: true,
		MaxPathLen:            0,
		IsCA:                  true,
	}

	byteCert, err := x509.CreateCertificate(rand.Reader, &tmpl, &tmpl, privKey.Public(), privKey)
	check(err)

	x509Cert, err := x509.ParseCertificate(byteCert)
	check(err)

	pubKey, err := certs.EncodePublicKeyPEM(&privKey.PublicKey)
	check(err)

	keyFile, err := os.Create("tls.key")
	check(err)
	defer keyFile.Close()

	crtFile, err := os.Create("tls.crt")
	check(err)
	defer crtFile.Close()

	pubFile, err := os.Create("tls.pub")
	check(err)
	defer pubFile.Close()

	keyFile.Write(certs.EncodePrivateKeyPEM(privKey))
	crtFile.Write(certs.EncodeCertPEM(x509Cert))
	pubFile.Write(pubKey)
}
