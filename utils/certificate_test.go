package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateCertificateSigningRequest(t *testing.T) {
	sk, err := GenerateSecp256k1Key()
	require.Nil(t, err)

	// sk, err := GenerateSm2Key()
	// require.Nil(t, err)

	// sk, err := GenerateRSA2048Key()
	// require.Nil(t, err)

	csrBytes, err := GenerateCertificateSigningRequest(sk, "node", "node1")
	require.Nil(t, err)

	fmt.Println(string(csrBytes))
}

func TestCreateCertificateSelf(t *testing.T) {
	sk, err := GenerateSecp256k1Key()
	require.Nil(t, err)

	// sk, err := GenerateSm2Key()
	// require.Nil(t, err)

	// sk, err := GenerateRSA2048Key()
	// require.Nil(t, err)

	csrBytes, err := GenerateCertificateSigningRequest(sk, "group-ca", "ca1")
	require.Nil(t, err)

	caBytes, err := GenerateCertificateSelf(sk, csrBytes, 365)
	require.Nil(t, err)

	fmt.Println(string(caBytes))

}

func TestCreateCertificate(t *testing.T) {
	// caSk, err := GenerateSecp256k1Key()
	// require.Nil(t, err)

	caSk, err := GenerateSm2Key()
	require.Nil(t, err)

	// caSk, err := GenerateRSA2048Key()
	// require.Nil(t, err)

	caCsrBytes, err := GenerateCertificateSigningRequest(caSk, "group-ca", "ca1")
	require.Nil(t, err)

	caBytes, err := GenerateCertificateSelf(caSk, caCsrBytes, 365)
	require.Nil(t, err)

	// sk, err := GenerateSecp256k1Key()
	// require.Nil(t, err)

	// sk, err := GenerateSm2Key()
	// require.Nil(t, err)

	sk, err := GenerateRSA2048Key()
	require.Nil(t, err)

	csrBytes, err := GenerateCertificateSigningRequest(sk, "node", "node1")
	require.Nil(t, err)

	certBytes, err := GenerateCertificate(caBytes, caSk, csrBytes, false, 365)
	require.Nil(t, err)

	fmt.Println(string(certBytes))
}
