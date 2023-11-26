package easy

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/hex"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// HashAndToken generates a bcrypt hash and a correlated 16-digit token based on the given password.
func HashAndToken(password string) (string, string, error) {
	// Generate bcrypt hash
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", "", err
	}

	// Convert the hash to a hex-encoded string
	hashString := hex.EncodeToString(hash)

	// Generate a 4-byte (8-character) random component
	randomComponent, err := generateRandomBytes(4)
	if err != nil {
		return "", "", err
	}

	// Combine hash and random component to create a 16-digit token
	token := hashString + fmt.Sprintf("%08x", randomComponent)

	return hashString, token, nil
}

// generateRandomBytes generates n random bytes.
func generateRandomBytes(n int) ([]byte, error) {
	randomBytes := make([]byte, n)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

// Authentication
func AuthHash(hash, password string) error {

	stored, err := hex.DecodeString(hash)
	if err != nil {
		fmt.Println(err)
	}

	return bcrypt.CompareHashAndPassword(stored, []byte(password))

}

// Transforms the DB connection string into a pem file for safekeeping
func MakeSecret() error {
	filename := "/etc/mhdev/keychain/secret.pem"
	data := "host=192.168.50.40 port=5432 user=pgremote password=pgremoteuser database=postgres sslmode=require"
	// Create or open the file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode string data to PEM format
	stringPEM := &pem.Block{
		Type:  "DATA",
		Bytes: []byte(data),
	}

	// Write PEM data to the file
	err = pem.Encode(file, stringPEM)
	if err != nil {
		return err
	}
	return nil
}

// Gets the connection string from the pem
func GetConn() (string, error) {
	filename := "/etc/mhdev/keychain/secret.pem"
	// Read the entire file
	fileData, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Decode PEM block
	block, _ := pem.Decode(fileData)
	if block == nil {
		return "", fmt.Errorf("encrypt - poad .pem failed")
	}

	// Extract the string data
	data := string(block.Bytes)

	return data, nil
}

// Creates self signed certs
func GenerateCerts() error {
	var certPath = "/etc/mhdev/keychain/tls/CA.crt"
	var keyPath = "/etc/mhdev/keychain/tls/secret/CA.key"
	var validityDays = 10000
	priv, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return err
	}

	notBefore := time.Now()
	notAfter := notBefore.Add(time.Duration(validityDays) * 24 * time.Hour)

	serialNumber, err := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))
	if err != nil {
		return err
	}

	template := x509.Certificate{
		SerialNumber: serialNumber,
		Subject: pkix.Name{
			Organization: []string{"Self-signed Certificate"},
		},
		NotBefore:             notBefore,
		NotAfter:              notAfter,
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &priv.PublicKey, priv)
	if err != nil {
		return err
	}

	certFile, err := os.Create(certPath)
	if err != nil {
		return err
	}
	defer certFile.Close()

	err = pem.Encode(certFile, &pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	if err != nil {
		return err
	}

	keyFile, err := os.Create(keyPath)
	if err != nil {
		return err
	}
	defer keyFile.Close()

	privBytes, err := x509.MarshalPKCS8PrivateKey(priv)
	if err != nil {
		return err
	}

	err = pem.Encode(keyFile, &pem.Block{Type: "PRIVATE KEY", Bytes: privBytes})
	if err != nil {
		return err
	}

	return nil
}
