package types

import (
	. "crypto/x509"
	"crypto/x509/pkix"
	"math/big"
	"time"
)

var (
	SignatureAlgorithms = map[SignatureAlgorithm]string{
		UnknownSignatureAlgorithm: "UNKNOWN",
		MD2WithRSA:                "RSA-MD2",
		MD5WithRSA:                "RSA-MD5",
		SHA1WithRSA:               "RSA-SHA1",
		SHA256WithRSA:             "RSA-SHA256",
		SHA384WithRSA:             "RSA-SHA384",
		SHA512WithRSA:             "RSA-SHA512",
		DSAWithSHA1:               "DSA-SHA512",
		DSAWithSHA256:             "DSA-SHA512",
		ECDSAWithSHA1:             "ECDSA-SHA1",
		ECDSAWithSHA256:           "ECDSA-SHA256",
		ECDSAWithSHA384:           "ECDSA-SHA384",
		ECDSAWithSHA512:           "ECDSA-SHA512",
		SHA256WithRSAPSS:          "RSAPSS-SHA256",
		SHA384WithRSAPSS:          "RSAPSS-SHA384",
		SHA512WithRSAPSS:          "RSAPSS-SHA512",
		PureEd25519:               "ED25519",
	}

	PublicKeyAlgorithms = map[PublicKeyAlgorithm]string{
		UnknownPublicKeyAlgorithm: "UNKNOWN",
		RSA:                       "RSA",
		DSA:                       "DSA",
		ECDSA:                     "ECDSA",
		Ed25519:                   "ED25519",
	}

	KeyUsages = map[KeyUsage]string{
		KeyUsageDigitalSignature:  "DIGITAL_SIGNATURE",
		KeyUsageContentCommitment: "CONTENT_COMMITMENT",
		KeyUsageKeyEncipherment:   "KEY_ENCIPHERMENT",
		KeyUsageDataEncipherment:  "DATA_ENCIPHERMENT",
		KeyUsageKeyAgreement:      "KEY_AGREEMENT",
		KeyUsageCertSign:          "CERT_SIGN",
		KeyUsageCRLSign:           "CRL_SIGN",
		KeyUsageEncipherOnly:      "ENCIPHER_ONLY",
		KeyUsageDecipherOnly:      "DECIPHER_ONLY",
	}

	ExtKeyUsages = map[ExtKeyUsage]string{
		ExtKeyUsageAny:                            "ANY",
		ExtKeyUsageServerAuth:                     "SERVER_AUTH",
		ExtKeyUsageClientAuth:                     "CLIENT_AUTH",
		ExtKeyUsageCodeSigning:                    "CODE_SIGNING",
		ExtKeyUsageEmailProtection:                "EMAIL_PROTECTION",
		ExtKeyUsageIPSECEndSystem:                 "IPSEC_END_SYSTEM",
		ExtKeyUsageIPSECTunnel:                    "IPSEC_TUNNEL",
		ExtKeyUsageIPSECUser:                      "IPSEC_USER",
		ExtKeyUsageTimeStamping:                   "TIME_STAMPING",
		ExtKeyUsageOCSPSigning:                    "OCSP_SIGNING",
		ExtKeyUsageMicrosoftServerGatedCrypto:     "MS_SERVER_GATED_CRYPTO",
		ExtKeyUsageNetscapeServerGatedCrypto:      "NETSCAPE_SERVER_GATED_CRYPTO",
		ExtKeyUsageMicrosoftCommercialCodeSigning: "MS_COMMERICIAL_CODE_SIGNING",
		ExtKeyUsageMicrosoftKernelCodeSigning:     "MS_KERNEL_CODE_SIGNING",
	}
)

type PKIX struct {
	values []pkix.AttributeTypeAndValue
}

type CriticalPKIX struct {
	values []pkix.Extension
}

type TLSCertificate struct {
	Signature          []byte
	SignatureAlgorithm SignatureAlgorithm
	PublicKeyAlgorithm PublicKeyAlgorithm
	PublicKey          interface{}
	Version            int
	SerialNumber       *big.Int

	Issuer      PKIX
	Subject     PKIX
	NotBefore   time.Time
	NotAfter    time.Time
	KeyUsage    KeyUsage
	ExtKeyUsage ExtKeyUsage

	Extensions CriticalPKIX
}

func NewCert(cert *Certificate) *TLSCertificate {
	return nil
}
