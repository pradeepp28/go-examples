package main

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func main() {
	privKey, _ := rsa.GenerateKey(rand.Reader, 2048)
	var buf bytes.Buffer
	b, _ := x509.MarshalPKIXPublicKey(privKey.Public())

	fmt.Println(base64.StdEncoding.EncodeToString(b))
	_ = pem.Encode(&buf, &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: b,
	})
	fmt.Println(buf.String())
}
