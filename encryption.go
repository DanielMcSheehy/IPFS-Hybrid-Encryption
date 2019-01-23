package main

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/pivotal-cf-experimental/bletchley"
)

func generateKeys() (*rsa.PublicKey, *rsa.PrivateKey) {
	privateKey, publicKey, err := bletchley.Generate()
	if err != nil {
		log.Fatal(err)
	}
	return publicKey, privateKey
}

func encrypt(publicKey *rsa.PublicKey, fileReader io.Reader) []byte {
	encryptedMessage, err := bletchley.Encrypt(publicKey, readerToByteArray(fileReader))
	if err != nil {
		fmt.Println(err)
	}
	encryptedBytes, err := json.Marshal(encryptedMessage)
	return encryptedBytes
}

func decrypt(privateKey *rsa.PrivateKey, storedEncryptedReader io.Reader) io.Reader {
	var encryptedMessage bletchley.EncryptedMessage
	err := json.Unmarshal(readerToByteArray(storedEncryptedReader), &encryptedMessage)
	if err != nil {
		log.Fatal(err)
	}
	plaintextBytes, err := bletchley.Decrypt(privateKey, encryptedMessage)
	return bytes.NewReader(plaintextBytes)
}

func readerToByteArray(reader io.Reader) []byte {
	body, _ := ioutil.ReadAll(reader)
	return body
}

func readerToString(reader io.Reader) string {
	return string(readerToByteArray(reader)[:])
}
