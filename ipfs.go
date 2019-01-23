package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func uploadToIPFS(file io.Reader) (string, error) {
	// Where your local node is running on localhost:5001
	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
		return "", err
	}
	//fmt.Println("added %s", cid)
	return cid, nil
}

func retreiveFromIPFS(ipfsAddress string) (io.Reader, error) {
	sh := shell.NewShell("localhost:5001")
	fileReader, err := sh.Cat(ipfsAddress)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
		return strings.NewReader(""), err //! Fix!
	}
	return fileReader, nil
}

func uploadFile(path string) io.Reader {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(data)
}

func download(fileReader io.Reader, fileName string) {
	w, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	n, err := io.Copy(w, fileReader)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}

func main() {

}
