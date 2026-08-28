package cryptoutil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

const nonceSize = 16

func deriveKey(passphrase string) []byte {
	h := sha256.Sum256([]byte(passphrase))
	return h[:]
}

func EncryptReader(r io.Reader, passphrase string) (io.Reader, error) {
	if passphrase == "" {
		return r, nil
	}

	key := deriveKey(passphrase)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating cipher: %w", err)
	}

	nonce := make([]byte, nonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return nil, fmt.Errorf("error generating nonce: %w", err)
	}

	stream := cipher.NewCTR(block, nonce)
	nonceReader := bytes.NewReader(nonce)
	encryptedReader := cipher.StreamReader{S: stream, R: r}

	return io.MultiReader(nonceReader, encryptedReader), nil
}

func DecryptFile(path string, passphrase string) error {
	if passphrase == "" {
		return nil
	}

	key := deriveKey(passphrase)

	encFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening encrypted file: %w", err)
	}
	defer encFile.Close()

	nonce := make([]byte, nonceSize)
	if _, err := io.ReadFull(encFile, nonce); err != nil {
		return fmt.Errorf("error reading nonce: %w", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("error creating cipher: %w", err)
	}
	stream := cipher.NewCTR(block, nonce)

	tmpPath := path + ".tmp"
	decFile, err := os.Create(tmpPath)
	if err != nil {
		return fmt.Errorf("error creating temp file: %w", err)
	}

	decryptedReader := cipher.StreamReader{S: stream, R: encFile}
	if _, err := io.Copy(decFile, decryptedReader); err != nil {
		decFile.Close()
		os.Remove(tmpPath)
		return fmt.Errorf("error decrypting file: %w", err)
	}

	decFile.Close()
	encFile.Close()

	if err := os.Rename(tmpPath, path); err != nil {
		os.Remove(tmpPath)
		return fmt.Errorf("error replacing file with decrypted version: %w", err)
	}

	return nil
}
