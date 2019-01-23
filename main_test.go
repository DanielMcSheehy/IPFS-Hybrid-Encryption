package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestUploadStringToIPFS(t *testing.T) {
	hashAddress, _ := uploadToIPFS(strings.NewReader("testing a string"))
	if len(hashAddress) < len("QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j") {
		t.Error("Expecting a valid hash address")
	}
}

func TestRetrieveFromIPFS(t *testing.T) {
	hashAddress, _ := uploadToIPFS(strings.NewReader("testing a string"))

	if len(hashAddress) < len("QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j") {
		t.Error("Expecting a valid hash address")
	}
	fileReader, err := retreiveFromIPFS(hashAddress)
	if (err != nil) && (fileReader == nil) {
		t.Error("Cannot retrieve data from IPFS")
	}
}

func TestUploadEncryptedStringToIPFS(t *testing.T) {
	testingBytes := strings.NewReader("testing a string")
	publicKey, privateKey := generateKeys()

	encryptedBytes := encrypt(publicKey, testingBytes)

	hashAddress, _ := uploadToIPFS(bytes.NewReader(encryptedBytes))
	if len(hashAddress) < len("QmTp2hEo8eXRp6wg7jXv1BLCMh5a4F3B7buAUZNZUu772j") {
		t.Error("Expecting a valid hash address")
	}
	storedEncryptedBytes, err := retreiveFromIPFS(hashAddress)
	if err != nil {
		t.Error("Cannot retrieve data from IPFS")
	}
	decryptedReader := decrypt(privateKey, storedEncryptedBytes)

	decryptedString := readerToString(decryptedReader)

	if decryptedString != "testing a string" {
		t.Error("Expecting a string to be encrypted and decrypted succesfully")
	}
}

func TestFileEncryptionUploadAndDownload(t *testing.T) {
	fileReader := uploadFile("example-file.js")
	publicKey, privateKey := generateKeys()
	encryptedBytes := encrypt(publicKey, fileReader)

	hashAddress, _ := uploadToIPFS(bytes.NewReader(encryptedBytes))
	storedEncryptedReader, err := retreiveFromIPFS(hashAddress)
	t.Log(storedEncryptedReader)
	if err != nil {
		t.Error("Cannot retrieve data from IPFS")
	}

	decryptedReader := decrypt(privateKey, storedEncryptedReader)
	download(decryptedReader, "outputTest.js")

	if readerToString(decryptedReader) != readerToString(fileReader) {
		t.Error("file encryption and decryption is not correct")
	}
}
