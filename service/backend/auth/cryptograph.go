package auth

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

// Encrypt the file
func EncryptFile(inputFilePath, outputFilePath string, secretKey []byte) error {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("unable to open input file: %w", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("unable to create output file: %w", err)
	}
	defer outputFile.Close()

	// Generate an AES cipher block with the secret key
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return fmt.Errorf("unable to create AES cipher: %w", err)
	}

	// Generate an IV (Initialization Vector) for encryption
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return fmt.Errorf("unable to generate IV: %w", err)
	}

	if _, err := outputFile.Write(iv); err != nil {
		return fmt.Errorf("unable to write IV to file: %w", err)
	}

	// Create a cipher stream with the AES block and the IV
	stream := cipher.NewCFBEncrypter(block, iv)

	writer := &cipher.StreamWriter{S: stream, W: outputFile}
	_, err = io.Copy(writer, inputFile)
	if err != nil {
		return fmt.Errorf("error while encrypting file: %w", err)
	}

	return nil
}

// Decrypt the file
func DecryptFile(inputFilePath, outputFilePath string, secretKey []byte) error {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return fmt.Errorf("unable to open input file: %w", err)
	}
	defer inputFile.Close()

	// Read the IV (Initialization Vector) from the input file
	iv := make([]byte, aes.BlockSize)
	if _, err := inputFile.Read(iv); err != nil {
		return fmt.Errorf("unable to read IV from file: %w", err)
	}

	// Create an AES cipher block with the secret key
	block, err := aes.NewCipher(secretKey)
	if err != nil {
		return fmt.Errorf("unable to create AES cipher: %w", err)
	}

	// Create a cipher stream with the AES block and the IV
	stream := cipher.NewCFBDecrypter(block, iv)

	// Create the output file
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("unable to create output file: %w", err)
	}
	defer outputFile.Close()

	// Decrypt the file
	reader := &cipher.StreamReader{S: stream, R: inputFile}
	_, err = io.Copy(outputFile, reader)
	if err != nil {
		return fmt.Errorf("error while decrypting file: %w", err)
	}

	return nil
}
