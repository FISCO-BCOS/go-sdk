package utils

import (
	"crypto/rand"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"time"

	bcx509 "github.com/liuxinfeng96/bc-crypto/x509"
)

const DefaultFiscoBcosCertOrganization = "fisco-bcos"

const (
	DefaultCertificateCountry  = "CN"
	DefaultCertificateProvince = "GuangDong"
	DefaultCertificateLocality = "ShenZhen"
)

// GenerateCertificateSigningRequest
// create the CSR with X.509
// return the CSR byte array with PEM
func GenerateCertificateSigningRequest(privateKeyBytes []byte,
	role, certName string) ([]byte, error) {

	sk, err := bcx509.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	signatureAlgorithm, err := bcx509.GetSignatureAlgorithm(sk)
	if err != nil {
		return nil, err
	}

	templateX509 := &bcx509.CertificateRequest{
		SignatureAlgorithm: signatureAlgorithm,
		Subject: pkix.Name{
			Country:            []string{DefaultCertificateCountry},
			Locality:           []string{DefaultCertificateLocality},
			Province:           []string{DefaultCertificateProvince},
			OrganizationalUnit: []string{role},
			Organization:       []string{DefaultFiscoBcosCertOrganization},
			CommonName:         certName,
		},
	}

	data, err := bcx509.CreateCertificateRequest(rand.Reader, templateX509, sk)
	if err != nil {
		return nil, err
	}

	pemCSR := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE REQUEST", Bytes: data})

	return pemCSR, nil
}

// GenerateCertificateSelf
// create a self-signed digital certificate with X.509
// return the certificate byte array with PEM
func GenerateCertificateSelf(privateKeyBytes, csrBytes []byte,
	days int) ([]byte, error) {

	csr, err := bcx509.ParseCertificateRequestFromPEM(csrBytes)
	if err != nil {
		return nil, err
	}

	if err = csr.CheckSignature(); err != nil {
		return nil, err
	}

	keyUsage, extKeyUsage := bcx509.GetKeyUsageAndExtKeyUsage(true)

	notBefore := time.Now().UTC()

	sn, err := rand.Int(rand.Reader, big.NewInt(1<<62))
	if err != nil {
		return nil, err
	}

	caKey, err := bcx509.ParsePrivateKey(privateKeyBytes)
	if err != nil {
		return nil, err
	}

	signatureAlgorithm, err := bcx509.GetSignatureAlgorithm(caKey)
	if err != nil {
		return nil, err
	}

	template := &bcx509.Certificate{
		SerialNumber:          sn,
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(time.Hour * 24 * time.Duration(days)).UTC(),
		BasicConstraintsValid: true,
		IsCA:                  true,
		KeyUsage:              keyUsage,
		ExtKeyUsage:           extKeyUsage,
		DNSNames:              csr.DNSNames,
		Subject:               csr.Subject,
		Extensions:            csr.Extensions,
		ExtraExtensions:       csr.ExtraExtensions,
		SignatureAlgorithm:    signatureAlgorithm,
		PublicKeyAlgorithm:    csr.PublicKeyAlgorithm,
		PublicKey:             csr.PublicKey,
	}

	template.SubjectKeyId, err = bcx509.ComputeSKI(csr.PublicKey)
	if err != nil {
		return nil, err
	}

	templateOfIssuer := new(bcx509.Certificate)
	templateOfIssuer = template

	certDER, err := bcx509.CreateCertificate(rand.Reader, template, templateOfIssuer,
		csr.PublicKey, caKey)
	if err != nil {
		return nil, err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	return certPEM, nil

}

// GenerateCertificate
// create a certificate with X.509
// return the certificate byte array with PEM
func GenerateCertificate(caCertBytes, caKeyBytes, csrBytes []byte, isCa bool,
	days int) ([]byte, error) {
	csr, err := bcx509.ParseCertificateRequestFromPEM(csrBytes)
	if err != nil {
		return nil, err
	}

	if err = csr.CheckSignature(); err != nil {
		return nil, err
	}

	keyUsage, extKeyUsage := bcx509.GetKeyUsageAndExtKeyUsage(isCa)

	notBefore := time.Now().UTC()

	sn, err := rand.Int(rand.Reader, big.NewInt(1<<62))
	if err != nil {
		return nil, err
	}

	caKey, err := bcx509.ParsePrivateKey(caKeyBytes)
	if err != nil {
		return nil, err
	}

	signatureAlgorithm, err := bcx509.GetSignatureAlgorithm(caKey)
	if err != nil {
		return nil, err
	}

	template := &bcx509.Certificate{
		SerialNumber:          sn,
		NotBefore:             notBefore,
		NotAfter:              notBefore.Add(time.Hour * 24 * time.Duration(days)).UTC(),
		BasicConstraintsValid: true,
		IsCA:                  isCa,
		KeyUsage:              keyUsage,
		ExtKeyUsage:           extKeyUsage,
		DNSNames:              csr.DNSNames,
		Subject:               csr.Subject,
		Extensions:            csr.Extensions,
		ExtraExtensions:       csr.ExtraExtensions,
		SignatureAlgorithm:    signatureAlgorithm,
		PublicKeyAlgorithm:    csr.PublicKeyAlgorithm,
		PublicKey:             csr.PublicKey,
	}

	template.SubjectKeyId, err = bcx509.ComputeSKI(csr.PublicKey)
	if err != nil {
		return nil, err
	}

	certOfIssuer, err := bcx509.ParseCertificateFromPEM(caCertBytes)
	if err != nil {
		return nil, err
	}

	template.Issuer = certOfIssuer.Subject

	if certOfIssuer.SubjectKeyId != nil {
		template.AuthorityKeyId = certOfIssuer.SubjectKeyId
	} else {
		template.AuthorityKeyId, err = bcx509.ComputeSKI(certOfIssuer.PublicKey)
		if err != nil {
			return nil, err
		}
	}

	certDER, err := bcx509.CreateCertificate(rand.Reader, template, certOfIssuer,
		csr.PublicKey, caKey)
	if err != nil {
		return nil, err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	return certPEM, nil
}
